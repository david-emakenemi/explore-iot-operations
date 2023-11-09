package site

import (
	"github.com/iot-for-all/device-simulation/components/observer"
	"github.com/iot-for-all/device-simulation/components/registry"
	"github.com/iot-for-all/device-simulation/lib/component"
)

type Store component.Store[Site, component.ID]

type Component struct {
	Name       string
	RegistryID component.ID
}

type Service struct {
	Store
	registryStore registry.Store
}

func NewStore() Store {
	return component.New[Site, component.ID]()
}

func NewService(store Store, registryStore registry.Store) *Service {
	return &Service{
		Store:         store,
		registryStore: registryStore,
	}
}

func (service *Service) Create(id component.ID, c *Component) error {
	var reg registry.Observable
	reg, err := service.registryStore.Get(c.RegistryID)
	if err != nil {
		_, ok := err.(*component.NotFoundError)
		if !ok {
			return err
		}
		reg = &observer.NoopObservable{}
	}

	return service.Store.Create(New(reg, func(ss *StaticSite) {
		ss.Name = c.Name
	}), id)
}
