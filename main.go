package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type roast struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Roaster string `json:"roaster"`
	Origin  string `json:"origin"`
}

func main() {
	fmt.Println("running...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbLocation := os.Getenv("DB_LOCATION")

	dbStr := fmt.Sprintf("postgresql://%s@%s/beandb?sslmode=disable", dbUser, dbLocation)
	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/roasts", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM roasts")
		defer rows.Close()
		if err != nil {
			log.Fatalln(err)
			c.JSON(500, "An error occured")
		}

		var roasts []roast

		for rows.Next() {
			var r roast
			rows.Scan(&r.ID, &r.Name, &r.Roaster, &r.Origin)
			roasts = append(roasts, r)
		}
		c.JSON(200, roasts)
	})

	router.POST("/roasts", func(c *gin.Context) {
		var roast roast
		c.BindJSON(&roast)
		result, err := db.Exec("INSERT INTO roasts (name, roaster, origin) VALUES ($1, $2, $3)", roast.Name, roast.Roaster, roast.Origin)
		if err != nil {
			log.Fatalln("FAILED", err)
			log.Fatalln(err)
			c.JSON(500, "An error occured")
		}
		fmt.Println(result)
		c.JSON(200, "success");
	})

	router.Run("localhost:8080")
}
