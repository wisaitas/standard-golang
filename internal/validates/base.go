package validates

import (
	"errors"
	"mime/multipart"
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/configs"
	"github.com/wisaitas/standard-golang/internal/dtos/request"
)

func validateCommonRequestJSONBody[T any](c *fiber.Ctx, req *T) error {
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := validator.New().Struct(req); err != nil {
		return err
	}

	return nil
}

func validateCommonRequestFormBody[T any](c *fiber.Ctx, req *T) error {
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if err := validator.New().Struct(req); err != nil {
		return err
	}

	if form, err := c.MultipartForm(); err == nil {
		val := reflect.ValueOf(req).Elem()
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.Type() == reflect.TypeOf([]*multipart.FileHeader{}) {
				formTag := typ.Field(i).Tag.Get("form")
				if files := form.File[formTag]; files != nil {
					field.Set(reflect.ValueOf(files))
				}
			}
		}
	}

	return nil
}

func validateImageFiles(files []*multipart.FileHeader) error {
	maxFileSize := configs.ENV.MAX_FILE_SIZE

	for _, file := range files {
		if file.Size > 1024*1024*maxFileSize {
			return errors.New("image file size must be less than " + strconv.FormatInt(maxFileSize, 10) + "MB")
		}

		if file.Size == 0 {
			return errors.New("image file is required")
		}

		if file.Filename == "" {
			return errors.New("image file name is required")
		}

		if file.Header.Get("content-type") != "image/jpeg" && file.Header.Get("content-type") != "image/png" && file.Header.Get("content-type") != "image/gif" {
			return errors.New("image file must be a valid image")
		}
	}

	return nil
}

func validateCommonPaginationQuery(c *fiber.Ctx, req *request.PaginationQuery) error {
	if err := c.QueryParser(req); err != nil {
		return err
	}

	if err := validatePageAndPageSize(req.Page, req.PageSize); err != nil {
		return err
	}

	if err := validateSortAndOrder(req.Sort, req.Order); err != nil {
		return err
	}

	return nil
}

func validatePageAndPageSize(page *int, pageSize *int) error {
	if page != nil && pageSize == nil {
		return errors.New("pageSize is required")
	}

	if page == nil && pageSize != nil {
		return errors.New("page is required")
	}

	if page != nil && pageSize != nil {
		if *page < 0 {
			return errors.New("page must be greater than 0")
		}

		if *pageSize < 0 {
			return errors.New("pageSize must be greater than 0")
		}
	}

	return nil
}

func validateSortAndOrder(sort *string, order *string) error {
	if sort != nil && order == nil {
		return errors.New("order is required")
	}

	if sort == nil && order != nil {
		return errors.New("sort is required")
	}

	if sort != nil && order != nil {
		if *sort == "" {
			return errors.New("sort must be a valid field")
		}

		if *order != "asc" && *order != "desc" {
			return errors.New("order must be asc or desc")
		}
	}

	return nil
}
