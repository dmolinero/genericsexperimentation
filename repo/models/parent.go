package models

type ModelInf interface {
	SetName() string
	GetName() string
}

type Model struct {
	name string
}

func (b *Model) GetName() string {
	return b.name
}

func (b *Model) SetName(name string) string {
	b.name = name
	return b.name
}
