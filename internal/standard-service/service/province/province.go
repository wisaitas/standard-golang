package province

type ProvinceService interface {
	Get
}

type provinceService struct {
	Get
}

func NewProvinceService(
	get Get,
) ProvinceService {
	return &provinceService{
		Get: get,
	}
}
