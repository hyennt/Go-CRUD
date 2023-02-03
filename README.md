# REST-API
## GO CRUD Using GIN Framework connecting with Docker (PostgreSQL)

### Air Setting CMD
Setting up environment
```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
```

Run project
```
./bin/air
```

### PostgreSQL and GORM
Enviroment

- https://customer.elephantsql.com/instance

### Setting up Env
GIN-Gonic
```
$ go get github.com/gin-gonic/gin
```
GORM 
```
$ go get gorm.io/gorm
$ go get gorm.io/driver/postgres
```

File .env to config
```
PORT =3000
DB_URL = "host=chunee.db.elephantsql.com user=jgpruamh password=I0abZRk20TlHX-owp01WqtfiP6H0eudD dbname=jgpruamh port=5432 sslmode=disable"
```

## Few note for project
- This project is used Gin-Gonic framework to handle RESTFUL-API included basic methods (GET/POST/PUT/DELETE ).
- Using PostgreSQL via TablePlus to manage database systems.
- Applying Hateoas principles to return another URL.

### References 
https://gorm.io/
https://github.com/gin-gonic/gin
https://github.com/go-gorm
