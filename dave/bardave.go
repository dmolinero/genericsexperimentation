package dave

import "generics/repo"
import "fmt"

type BarDave struct {
	repo repo.BarRepo
}

func (d *BarDave) BarPrint() {
	foo := d.repo.GetBar()
	fmt.Printf("%+v\n", foo)
}

func NewBarDave(barRepo repo.BarRepo) BarDave {
	return BarDave{
		repo: barRepo,
	}
}
