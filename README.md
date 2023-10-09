# GO BOOTCAM #1

## Firs Task
Crear un CRUD para el manejo de usuarios.

Los usuarios se deben almacenar en un map en memoria en donde el índice debe ser un id de tipo UUID. Se debe crear un struct que sustente los métodos create, get, get all, update y delete. Este struct debe ser privado y debe contar con un método constructor. 

Cada usuario debe contar con los siguientes atributos:

- ID
- Name 
- Lastname
- Email
- Active (bool)
- Address - City, Country, address


Se deben usar data types de Go, no usar JSON.

El map de usuarios debe tener 5 usuarios por default. Es decir ya creados antes de arrancar la aplicación.

## Second Task

Tomando como base el ejercicio anterior. 
Convertir la API a una Restful API, para ello se debe agregar un router y asociar los endpoints con los métodos anteriormente creados. Con el fin de implementar los endpoints se debe usar Gorilla/mux.
Los endpoints deben soportar JSON tanto en los request como en los responses. 
Los errores deben ser un objeto JSON con un campo message y otro code. 
Deben haber validaciones de los request. (tipo de dato, usuario no encontrado)
Deben haber logs de los métodos principales (usar slog)

Igualmente la API debe soportar dos tipos de storage. El primero que ya se usó es el in-memory y el segundo debe ser Redis. Por configuración se debe usar uno o el otro. (usar interfaces para lograr esta funcionalidad)


## Third Task
Para esta tercera parte se debe implementar unit testing para el package principal (service).
Para ellos se debe hacer uso de la inyección de dependencias y el mockeo para lograr cubrir la mayor cantidad de escenarios posibles. Se requiere un coverage por encima del 85%.

De igual manera se deben implementar errores custom que agreguen más detalle sobre los errores ocurridos dentro de la aplicación. 
Se deben crear por lo menos tres tipos de custom errors. StorageError que captura los errores relacionados a las dbs in-memory y Redis. ServiceError relacionado a los errores ocurridos dentro del package service y HTTPError que son los errores encontrados a nivel de handlers.
Cada tipo de error debe contener los atributos code (codigo del error - string), description (error encontrado en formato string).

Ejemplo de los errores internos:
ServiceError{
    Code: "DBConnectionFailed",
    Description: "Unable to connect to the DB"
}


Ejemplo de los errores HTTP:
HTTPError{
    Code: "NotFound",
    Status: 404,
    Description: "User not found"
}
--

A nivel de handlers van a necesitar identificar que tipo de error se envió desde la API, para eso deben hacer uso de errors.Is() o errors.As() a fin de saber que status code enviar como response.

Igual que siempre no duden en contactarme si queda alguna duda.

## Usage
you can check the package by running its tests which are writen under the file /structures/users_test.go
```go
//run and verbose all tests
go run ./... -v

```
```go
//run and verbose all tests
go run main.go
```
## Environment variables

this is an example of the .env file required.
```go
#STORAGE = Redis , to use redis, any other value will use Local Storage
STORAGE = Redis
#If STORAGE = Redis , use REDIS_ADDR var to tell the program which redis db to use
REDIS_ADDR = localhost:6379
```


