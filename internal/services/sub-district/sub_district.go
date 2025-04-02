package sub_district

type SubDistrictService interface {
	Get
}

type subDistrictService struct {
	Get
}

func NewSubDistrictService(
	get Get,
) SubDistrictService {
	return &subDistrictService{
		Get: get,
	}
}
