package repo

import (
	"fmt"
	"generics/repo/models"
	"reflect"
)

type FooBar interface {
	models.Foo | models.Bar
}

func get[T FooBar](model T) T {
	fmt.Println(reflect.TypeOf(model))
	// I think bun will take model here
	return model
}
