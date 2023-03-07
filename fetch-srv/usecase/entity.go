package usecase

type (
	List struct {
		UUID         string
		Komoditas    string
		AreaProvinsi *string
		AreaKota     *string
		Size         float64
		Price        float64
	}
)
