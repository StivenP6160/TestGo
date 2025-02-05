//Gestor de tareas

package main

import "fmt"

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

//estructura que contiene como campo un slice
//Esta estructura es la que se va a encargar de almacenar la cantiad de tareas, se le debe implementar los metedos para agregar, eliminar y asignar
type ListaTarea struct {
	tareas []Tarea
}

//Método para agregar tarea
func (l *ListaTarea) agregarTareas(t Tarea) {
	l.tareas = append(l.tareas, t)
}

//Método para marcar completado la tarea
func (l *ListaTarea) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

//Método para editar tarea
func (l *ListaTarea) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

//Método para eliminar tarea
func (l *ListaTarea) eliminarTareas(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

//
func main() {
	//Instancia de lista de tareas
	lista := ListaTarea{}

	//Listar todas las tareas
	fmt.Println("Lista de tareas:")
	fmt.Println("-----------------------------------------------")
	for i, t := range lista.tareas {
		fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
	}
	fmt.Println("------------------------------------------------")
}
