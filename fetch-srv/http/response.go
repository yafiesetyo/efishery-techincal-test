package http

import "fetch-srv/usecase"

type (
	List struct {
		UUID         string  `json:"uuid"`
		Komoditas    string  `json:"komoditas"`
		AreaProvinsi *string `json:"area_provinsi"`
		AreaKota     *string `json:"area_kota"`
		Size         float64 `json:"size"`
		Price        float64 `json:"price"`
	}
)

func (l List) ToResponse(in usecase.List) List {
	return List{
		UUID:         in.UUID,
		Komoditas:    in.Komoditas,
		AreaProvinsi: in.AreaProvinsi,
		AreaKota:     in.AreaKota,
		Size:         in.Size,
		Price:        in.Price,
	}
}
