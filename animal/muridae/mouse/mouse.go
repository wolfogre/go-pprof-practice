package mouse

import (
	"log"
	"time"

	"github.com/wolfogre/go-pprof-practice/constant"
)

type Mouse struct {
	buffer     [][constant.Mi]byte
	slowBuffer [][constant.Mi]byte
}

func (*Mouse) Name() string {
	return "mouse"
}

func (m *Mouse) Live() {
	m.Eat()
	m.Drink()
	m.Shit()
	m.Pee()
	m.Hole()
	m.Steal()
}

func (m *Mouse) Eat() {
	log.Println(m.Name(), "eat")
}

func (m *Mouse) Drink() {
	log.Println(m.Name(), "drink")
}

func (m *Mouse) Shit() {
	log.Println(m.Name(), "shit")
}

func (m *Mouse) Pee() {
	log.Println(m.Name(), "pee")
	go func() {
		time.Sleep(time.Second * 10)
		max := constant.Gi
		for len(m.slowBuffer)*constant.Mi < max {
			m.slowBuffer = append(m.slowBuffer, [constant.Mi]byte{})
			time.Sleep(time.Millisecond * 500)
		}
	}()
}

func (m *Mouse) Hole() {
	log.Println(m.Name(), "hole")
}

func (m *Mouse) Steal() {
	log.Println(m.Name(), "steal")
	max := constant.Gi
	for len(m.buffer)*constant.Mi < max {
		m.buffer = append(m.buffer, [constant.Mi]byte{})
	}
}
