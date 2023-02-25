package service

import "github.com/sk10az/go_anime_crud/pkg/logger"

type Interface interface {
	Data() string
}

type Service struct {
	logger *logger.Interface
}

func New() (*Service, error) {
	return &Service{}, nil
}

func (s *Service) Data() string {
	return "Default"
}
