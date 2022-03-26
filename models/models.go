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

	return nil, errors.New("Service not found")
}

func (service *Service) NextTask() (*Task, error) {
	tasksCount := len(service.Tasks)
	if (tasksCount) == 0 {
		return nil, errors.New("Service doesn't have any task yet")
	}

	if tasksCount-1 == service.LastTask {
		service.LastTask = 0
	} else {
		service.LastTask += 1
	}

	task := service.Tasks[service.LastTask]

	log.Printf("Found Task %s\n", task.Id)
	return task, nil
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
