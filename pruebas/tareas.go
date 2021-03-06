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

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("nombre", tarea.nombre)
		fmt.Println("descripcion", tarea.descripcion)
		fmt.Println("completado", tarea.completado)
	}
}

func (t *taskList) imprimirListaCompletos() {
	for _, tarea := range t.tasks {
		if !tarea.completado {
			continue
		}
		fmt.Println("nombre", tarea.nombre)
		fmt.Println("descripcion", tarea.descripcion)

	}
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
		completado:  true,
	}
	t2 := &task{
		nombre:      "completar mi 2",
		descripcion: "complesahdkash sdkj dhas kjhdaskjdhaskjdashk",
		completado:  false,
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	lista.imprimirListaCompletos()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Lucas"] = lista
}
