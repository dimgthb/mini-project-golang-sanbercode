package main

import (
	"database/sql"
	"fmt"
	"os"
	"practice/controllers"
	"practice/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin1"
	dbname   = "practice"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("succes read file environment")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv(`PGHOST`),
		os.Getenv(`PGPORT`),
		os.Getenv(`PGUSER`),
		os.Getenv(`PGPASSWORD`),
		os.Getenv(`PGDATABASE`))
	// psqlInfo := fmt.Sprintf("host=#{os.Getenv(`DB_HOST`)} port=#{os.Getenv(`DB_PORT`)} user=#{os.Getenv(`DB_USER`)} password=#{os.Getenv(`DB_PASSWORD`)} dbname=#{os.Getenv(`DB_NAME`)} sslmode=disable")

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)
	defer DB.Close()

	// Router GIN
	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run(":" + os.Getenv("PGPORT"))
}
