package request

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"resty.dev/v3"
)

type H map[string]any

type HttpDefaults struct {
	headers map[string]string
}

type HttpUrlConfig struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type HttpConfig struct {
	Params  H `json:"params"`
	Data    H `json:"data"`
	Headers H `json:"headers"`
}

type HttpRequestConfig struct {
	HttpUrlConfig
	HttpConfig
}

type Http struct {
	resty *resty.Client
}

func New() *Http {
	return &Http{
		resty: resty.New(),
	}
}

func (h *Http) SetTimeout(timeout time.Duration) *Http {
	h.resty.SetTimeout(timeout)
	return h
}

func (h *Http) Request(config HttpRequestConfig) (*resty.Response, error) {

	url := strings.TrimSpace(config.Url)
	if url == "" {
		return nil, errors.New("request URL cannot be empty")
	}

	method := strings.ToUpper(strings.TrimSpace(config.Method))
	if method == "" {
		method = http.MethodGet
	}

	r := h.resty.R()

	if config.Params != nil {
		for k, v := range config.Params {
			r.SetQueryParamAny(k, v)
		}
	}

	if config.Headers != nil {
		for k, v := range config.Headers {
			r.SetHeaderAny(k, v)
		}
	}

	if config.Data != nil {
		r.SetBody(config.Data)
	}

	if config.Headers == nil {
		r.SetHeaderAny("Accept", "*/*")
	} else if _, ok := config.Headers["Accept"]; !ok {
		if _, ok := config.Headers["accept"]; !ok {
			r.SetHeaderAny("Accept", "*/*")
		}
	}

	return r.Execute(method, url)
}

func (h *Http) R() *resty.Request {
	return h.resty.R()
}

func (h *Http) Close() error {
	return h.resty.Close()
}

func (h *Http) Get(url string) (*resty.Response, error) {
	return h.resty.R().Get(url)
}

func (h *Http) GetConfig(url string, config HttpConfig) (*resty.Response, error) {
	return h.Request(HttpRequestConfig{
		HttpUrlConfig{
			Url:    url,
			Method: http.MethodGet,
		},
		config,
	})
}

func (h *Http) Post(url string) (*resty.Response, error) {
	return h.resty.R().Post(url)
}

func (h *Http) PostConfig(url string, config HttpConfig) (*resty.Response, error) {
	return h.Request(HttpRequestConfig{
		HttpUrlConfig{
			Url:    url,
			Method: http.MethodPost,
		},
		config,
	})
}
