package sub_district

type SubDistrictService interface {
	Read
}

type subDistrictService struct {
	Read
}

func NewSubDistrictService(
	read Read,
) SubDistrictService {
	return &subDistrictService{
		Read: read,
	}
}
