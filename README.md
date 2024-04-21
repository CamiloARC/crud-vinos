# CRUD Vinos

Este proyecto es una aplicación CRUD (Create, Read, Update, Delete) para la gestión de vinos. Permite realizar operaciones de creación, lectura, actualización y eliminación de registros de vinos en una base de datos.

## Tabla de contenido

- [Características](#características)
- [Tecnologías](#tecnologías)
- [Requisitos previos](#requisitos-previos)
- [Instalación](#instalación)
- [Uso](#uso)
- [API](#api)

## Características

- **Creación**: Agrega nuevos vinos a la base de datos.
- **Lectura**: Lee los registros de vinos existentes, incluyendo todos los vinos o uno específico por ID.
- **Actualización**: Modifica los detalles de un vino existente.
- **Eliminación**: Elimina un vino existente de la base de datos.

## Tecnologías

- **Lenguaje**: Go
- **Base de datos**: MySQL (u otra base de datos compatible con el paquete `database/sql`)
- **Framework web**: `net/http` de Go
- **Manejo de datos**: JSON

## Requisitos previos

- [Go](https://golang.org/doc/install) instalado en tu máquina.
- Base de datos compatible con el paquete `database/sql` (por ejemplo, MySQL).
- Asegúrate de tener acceso a la base de datos y que las credenciales de conexión estén configuradas correctamente.

## Instalación

1. Clona el repositorio

2. Instala las dependencias:

3. Configura la conexión a la base de datos en el archivo correspondiente.

## Uso

1. Ejecuta la aplicación:

    ```shell
    go run cmd\app\main.go
    ```

2. Accede a la aplicación desde tu navegador web o un cliente HTTP en el puerto definido (por defecto, el puerto 8080):

    ```
    http://localhost:8080/
    ```

## API

- **GET /vinos**: Devuelve una lista de todos los vinos.
- **GET /vinos?id=ID**: Devuelve un vino específico por ID.
- **POST /vinos**: Crea un nuevo vino. Los datos deben enviarse en formato JSON.
- **PUT /vinos?id=ID**: Actualiza un vino específico por ID.
- **DELETE /vinos?id=ID**: Elimina un vino específico por ID.