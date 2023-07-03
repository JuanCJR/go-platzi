package main

import (
	"go-platzi/src/array"
	"go-platzi/src/ciclos"
	"go-platzi/src/condicional"
	"go-platzi/src/fmtPackage"
	"go-platzi/src/functions"
	"go-platzi/src/interfaces"
	"go-platzi/src/maps"
	"go-platzi/src/operaciones"
	"go-platzi/src/slice"
	"go-platzi/src/structsPunteros"
	"go-platzi/src/variables"
)

func main() {
	var onTest bool = true
	if onTest {

	} else {
		variables.Variables()
		operaciones.Operaciones()
		structsPunteros.StructsPunteros()
		interfaces.InterfaceFunc()
		array.ArrayFunc()
		ciclos.CiclosFunc()
		condicional.CondicionalFunc()
		fmtPackage.FmtFunc()
		functions.FunctionsFunc()
		maps.MapsFunc()
		slice.SliceFunc()
	}

}
