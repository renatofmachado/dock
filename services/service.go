package services

import (
	"errors"
)

type Parameter struct {
	Key          string
	Question     string
	DefaultValue string
}

type Service struct {
	Name         string
	Organization string
	Image        string
	Port         string
	Template     string
	Parameters   []Parameter
}

var services = map[string]*Service{
	"redis": redis(),
}

func Resolve(serviceName string) (*Service, error) {
	service, ok := services[serviceName]

	if !ok {
		return nil, errors.New("Service is not implemented")
	}

	return service, nil
}
