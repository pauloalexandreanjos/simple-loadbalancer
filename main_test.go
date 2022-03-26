package main

import (
	"testing"

	"github.com/pauloalexandreanjos/simple-loadbalancer/models"
)

func Test_GetTask(t *testing.T) {
	server := models.NewServer("My Test Server")
	server.MockService()
	task1 := server.MockTask()
	task2 := server.MockTask()

	service, err := server.GetService("my-service-token_v1")
	if err != nil {
		t.Errorf("Expected a service got %p", service)
	}

	task, err := service.NextTask()
	if err != nil {
		t.Error(err)
	}

	if task.Id != task1 {
		t.Errorf("Expected a %q got %q", task1, task.Id)
	}

	task, err = service.NextTask()
	if err != nil {
		t.Error(err)
	}

	if task.Id != task2 {
		t.Errorf("Expected a %q got %q", task2, task.Id)
	}
}
