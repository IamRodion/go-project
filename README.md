# go-project

Este proyecto es un script en Go que facilita la creación de una estructura básica para un nuevo proyecto en Go. El script automatiza la creación de directorios y archivos necesarios, así como la inicialización de un módulo Go.

## Características

- Solicita al usuario el nombre del proyecto a través de la terminal.
- Crea una estructura de directorios estándar para un proyecto en Go.
- Inicializa un módulo Go con el nombre proporcionado.
- Genera un archivo `main.go` vacío listo para ser editado.

## Estructura del Proyecto

La estructura generada por el script es la siguiente:

```
nombre-del-proyecto/
│
├── main.go
└── pkg/
    ├── models/
    └── utils/
```

## Requisitos Previos

- Tener Go instalado en el sistema. Puedes descargarlo desde [golang.org](https://golang.org/dl/).
- Asegúrate de que el comando `go` esté disponible en tu PATH.

## Uso

1. Clona este repositorio en tu máquina local.
2. Navega al directorio del proyecto.
3. Ejecuta el script con el comando `go run go-project.go`.
4. Introduce el nombre del proyecto cuando se te solicite.

O también puedes intentar ejecutar el compilado que encontrarás [aquí](https://github.com/IamRodion/go-project/releases)

El script creará automáticamente la estructura de directorios y archivos necesarios.

## Detalles Técnicos

El script consta de las siguientes funciones principales:

- **GetName**: Solicita al usuario el nombre del proyecto y lo devuelve como un string.
- **RunCommand**: Ejecuta comandos del sistema en el directorio especificado.
- **CreateProject**: Utiliza `RunCommand` para crear la estructura de directorios y archivos, e inicializa el módulo Go.

## Licencia

Este proyecto está bajo la Licencia MIT. Para más detalles, consulta el archivo [LICENSE](LICENSE).
