package dave

import "generics/repo"
import "fmt"

type FooDave struct {
	repo repo.FooRepo
}

func (d *FooDave) FooPrint() {
	foo := d.repo.GetFoo()
	fmt.Printf("%+v\n", foo)
}

func NewFooDave(fooRepo repo.FooRepo) FooDave {
	return FooDave{
		repo: fooRepo,
	}
}
