# goservice

Базовый пакет для сервисов написанных на golang

Возможности:

* Регистрация сервиса в consul
* Получение списка сервисов из consul


### Регистрация сервиса в consul

```go
package main

import "github.com/dsxack/goservice"

const (
	serviceName = "go-blank-http-server"
	servicePort = 80
)

func main() {
	err := goservice.InitConsulClient("http://consul:8500")
	if err != nil {
		panic(err)
	}
	
	err = goservice.RegisterConsulService(goservice.ConsulServiceRegistration{
		Name:          serviceName,
		Port:          servicePort,
		HealthCheck:   "/health",
	})
	if err != nil {
		panic(err)
	}
}
```

### Получение списка сервисов из consul

```go
package main

import (
	"github.com/dsxack/goservice"
)

func main() {
	err := goservice.InitConsulClient("http://consul:8500")
	if err != nil {
		panic(err)
	}
	
	services, _, err := goservice.DiscoverConsulService("go-blank-http-server")
	if err != nil {
		panic(err)
	}
}
```
