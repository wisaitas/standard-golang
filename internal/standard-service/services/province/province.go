package province

type ProvinceService interface {
	Read
}

type provinceService struct {
	Read
}

func NewProvinceService(
	read Read,
) ProvinceService {
	return &provinceService{
		Read: read,
	}
}
