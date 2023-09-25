# GO BOOTCAM #1

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


## Usage
you can check the package by running its tests which are writen under the file /structures/users_test.go
```go
//run and verbose all tests
go run ./... -v

```
