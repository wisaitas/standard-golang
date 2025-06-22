package query

import (
	"github.com/google/uuid"
	"github.com/wisaitas/share-pkg/db/repository"
)

type DistrictQuery struct {
	repository.PaginationQuery
	ProvinceID uuid.UUID `query:"province_id"`
}
