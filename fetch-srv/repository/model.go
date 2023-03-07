package repository

type (
	List struct {
		UUID         string  `json:"uuid"`
		Komoditas    string  `json:"komoditas"`
		AreaProvinsi *string `json:"area_provinsi"`
		AreaKota     *string `json:"area_kota"`
		Size         string  `json:"size"`
		Price        string  `json:"price"`
	}

	Currency struct {
		Rates CurrencyRates `json:"rates"`
	}

	CurrencyRates struct {
		USD float64 `json:"USD"`
	}
)
