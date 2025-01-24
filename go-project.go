package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
Obtiene un nombre en formato string a través de la terminal y lo devuelve

	return => var name string // Devuelve una variable de tipo string
*/
func GetName() string {
	var name string
	fmt.Print("Ingrese el nombre del proyecto de Go a crear: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error al obtener el nombre:", err)
	}

	return name
}

/*
Toma varios argumentos y ejecuta un comando con ellos de la siguiente forma:

	dir string      => // Un string con la ruta donde se ejecutará el comando
	command string  => // Un string con el comando a ejecutar
	arg ...string   => // Una cantidad indefinida de strings para utilizar como argumentos del comando a ejecutar
*/
func RunCommand(dir string, command string, arg ...string) {
	cmd := exec.Command(command, arg...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr // Esto es útil para ver errores del comando
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error al crear el proyecto:", err)
	}

}

/*
Una lista de comandos a ejecutar, para crear el proyecto de Go con el nombre indicado

	name string => // Un string con el nombre del proyecto
*/
func CreateProject(name string) {
	RunCommand(".", "mkdir", "-p", fmt.Sprintf("%v/pkg/models", name), fmt.Sprintf("%v/pkg/utils", name))
	RunCommand(name, "touch", "main.go")
	// RunCommand(name, "go", "mod", "init", name)
	RunCommand(name, "go", "mod", "init", fmt.Sprintf("github.com/IamRodion/%v", name))

}

func main() {
	name := GetName()
	CreateProject(name)
}
