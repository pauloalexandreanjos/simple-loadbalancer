package models

type ServiceType int16
type TaskStatus int16

const (
	LoadBalancer ServiceType = 0
	StaticRoute              = 1
	StaticFiles              = 2
)

const (
	Healthy  TaskStatus = 0
	Degraded            = 1
	Unknown             = 2
	NotReady            = 3
	Live                = 4
)

type Server struct {
	Name     string
	Services []Service
}

type Service struct {
	Type            ServiceType
	Id              string // "0f04f1f9-8aaf-4eaf-8109-1dbb47dbff36",
	Name            string //"My Service v1",
	Token           string //"my-service-token_v1",
	Path            string //"myservice/v1",
	DefaultTaskPath string //"if present populate the taskPath on task register, else = servicePath",
	Tasks           []Task
	Rules           []Rule
}

type Task struct {
	serviceToken string // "my-service-token_v1",
	id           string // "cbd32892-8115-4ee9-b556-c33f5def0ce1",
	address      string // "localhost:1234",
	taskPath     string // "if present is the custom task path, else use service's defaultTaskPath" ,
	status       string // "HEALTHY"
}

type Rule struct {
}
