package muridae

import "github.com/wolfogre/go-pprof-practice/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}
