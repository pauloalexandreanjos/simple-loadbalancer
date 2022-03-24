package main

import (
	"testing"

	"github.com/pauloalexandreanjos/simple-loadbalancer/models"
)

func Test_GetTask(t *testing.T) {
	server := models.NewServer("My Test Server")
	server.MockServer()

	service, err := server.GetService("my-service-token_v1")
	if err != nil {
		t.Errorf("Expected a service got %q", service)
	}

	task := service.NextTask()
	if task == nil {
		t.Errorf("Expected a task got %q", task)
	}

	if task.Id != "1304f5ec-aa4f-11ec-b909-0242ac120002" {
		t.Errorf("Expected a %q got %q", "1304f5ec-aa4f-11ec-b909-0242ac120002", task.Id)
	}

	task = service.NextTask()
	if task == nil {
		t.Errorf("Expected a task got %q", task)
	}

	if task.Id != "77b242a6-8ffe-464e-81f8-79fc9b1fd843" {
		t.Errorf("Expected a %q got %q", "77b242a6-8ffe-464e-81f8-79fc9b1fd843", task.Id)
	}

	task = service.NextTask()
	if task == nil {
		t.Errorf("Expected a task got %q", task)
	}

	if task.Id != "1304f5ec-aa4f-11ec-b909-0242ac120002" {
		t.Errorf("Expected a %q got %q", "1304f5ec-aa4f-11ec-b909-0242ac120002", task.Id)
	}

	task = service.NextTask()
	if task == nil {
		t.Errorf("Expected a task got %q", task)
	}

	if task.Id != "77b242a6-8ffe-464e-81f8-79fc9b1fd843" {
		t.Errorf("Expected a %q got %q", "77b242a6-8ffe-464e-81f8-79fc9b1fd843", task.Id)
	}

	task = service.NextTask()
	if task == nil {
		t.Errorf("Expected a task got %q", task)
	}

	if task.Id != "1304f5ec-aa4f-11ec-b909-0242ac120002" {
		t.Errorf("Expected a %q got %q", "1304f5ec-aa4f-11ec-b909-0242ac120002", task.Id)
	}

}
