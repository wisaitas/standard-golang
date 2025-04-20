package district

type DistrictService interface {
	Get
}

type districtService struct {
	Get
}

func NewDistrictService(
	get Get,
) DistrictService {
	return &districtService{
		Get: get,
	}
}
