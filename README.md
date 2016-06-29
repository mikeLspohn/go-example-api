#Go Example Api

This is an example go api for learning purposes. 

Installation
  * Clone the repo into the directory where your Go projects live
  * `$ git clone https://github.com/mikeLspohn@gmail.com/go-test-api.git`
  * `$ cd go-test-api`
  * Setup a postgres database
  * Set your ENV vars in `.env`
    * `PORT` // port the server should listen on
    * `DATABASE_URL` // postgres url (e.g postgres://someone@localhost/gotest?sslmode=disable)
  * set the fallback `db_url` in `line 18 - db.go` to your db connection
  * Run Server: `go run *.go`

Packages Used: 
  * [gorilla/mux](https://github.com/gorilla/mux)
  * [justinas/alice](https://github.com/justinas/alice)
  * [rs/cors](https://github.com/rs/cors)
  * [pborman/uuid](https://github.com/pborman/uuid)
  * [joho/godotenv](https://github.com/joho/godotenv)
  * [jmoiron/sqlx](https://github.com/jmoiron/sqlx)


