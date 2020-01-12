[![GCI Badge](https://img.shields.io/badge/Google%20Code--in-JBoss%20Community-red?labelColor=2096F3)](https://gitter.im/JBossOutreach/gci)

# go-crud

## Prerequisites
- Have Go installed
- Have a working internet connection

## Run
- `go -u get https://github.com/gorilla/mux`
- `go run ./main.go`

## Use:
### Create, Read, Update, Delete Food Items with [Name] and [Price]
#### Running On: localhost:4000

### Requests Available:
> #### localhost:4000/
- HTTP Method: GET
- Use: Home Dir
- Request Body: N/A

> #### localhost:4000/api
- HTTP Method: GET
- Use: List all Foods
- Request Body: N/A

> #### localhost:4000/api/{id}
- HTTP Method: GET
- Use: Get the food with id mentioned
- Request Body: N/A

> #### localhost:4000/api/
- HTTP Method: POST
- Use: Add food to the api
- Request Body: JSON

> #### localhost:4000/api/{id}
- HTTP Method: PUT
- Use: Update food with Specified ID
- Request Body: JSON

> #### localhost:4000/api/{id}
- HTTP Method: DELETE
- Use: Delete food with Specified ID
- Request Body: JSON

### Request Body JSON Structure:
```
{
  "Name": String
  "Price": Int
}
```

## Working Demo

[Here](https://youtu.be/4dy_DAO7bDc)

Google Code-in and the Google Code-in logo are trademarks of Google Inc.
