package district

type DistrictService interface {
	Read
}

type districtService struct {
	Read
}

func NewDistrictService(
	read Read,
) DistrictService {
	return &districtService{
		Read: read,
	}
}
