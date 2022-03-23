# Simple Load Balancer - SLB

This is simple load balancer a personal project to develop a simple but usable load balancer
using golang's powerfull features to achieve great performance

### Installiing

> go install

### Create redis container

> docker run -d --name redis -p 6379:6379 redis

### Running

Define in .env file REDIS_URL which is usually localhost:6379 and REDIS_PASS usually a blank space

> go build && simple-loadbalancer