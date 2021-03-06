# dependencies
- gorilla/mux — A powerful URL router and dispatcher. We use this package to match URL paths with their handlers.
- jinzhu/gorm — The fantastic ORM library for Golang, aims to be developer friendly. We use this ORM(Object relational mapper) package to interact smoothly with our database
- dgrijalva/jwt-go — Used to sign and verify JWT tokens
- joho/godotenv — Used to load .env files into the project
# Simple client api in Go - v2

1. install mux and mysql:
```
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
go get github.com/go-sql-driver/mysql
```
   

# Test
go test -v


# Compile
go build

# modules
- go mod init github.com/hfantin/clientgo


# Links
- [go + gorilla](https://medium.com/@rafaelacioly/construindo-uma-api-restful-com-go-d6007e4faff6)
- [go + gorilla + mysql](https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6)
- [go nullable](https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267)
- [rest api](https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b)



