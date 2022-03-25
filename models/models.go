package models

import (
	"errors"
	"log"
)

type ServiceType int16
type TaskStatus int16

const (
	LoadBalancer ServiceType = 0
	StaticRoute              = 1
	StaticFiles              = 2
)

const (
	Unknown  TaskStatus = 0
	Healthy             = 1
	Degraded            = 2
	NotReady            = 3
	Live                = 4
)

type Server struct {
	Name     string
	Services []*Service
}

func NewServer(name string) *Server {

	services := make([]*Service, 0)
	server := Server{
		Name:     name,
		Services: services,
	}

	return &server
}

func (server *Server) GetService(serviceToken string) (*Service, error) {
	for i, service := range server.Services {
		if service.Token == serviceToken {
			return server.Services[i], nil
		}
	}

	return nil, errors.New("Service not found!")
}

func (server *Server) MockServer() {
	tasks := make([]*Task, 0)
	tasks = append(tasks, &Task{
		ServiceToken: "my-service-token_v1",
		Id:           "1304f5ec-aa4f-11ec-b909-0242ac120002",
		Address:      "http://localhost:8000",
		Status:       Healthy,
	})
	tasks = append(tasks, &Task{
		ServiceToken: "my-service-token_v1",
		Id:           "77b242a6-8ffe-464e-81f8-79fc9b1fd843",
		Address:      "http://localhost:8000",
		Status:       Healthy,
	})
	server.Services = append(server.Services, &Service{
		Type:     LoadBalancer,
		Id:       "be2d3f4c-eec7-4cab-a782-6c262e6f04d0",
		Name:     "My Service V1",
		Token:    "my-service-token_v1",
		Path:     "/main.go",
		Tasks:    tasks,
		LastTask: -1,
		Rules:    make([]*Rule, 0),
	})
}

func (service *Service) NextTask() *Task {
	tasksCount := len(service.Tasks)

	if tasksCount-1 == service.LastTask {
		service.LastTask = 0
	} else {
		service.LastTask += 1
	}

	task := service.Tasks[service.LastTask]

	log.Printf("Found Task %s\n", task.Id)
	return task
}

func (service *Service) AddTask(task *Task) {
	service.Tasks = append(service.Tasks, task)
}

type Service struct {
	Type            ServiceType
	Id              string
	Name            string
	Token           string
	Path            string
	DefaultTaskPath string
	Tasks           []*Task
	Rules           []*Rule
	LastTask        int
}

type Task struct {
	ServiceToken string
	Id           string
	Address      string
	TaskPath     string
	Status       TaskStatus
}

type Rule struct {
	Id             string
	Route          string
	Rate           string
	MaxPayloadSize string
	KeywordFilter  string
}

type Register struct {
	ServiceToken string `json:"serviceToken"`
	HealthUrl    string `json:"healthUrl"`
}
