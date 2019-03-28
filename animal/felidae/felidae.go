package felidae

import "github.com/wolfogre/go-pprof-practice/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}
