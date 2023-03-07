package usecase

import (
	"fetch-srv/repository"
	"fetch-srv/utils/logger"
	"fmt"
	"sort"
	"strconv"
)

var (
	USD float64
)

type (
	IUsecase interface {
		FetchAggregate() (map[string]interface{}, error)
		Fetch() (out []List, err error)
	}

	usecase struct {
		repo repository.IRepo
	}
)

func New(repo repository.IRepo) IUsecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Fetch() (out []List, err error) {
	logCtx := fmt.Sprintf("%T.Fetch", *u)

	if USD < 1 {
		currency, err := u.repo.GetCurrency()
		if err != nil {
			logger.Error(logCtx, "repo.GetCurrency error: %v", err)
			return out, err
		}

		USD = currency.Rates.USD
	}

	data, err := u.repo.GetList()
	if err != nil {
		logger.Error(logCtx, "repo.GetList error: %v", err)
		return
	}

	for _, d := range data {
		price, err := strconv.ParseFloat(d.Price, 64)
		if err != nil {
			logger.Error(logCtx, "parseFloat error: %v", err)
			return out, err
		}
		price = price * USD

		size, err := strconv.ParseFloat(d.Size, 64)
		if err != nil {
			logger.Error(logCtx, "parseFloat error: %v", err)
			return out, err
		}

		out = append(out, List{
			UUID:         d.UUID,
			Komoditas:    d.Komoditas,
			AreaProvinsi: d.AreaProvinsi,
			AreaKota:     d.AreaKota,
			Size:         size,
			Price:        price,
		})
	}

	return
}

func (u *usecase) FetchAggregate() (out map[string]interface{}, err error) {
	logCtx := fmt.Sprintf("%T.FetchAggregate", *u)

	if USD < 1 {
		currency, err := u.repo.GetCurrency()
		if err != nil {
			logger.Error(logCtx, "repo.GetCurrency error: %v", err)
			return out, err
		}

		USD = currency.Rates.USD
	}

	list, err := u.repo.GetList()
	if err != nil {
		logger.Error(logCtx, "repo.GetList error: %v", err)
		return
	}

	res := map[string]interface{}{}
	priceMedian := map[string][]float64{}
	sizeMedian := map[string][]float64{}

	for _, l := range list {
		price, err := strconv.ParseFloat(l.Price, 64)
		if err != nil {
			logger.Error(logCtx, "parseFloat error: %v", err)
			return out, err
		}
		price = price * USD

		size, err := strconv.ParseFloat(l.Size, 64)
		if err != nil {
			logger.Error(logCtx, "parseFloat error: %v", err)
			return out, err
		}

		if l.AreaProvinsi != nil {
			if _, ok := res[*l.AreaProvinsi]; !ok {
				res[*l.AreaProvinsi] = map[string]interface{}{
					"price": map[string]interface{}{
						"avg":    price,
						"min":    price,
						"max":    price,
						"median": price,
						"count":  uint64(1),
						"total":  price,
					},
					"size": map[string]interface{}{
						"avg":    size,
						"min":    size,
						"max":    size,
						"median": size,
						"count":  uint64(1),
						"total":  size,
					},
				}
				priceMedian[*l.AreaProvinsi] = append(priceMedian[*l.AreaProvinsi], price)
				sizeMedian[*l.AreaProvinsi] = append(sizeMedian[*l.AreaProvinsi], size)

				continue
			}

			priceMedian[*l.AreaProvinsi] = append(priceMedian[*l.AreaProvinsi], price)
			sizeMedian[*l.AreaProvinsi] = append(sizeMedian[*l.AreaProvinsi], size)

			currentPriceAvg := res[*l.AreaProvinsi].(map[string]interface{})["price"].(map[string]interface{})["avg"].(float64)
			currentPriceMin := res[*l.AreaProvinsi].(map[string]interface{})["price"].(map[string]interface{})["min"].(float64)
			currentPriceMax := res[*l.AreaProvinsi].(map[string]interface{})["price"].(map[string]interface{})["max"].(float64)
			currentPriceMedian := res[*l.AreaProvinsi].(map[string]interface{})["price"].(map[string]interface{})["median"].(float64)
			currentPriceCount := res[*l.AreaProvinsi].(map[string]interface{})["price"].(map[string]interface{})["count"].(uint64)
			currentPriceTotal := res[*l.AreaProvinsi].(map[string]interface{})["price"].(map[string]interface{})["total"].(float64)

			currentSizeAvg := res[*l.AreaProvinsi].(map[string]interface{})["size"].(map[string]interface{})["avg"].(float64)
			currentSizeMin := res[*l.AreaProvinsi].(map[string]interface{})["size"].(map[string]interface{})["min"].(float64)
			currentSizeMax := res[*l.AreaProvinsi].(map[string]interface{})["size"].(map[string]interface{})["max"].(float64)
			currentSizeMedian := res[*l.AreaProvinsi].(map[string]interface{})["size"].(map[string]interface{})["median"].(float64)
			currentSizeCount := res[*l.AreaProvinsi].(map[string]interface{})["size"].(map[string]interface{})["count"].(uint64)
			currentSizeTotal := res[*l.AreaProvinsi].(map[string]interface{})["size"].(map[string]interface{})["total"].(float64)

			if price > currentPriceMax {
				currentPriceMax = price
			}
			if price < currentPriceMin {
				currentPriceMin = price
			}
			currentPriceAvg = (float64(currentPriceTotal) + price) / (float64(currentPriceCount) + 1)
			currentPriceMedian = getMedian(priceMedian[*l.AreaProvinsi])

			if size > currentSizeMax {
				currentSizeMax = size
			}
			if size < currentSizeMin {
				currentSizeMin = size
			}
			currentSizeAvg = (float64(currentSizeTotal) + size) / (float64(currentSizeCount) + 1)
			currentSizeMedian = getMedian(sizeMedian[*l.AreaProvinsi])

			res[*l.AreaProvinsi] = map[string]interface{}{
				"price": map[string]interface{}{
					"avg":    currentPriceAvg,
					"min":    currentPriceMin,
					"max":    currentPriceMax,
					"median": currentPriceMedian,
					"count":  uint64(currentPriceCount + 1),
					"total":  float64(currentPriceTotal) + price,
				},
				"size": map[string]interface{}{
					"avg":    currentSizeAvg,
					"min":    currentSizeMin,
					"max":    currentSizeMax,
					"median": currentSizeMedian,
					"count":  uint64(currentSizeCount + 1),
					"total":  float64(currentSizeTotal) + size,
				},
			}
		}

	}

	return res, nil
}

func getMedian(in []float64) float64 {
	sort.Float64s(in)

	medIdx := len(in) / 2
	if medIdx%2 > 0 {
		return in[medIdx]
	}

	return (in[medIdx-1] + in[medIdx]) / 2
}
