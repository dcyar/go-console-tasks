package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dcyar/go-console-tasks/models"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(">>>>> Bienvenido <<<<<")
		fmt.Println("Ingresa el nombre de tu proyecto o escribe \"salir\" para cerrar el programa")
		fmt.Print("-> ")
		scanner.Scan()
		if scanner.Text() == "salir" {
			break
		}

		projectName := scanner.Text()

		project := models.Project{
			Name: projectName,
		}

		for {
			fmt.Println("")
			fmt.Println(">> PROYECTO:", project.Name)
			fmt.Println("Ingresa una de las opciones")
			fmt.Println("1. Listar tareas")
			fmt.Println("2. Agregar nueva tarea")
			fmt.Println("3. Marcar tarea como completada")
			fmt.Println("4. Remover tarea")
			fmt.Println("5. Salir")
			fmt.Println("")
			fmt.Print("-> ")
			scanner.Scan()
			fmt.Println("")

			optionMenu := scanner.Text()

			if optionMenu == "5" || optionMenu == "salir" {
				goto Exit
			}

			switch optionMenu {
			case "1":
				project.PrintTasks()

			case "2":
				name := inputText("Ingresa el nombre de la tarea", *scanner)
				description := inputText("Ingresa la descripción de la tarea", *scanner)

				task := models.Task{
					Name:        name,
					Description: description,
				}

				project.AddTask(task)

				fmt.Println(">> Tarea agregada correctamente")

			case "3":
				input := inputText("Ingresa el ID de la tarea", *scanner)

				id, err := strconv.Atoi(input)

				if err != nil || len(project.Tasks) < id {
					fmt.Printf("El ID de la tarea no es válido o no existe una tarea con el ID \"%s\"\n", input)
					continue
				}

				project.MarkTaskAsCompleted(id)

				fmt.Println(">> La tarea se ha completado correctamente")

			case "4":
				input := inputText("Ingresa el ID de la tarea", *scanner)

				id, err := strconv.Atoi(input)

				if err != nil || len(project.Tasks) < id {
					fmt.Printf("El ID de la tarea no es válido o no existe una tarea con el ID \"%s\"\n", input)
					continue
				}

				project.RemoveTask(id - 1)

				fmt.Println(">> La tarea se ha removido correctamente")

			default:
				fmt.Println("Opción no válida")
			}
		}
	}

Exit:
	fmt.Println("Adios")
}

func inputText(message string, scanner bufio.Scanner) string {
	fmt.Println(message)
	fmt.Print("-> ")
	scanner.Scan()
	fmt.Println("")

	return scanner.Text()
}
