package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hfantin/clientgo-echo/client"
	"github.com/hfantin/clientgo-echo/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

type Cli struct {
	ID        int             `json:"id"`
	Name      *string         `json:"name"`
	BirthDate *string         `json:"bithDate"`
	Email     *sql.NullString `json:"email"`
}

var db *sql.DB
var logger = logrus.New()
var consul client.Consul

// var redis client.RedisClient

func main() {
	env := utils.Config()
	db = initDB(env)
	initLogger(env.LoggerLevel)
	consul = client.Consul{}
	// redis = client.RedisClient{}
	// redis.New()
	consul.Register(env.ServerPort)
	initServer(env.ServerPort)

}

func initServer(port string) {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/v1/clients", getClients)
	e.GET("/actuator/health", health)
	e.GET("/actuator/info", info)
	logger.Info("Starting clientgo-echo server on port ", port)
	e.Server.Addr = ":" + port
	graceful.ListenAndServe(e.Server, 5*time.Second)
	logger.Info("Stopping server...")
	consul.Deregister()

}

// loggger https://github.com/antonfisher/nested-logrus-formatter
func initLogger(level string) {
	lv, _ := logrus.ParseLevel(level)
	logger.SetLevel(lv)
	logger.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        true,
		ShowFullLevel:   false,
		FieldsOrder:     []string{"component", "category"},
	})
}
func initDB(env utils.Environment) *sql.DB {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s", env.DBUser, env.DBPass, env.DBHost, env.DBName)
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	return db
}

func getClients(ctx echo.Context) error {
	logger.Debug("get clients")

	// chave := "teste"
	// cache, err := redis.Get(chave)
	// if err != nil {
	// 	logger.Info("testando cache")
	// 	redis.Set(chave, []byte("teste"))
	// } else {
	// 	logger.Info("cache: ", string(cache))

	// }

	rows, err := db.Query("SELECT * FROM CLI")
	if err != nil {
		logger.Error("Falha ao executar query de clientes: ", err)
		// https://golang.org/src/net/http/status.go
		return ctx.JSON(http.StatusBadRequest, err)
	}
	defer rows.Close()
	clients := make([]Cli, 0)
	for rows.Next() {
		cli := Cli{}
		if err := rows.Scan(&cli.ID, &cli.Name, &cli.BirthDate, &cli.Email); err != nil {
			logger.Error("Falha ao executar ler cursor de clientes: ", err)
			return ctx.JSON(http.StatusBadRequest, err)
		}
		clients = append(clients, cli)
	}

	return ctx.JSON(http.StatusOK, clients)
}

func info(ctx echo.Context) error {
	build := make(map[string]interface{})
	status := make(map[string]interface{})
	status["name"] = "clientgo-echo"
	status["version"] = "0.0.1"
	build["build"] = status
	return ctx.JSON(http.StatusOK, build)
}

func health(ctx echo.Context) error {
	health := make(map[string]interface{})
	health["status"] = "UP"
	return ctx.JSON(http.StatusOK, health)
}
