package services

import "backend/internal/application/interfaces"

type ServerService struct {
}

func NewServerService() interfaces.ServerService {
	return &ServerService{}
}
