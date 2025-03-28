package validates

import (
	"errors"
	"mime/multipart"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/configs"
	"github.com/wisaitas/standard-golang/pkg"
)

func validateCommonRequestJSONBody[T any](c *fiber.Ctx, req *T, validator pkg.ValidatorUtil) error {
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(err)
	}

	if err := validator.Validate(req); err != nil {
		return pkg.Error(err)
	}

	return nil
}

func validateCommonRequestParams[T any](c *fiber.Ctx, req *T, validator pkg.ValidatorUtil) error {
	if err := c.ParamsParser(req); err != nil {
		return pkg.Error(err)
	}

	if err := validator.Validate(req); err != nil {
		return pkg.Error(err)
	}

	return nil
}

func validateCommonRequestFormBody[T any](c *fiber.Ctx, req *T, validator pkg.ValidatorUtil) error {
	if err := c.BodyParser(req); err != nil {
		return pkg.Error(err)
	}

	if err := validator.Validate(req); err != nil {
		return pkg.Error(err)
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

func validateImageFiles(files []*multipart.FileHeader) error {
	maxFileSize := configs.ENV.MAX_FILE_SIZE

	for _, file := range files {
		if file.Size > 1024*1024*maxFileSize {
			return pkg.Error(errors.New("image file size must be less than " + strconv.FormatInt(maxFileSize, 10) + "MB"))
		}

		if file.Size == 0 {
			return pkg.Error(errors.New("image file is required"))
		}

		if file.Filename == "" {
			return pkg.Error(errors.New("image file name is required"))
		}

		if file.Header.Get("content-type") != "image/jpeg" && file.Header.Get("content-type") != "image/png" && file.Header.Get("content-type") != "image/gif" {
			return pkg.Error(errors.New("image file must be a valid image"))
		}
	}

	return nil
}

func validateCommonPaginationQuery[T any](c *fiber.Ctx, req *T, validator pkg.ValidatorUtil) error {
	if err := c.QueryParser(req); err != nil {
		return pkg.Error(err)
	}

	val := reflect.ValueOf(req).Elem()
	paginationField := val.FieldByName("PaginationQuery")

	if paginationField.IsValid() {
		pagination := paginationField.Addr().Interface().(*pkg.PaginationQuery)

		if err := validatePageAndPageSize(pagination.Page, pagination.PageSize); err != nil {
			return pkg.Error(err)
		}

		if err := validateSortAndOrder(pagination.Sort, pagination.Order); err != nil {
			return pkg.Error(err)
		}
	}

	if err := validator.Validate(req); err != nil {
		return err
	}

	return nil
}

func validatePageAndPageSize(page *int, pageSize *int) error {
	if page != nil && pageSize == nil {
		return pkg.Error(errors.New("pageSize is required"))
	}

	if page == nil && pageSize != nil {
		return pkg.Error(errors.New("page is required"))
	}

	if page != nil && pageSize != nil {
		if *page < 0 {
			return pkg.Error(errors.New("page must be greater than 0"))
		}

		if *pageSize < 0 {
			return pkg.Error(errors.New("pageSize must be greater than 0"))
		}
	}

	return nil
}

func validateSortAndOrder(sort *string, order *string) error {
	if sort != nil && order == nil {
		return pkg.Error(errors.New("order is required"))
	}

	if sort == nil && order != nil {
		return pkg.Error(errors.New("sort is required"))
	}

	if sort != nil && order != nil {
		if *sort == "" {
			return pkg.Error(errors.New("sort must be a valid field"))
		}

		if *order != "asc" && *order != "desc" {
			return pkg.Error(errors.New("order must be asc or desc"))
		}
	}

	return nil
}
