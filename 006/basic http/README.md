# Golang HTTP server

- A basic HTTP server has a few key jobs to take care of.
    - *Process dynamic requests*: Process incoming requests from users who browse the website, log into their accounts or post images.
    - *Serve static assets*: Serve JavaScript, CSS and images to browsers to create a dynamic experience for the user.
    - *Accept connections*: The HTTP Server must listen on a specific port to be able to accept connections from the internet.


## Introducing Gorilla mux

- Go’s net/http package provides a lot of functionalities for the HTTP protocol. 
 One thing it doesn’t do very well is complex request routing like segmenting a 
 request url into single parameters. Fortunately there is a very popular package 
 for this, which is well known for the good code quality in the Go community. 
 
 
```bash
go get -u github.com/gorilla/mux
```

## Installing the go-sql-driver/mysql package
```bash
go get -u github.com/go-sql-driver/mysql
```

connecting to the database 

```go

	db, err := sql.Open("mysql", "root:password(127.0.0.1:3306)/GORMAX?parseTime=true")
```
- The error checks if there's any error when connecting to the database