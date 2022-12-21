package handler

import "sales/service"

type handler struct{

	s	service.Service
}


func NewHandler(s service.Service) *handler{
	return &handler{s}
}