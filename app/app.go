package app

import "reflect"

type Port interface {
	GetPortName() string
}

type App struct {
	Ports map[reflect.Type]Port
}

func NewApp() *App {
	return &App{
		Ports: make(map[reflect.Type]Port),
	}
}

func (a *App) RegisterPort(port Port) {
	if _, exists := a.Ports[reflect.TypeOf(port)]; exists {
		panic("Port already registered: " + port.GetPortName())
	}
	a.Ports[reflect.TypeOf(port)] = port
}

func (a *App) GetPort(t reflect.Type) Port {
	if port, exists := a.Ports[t]; exists {
		return port
	}
	panic("Port not found: " + t.Name())
}
