package repo

import (
	"generics/repo/models"
)

type FooRepo struct {
}

func (b *FooRepo) GetFoo() models.Foo {
	return get[models.Foo](models.Foo{})
}

func NewFooRepo() FooRepo {
	return FooRepo{}
}
