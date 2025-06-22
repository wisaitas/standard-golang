package query

import (
	"github.com/google/uuid"
	"github.com/wisaitas/share-pkg/db/repository"
)

type SubDistrictQuery struct {
	repository.PaginationQuery
	DistrictID uuid.UUID `query:"district_id"`
}
