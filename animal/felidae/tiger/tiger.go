package tiger

import "log"

type Tiger struct {
}

func (t *Tiger) Name() string {
	return "tiger"
}

func (t *Tiger) Live() {
	t.Eat()
	t.Drink()
	t.Shit()
	t.Pee()
	t.Climb()
	t.Sneak()
}

func (t *Tiger) Eat() {
	log.Println(t.Name(), "eat")
	loop := 10000000000
	for i := 0; i < loop; i++ {
		// do nothing
	}
}

func (t *Tiger) Drink() {
	log.Println(t.Name(), "drink")
}

func (t *Tiger) Shit() {
	log.Println(t.Name(), "shit")
}

func (t *Tiger) Pee() {
	log.Println(t.Name(), "pee")
}

func (t *Tiger) Climb() {
	log.Println(t.Name(), "climb")
}

func (t *Tiger) Sneak() {
	log.Println(t.Name(), "sneak")
}
