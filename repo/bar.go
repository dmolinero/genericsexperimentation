package repo

import (
	"generics/repo/models"
)

type BarRepo struct {
}

func (br *BarRepo) GetBar() models.Bar {
	return get[models.Bar](models.Bar{})
}

func NewBarRepo() BarRepo {
	return BarRepo{}
}
