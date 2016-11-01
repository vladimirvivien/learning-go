package main

import "fmt"

type fuel int

const (
	GASOLINE fuel = iota
	BIO
	ELECTRIC
	JET
)

type vehicle interface {
	start()
	stop()
}

type desc struct {
	make  string
	model string
}

type engine struct {
	fuel   fuel
	thrust int
}

func (e engine) start() {
	fmt.Println("Engine started.")
}

type truck struct {
	desc
	engine
	axels  int
	wheels int
	class  int
}

func (t truck) drive() {
	fmt.Printf("Truck %s %s, on the go!\n", t.make, t.model)
}
func (t truck) operate() {
	t.start()
	t.drive()
}

type plane struct {
	desc
	engine
	engineCount int
	fixedWings  bool
	maxAltitude int
}

func (p plane) fly() {
	fmt.Printf("Aircraft %s %s clear for takeoff!\n", p.make, p.model)
}

func (p plane) operate() {
	p.start()
	p.fly()
}

func move(v vehicle) {
	v.operate()
}

func main() {
	t := truck{
		desc:   desc{"Ford", "F750"},
		engine: engine{GASOLINE + BIO, 700},
		axels:  1,
		wheels: 6,
		class:  3,
	}
	move(t)

	p := plane{}
	p.make = "HondaJet"
	p.model = "HA-420"
	p.fuel = JET
	p.thrust = 2050
	p.engineCount = 2
	p.fixedWings = true
	p.maxAltitude = 43000
	move(p)
}
