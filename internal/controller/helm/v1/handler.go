package v1

import (
	helmservice "githup.com/dierbei/go-helm-api/internal/service/helm"
)

//var _ helmcontroller.Handler = (*handler)(nil)

type handler struct {
	Svc helmservice.Service
}

func NewHandler(svc helmservice.Service) *handler {
	return &handler{Svc: svc}
}
