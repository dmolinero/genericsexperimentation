package main

import (
	"generics/dave"
	"generics/repo"
)

func main() {
	fooRepo := repo.NewFooRepo()
	barRepo := repo.NewBarRepo()

	fooDave := dave.NewFooDave(fooRepo)
	barDave := dave.NewBarDave(barRepo)

	fooDave.FooPrint()
	barDave.BarPrint()
}
