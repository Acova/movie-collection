package app

type Port interface {
	GetPortName() string
}

type App struct {
	Ports map[string]Port
}

func NewApp() *App {
	return &App{
		Ports: make(map[string]Port),
	}
}

func (a *App) RegisterPort(port Port) {
	if _, exists := a.Ports[port.GetPortName()]; exists {
		panic("Port already registered: " + port.GetPortName())
	}
	a.Ports[port.GetPortName()] = port
}

func (a *App) GetPort(name string) Port {
	if port, exists := a.Ports[name]; exists {
		return port
	}
	panic("Port not found: " + name)
}
