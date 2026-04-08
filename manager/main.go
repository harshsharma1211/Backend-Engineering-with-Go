package main

type Truck struct {
	id    string
	cargo int
}

type TruckManager struct {
	truck_id map[string]*truck
}

type FleetManager interface {
	AddTruck(truck *Truck)
	GetTruck(id string) (*Truck, error)
	DeleteTruck(id string) error
	GetAllTrucks() []*Truck
}
