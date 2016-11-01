package main

import "fmt"

type fuel int

const (
	GASOLINE fuel = iota
	BIO
	ELECTRIC
	JET
)

type vehicle struct {
	make  string
	model string
}

type engine struct {
	fuel   fuel
	thrust int
}

func (e *engine) start() {
	fmt.Println("Engine started.")
}

type truck struct {
	vehicle
	engine
	axels  int
	wheels int
	class  int
}

func (t *truck) drive() {
	fmt.Printf("Truck %s %s, on the go!\n", t.make, t.model)
}

func newTruck(mk, mdl string) *truck {
	return &truck{vehicle: vehicle{mk, mdl}}
}

type plane struct {
	vehicle
	engine
	engineCount int
	fixedWings  bool
	maxAltitude int
}

func newPlane(mk, mdl string) *plane {
	p := &plane{}
	p.make = mk
	p.model = mdl
	return p
}

func (p *plane) fly() {
	fmt.Printf("Aircraft %s %s clear for takeoff!\n", p.make, p.model)
}

func main() {
	t := newTruck("Ford", "F750")
	t.axels = 2
	t.wheels = 6
	t.class = 3
	t.start()
	t.drive()

	p := newPlane("HondaJet", "HA-420")
	p.fuel = JET
	p.thrust = 2050
	p.engineCount = 2
	p.fixedWings = true
	p.maxAltitude = 43000
	p.start()
	p.fly()

}
