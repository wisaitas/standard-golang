package pkg

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"

	validatorLib "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator interface {
	ValidateStruct(data any) error
	ValidateCommonRequestJSONBody(c *fiber.Ctx, data any) error
	ValidateCommonRequestParam(c *fiber.Ctx, data any) error
	ValidateCommonRequestFormBody(c *fiber.Ctx, data any) error
	ValidateImageFiles(files []*multipart.FileHeader, maxFileSize int64) error
	ValidateCommonQuery(c *fiber.Ctx, data any) error
}

type validator struct {
	validator *validatorLib.Validate
}

func NewValidator() Validator {
	return &validator{validator: validatorLib.New()}
}

func (v *validator) ValidateStruct(data any) error {
	return v.validator.Struct(data)
}

func (v *validator) ValidateCommonRequestJSONBody(c *fiber.Ctx, data any) error {
	if err := c.BodyParser(&data); err != nil {
		return fmt.Errorf("[validator] %w", err)
	}

	if err := v.validator.Struct(data); err != nil {
		return fmt.Errorf("[validator] %w", err)
	}

	return nil
}

func (v *validator) ValidateCommonRequestParam(c *fiber.Ctx, data any) error {
	if err := c.ParamsParser(data); err != nil {
		return fmt.Errorf("[validator] %w", err)
	}

	if err := v.validator.Struct(data); err != nil {
		return fmt.Errorf("[validator] %w", err)
	}

	return nil
}

func (v *validator) ValidateCommonRequestFormBody(c *fiber.Ctx, req any) error {
	if err := c.BodyParser(req); err != nil {
		return fmt.Errorf("[validator] %w", err)
	}

	if err := v.validator.Struct(req); err != nil {
		return fmt.Errorf("[validator] %w", err)
	}

	if form, err := c.MultipartForm(); err == nil {
		val := reflect.ValueOf(req).Elem()
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.Type() == reflect.TypeOf((*multipart.FileHeader)(nil)) {
				formTag := typ.Field(i).Tag.Get("form")

				files := form.File[formTag]

				if len(files) > 1 {
					return errors.New("multiple files are not allowed")
				}

				if len(files) > 0 {
					field.Set(reflect.ValueOf(files[0]))
				}
			}
		}
	}

	return nil
}

func (v *validator) ValidateImageFiles(files []*multipart.FileHeader, maxFileSize int64) error {
	for _, file := range files {
		if file.Size > 1024*1024*maxFileSize {
			return fmt.Errorf("[validator] image file size must be less than %dMB", maxFileSize)
		}

		if file.Size == 0 {
			return fmt.Errorf("[validator] image file is required")
		}

		if file.Filename == "" {
			return fmt.Errorf("[validator] image file name is required")
		}

		if file.Header.Get("content-type") != "image/jpeg" && file.Header.Get("content-type") != "image/png" && file.Header.Get("content-type") != "image/gif" {
			return fmt.Errorf("[validator] image file must be a valid image")
		}
	}

	return nil
}

func (v *validator) ValidateCommonQuery(c *fiber.Ctx, req any) error {
	if err := c.QueryParser(req); err != nil {
		return fmt.Errorf("[validator] %w", err)
	}

	val := reflect.ValueOf(req).Elem()
	paginationField := val.FieldByName("PaginationQuery")

	if paginationField.IsValid() {
		pagination := paginationField.Addr().Interface().(*PaginationQuery)

		if err := validatePageAndPageSize(pagination.Page, pagination.PageSize); err != nil {
			return fmt.Errorf("[validator] %w", err)
		}

		if err := validateSortAndOrder(pagination.Sort, pagination.Order); err != nil {
			return fmt.Errorf("[validator] %w", err)
		}
	}

	if err := v.validator.Struct(req); err != nil {
		return fmt.Errorf("[validator] %w", err)
	}

	return nil
}

func validatePageAndPageSize(page *int, pageSize *int) error {
	if page != nil && pageSize == nil {
		return fmt.Errorf("[validator] page_size is required")
	}

	if page == nil && pageSize != nil {
		return fmt.Errorf("[validator] page is required")
	}

	if page != nil && pageSize != nil {
		if *page < 0 {
			return fmt.Errorf("[validator] page must be greater than 0")
		}

		if *pageSize < 0 {
			return fmt.Errorf("[validator] page_size must be greater than 0")
		}
	}

	return nil
}

func validateSortAndOrder(sort *string, order *string) error {
	if sort != nil && order == nil {
		return fmt.Errorf("[validator] order is required")
	}

	if sort == nil && order != nil {
		return fmt.Errorf("[validator] sort is required")
	}

	if sort != nil && order != nil {
		if *sort == "" {
			return fmt.Errorf("[validator] sort must be a valid field")
		}

		if *order != "asc" && *order != "desc" {
			return fmt.Errorf("[validator] order must be asc or desc")
		}
	}

	return nil
}
