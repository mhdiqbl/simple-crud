package handler

import "ottoDigital/repository"

type handler struct {
	repository repository.IRepository
}

func NewHandler() *handler {
	return &handler{
		repository: repository.NewRepository(),
	}
}
