package models

import (
	"math/rand"
	"strconv"
)

func (server *Server) MockService() {
	server.Services = append(server.Services, &Service{
		Type:     LoadBalancer,
		Id:       "be2d3f4c-eec7-4cab-a782-6c262e6f04d0",
		Name:     "My Service V1",
		Token:    "my-service-token_v1",
		Path:     "/",
		Tasks:    make([]*Task, 0),
		LastTask: -1,
		Rules:    make([]*Rule, 0),
	})
}

func (server *Server) MockTask() string {
	service := server.Services[0]
	taskId := strconv.Itoa(rand.Int())
	service.Tasks = append(service.Tasks, &Task{
		ServiceToken: "my-service-token_v1",
		Id:           taskId,
		Address:      "http://localhost:8000",
	})

	services := make([]*Service, 0)
	services = append(services, service)

	server.Services = services
	return taskId
}
