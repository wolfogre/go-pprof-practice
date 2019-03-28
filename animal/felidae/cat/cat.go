package cat

import (
	"log"
	"time"
)

type Cat struct {
}

func (c *Cat) Name() string {
	return "cat"
}

func (c *Cat) Live() {
	c.Eat()
	c.Drink()
	c.Shit()
	c.Pee()
	c.Climb()
	c.Sneak()
}

func (c *Cat) Eat() {
	log.Println(c.Name(), "eat")
}

func (c *Cat) Drink() {
	log.Println(c.Name(), "drink")
}

func (c *Cat) Shit() {
	log.Println(c.Name(), "shit")
}

func (c *Cat) Pee() {
	log.Println(c.Name(), "pee")

	<-time.After(time.Second)
}

func (c *Cat) Climb() {
	log.Println(c.Name(), "climb")
}

func (c *Cat) Sneak() {
	log.Println(c.Name(), "sneak")
}
