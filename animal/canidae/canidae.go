package canidae

import "github.com/wolfogre/go-pprof-practice/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}
