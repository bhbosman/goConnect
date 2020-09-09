package rest

import "net/http"

type baseService struct {
	apiKey  string
	client  http.Client
	address string
}
