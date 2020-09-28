package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tarea *task) {
	t.tasks = append(t.tasks, tarea)
}

func (t *taskList) quitarDeLista(indice int) {
	t.tasks = append(t.tasks[:indice], t.tasks[indice+1:]...)
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
	t1 := &task{
		nombre:      "completar mi curso",
		descripcion: "complesahdkash sdkj dhas kjhdaskjdhaskjdashk",
		completado:  false,
	}
	t2 := &task{
		nombre:      "completar mi 2",
		descripcion: "complesahdkash sdkj dhas kjhdaskjdhaskjdashk",
		completado:  false,
	}
	t3 := &task{
		nombre:      "completar mi 3",
		descripcion: "complesahdkash sdkj dhas kjhdaskjdhaskjdashk",
		completado:  false,
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(lista.tasks[0])
	lista.agregarALista(t3)

	lista.quitarDeLista(1)
	fmt.Println(len(lista.tasks))
}
