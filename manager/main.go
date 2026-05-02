package main

import (
	"errors"
)
var ErrTruckNotFound = errors.New("Truck not found ")

type Truck struct {
	id    string
	cargo int
}

type TruckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() *TruckManager {
	return &TruckManager{
		trucks: make(map[string]*Truck),
	}
}

type FleetManager interface {
	AddTruck(truck *Truck)
	GetTruck(id string) (*Truck, error)
	DeleteTruck(id string) error
	GetAllTrucks() []*Truck
}

func AddTruck(truck *Truck) {

}

// func main() {

// 	trucks := []Truck{
// 		{id: "Truck_1"},
// 		{id: "Truck_2"},
// 	}

// 	for _, truck := range trucks {

// 		fmt.Printf("Trucks %s loaded /n",truck.id)
// 		// ProcessTruck()
// 	}

// }
