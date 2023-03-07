package infra

import (
	"fetch-srv/utils/logger"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	HttpClientIface interface {
		Get(url string, headers *map[string]string) (bt []byte, err error)
	}

	httpClient struct {
		client http.Client
	}
)

func NewHttpClient(client http.Client) HttpClientIface {
	return &httpClient{
		client: client,
	}
}

func (h *httpClient) Get(url string, headers *map[string]string) (bt []byte, err error) {
	logCtx := fmt.Sprintf("%T.Get", *h)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error(logCtx, "http.NewRequest error: %v", err)
		return
	}
	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}

	res, err := h.client.Do(req)
	if err != nil {
		logger.Error(logCtx, "client.Do error: %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return bt, fmt.Errorf("getting %d status from %s", res.StatusCode, url)
	}

	bt, err = ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(logCtx, "ioutil.ReadAll error: %v", err)
		return
	}

	return
}
