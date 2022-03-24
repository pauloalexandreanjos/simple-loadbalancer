package main

import (
	"errors"

	"github.com/pauloalexandreanjos/simple-loadbalancer/models"
)

func getTask(serviceToken string) (*models.Task, error) {
	for _, service := range server.Services {
		if service.Token == serviceToken {
			return service.NextTask(), nil
		}
	}

	return nil, errors.New("Task not found!")
}
