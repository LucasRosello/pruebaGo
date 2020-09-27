package main

import "fmt"

type taskList struct {
	tasks []*task
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := task{
		nombre:      "completar mi curso",
		descripcion: "complesahdkash sdkj dhas kjhdaskjdhaskjdashk",
		completado:  false,
	}
	fmt.Println(t)
	fmt.Printf("%+v!\n", t)
	t.marcarCompleta()
	t.actualizarNombre("final")
	t.actualizarDescripcion("chacha")
	fmt.Printf("%+v!\n", t)
}
