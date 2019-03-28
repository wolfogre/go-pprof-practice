package wolf

import (
	"log"
	"sync"
	"time"
)

type Wolf struct {
	buffer []byte
}

func (w *Wolf) Name() string {
	return "wolf"
}

func (w *Wolf) Live() {
	w.Eat()
	w.Drink()
	w.Shit()
	w.Pee()
	w.Run()
	w.Howl()
}

func (w *Wolf) Eat() {
	log.Println(w.Name(), "eat")
}

func (w *Wolf) Drink() {
	log.Println(w.Name(), "drink")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}

func (w *Wolf) Shit() {
	log.Println(w.Name(), "shit")
}

func (w *Wolf) Pee() {
	log.Println(w.Name(), "pee")
}

func (w *Wolf) Run() {
	log.Println(w.Name(), "run")
}

func (w *Wolf) Howl() {
	log.Println(w.Name(), "howl")

	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}
