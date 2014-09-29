package reveldi

// Simple service container
type Container struct {
	services map[string]Service
}

type Service interface{}

func (c *Container) Register(name string, serviceStruct Service) {
	if len(c.services) == 0 {
		c.services = make(map[string]Service)
	}
	c.services[name] = serviceStruct
}

func (c *Container) Get(name string) Service {
	return c.services[name]
}
