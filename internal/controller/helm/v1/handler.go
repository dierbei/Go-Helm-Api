package v1

import (
	helmservice "githup.com/dierbei/go-helm-api/internal/service/helm"
)

type handler struct {
	Svc helmservice.Service
}

func NewHandler(svc helmservice.Service) *handler {
	return &handler{Svc: svc}
}
