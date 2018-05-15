package go_consul

import (
	"net"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"strconv"
)

const (
	PublicWebServiceTag = "web"
)

var consulClient *consul.Client

// Service config
type ConsulServiceRegistration struct {
	// Consul service name
	Name string

	// Consul service id
	ID string

	// Consul service port
	Port int

	// Consul service health check url address
	HealthCheck string

	// Consul service tags
	Tags []string
}

func (c *ConsulServiceRegistration) validate() error {
	if c.Name == "" {
		return fmt.Errorf("service name must be passed")
	}

	if c.Port <= 0 {
		return fmt.Errorf("service port must be passed")
	}

	return nil
}

// Register new consul service by config params
func RegisterConsulService(config ConsulServiceRegistration) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("error to register consul service: %v", err)
		}
	}()

	err = config.validate()
	if err != nil {
		return
	}

	if consulClient == nil {
		err = fmt.Errorf("consul client must be initialized first")
		return
	}

	ip, err := getSelfIP()
	if err != nil {
		return
	}

	reg := &consul.AgentServiceRegistration{
		ID:      config.ID,
		Name:    config.Name,
		Port:    config.Port,
		Address: ip,
		Tags:    config.Tags,
	}

	if config.HealthCheck != "" {
		reg.Check = &consul.AgentServiceCheck{
			Interval: "10s",
			HTTP:     "http://" + ip + ":" + strconv.Itoa(config.Port) + config.HealthCheck,
		}
	}

	err = consulClient.Agent().ServiceRegister(reg)
	if err != nil {
		return
	}

	return
}

// Discover consul service by name
func DiscoverConsulService(name string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	addrs, meta, err := consulClient.Health().Service(name, "", true, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", name)
	}
	if err != nil {
		return nil, nil, err
	}
	return addrs, meta, nil
}

// Init consul client by address
func InitConsulClient(address string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("error to init consul client: %v", err)
		}
	}()

	if address == "" {
		err = fmt.Errorf("address must be passed")
		return
	}

	consulClient, err = consul.NewClient(&consul.Config{
		Address: address,
	})

	return
}

func getSelfIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("error to fetch ip")
}
