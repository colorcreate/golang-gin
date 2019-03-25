package main

import (
	"log"

	local "gitlab.com/go-gin-curd/person-local"
	remote "gitlab.com/go-gin-curd/person-remote"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var remPg remote.PersonConnection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	remPg.ConnectionManager()
}

func main() {
	r := gin.Default()
	r.GET("hello", func(c *gin.Context) {
		c.String(200, "hello world!")
	})

	loc := r.Group("person-local")
	loc.GET("persons", local.GetAllPerson)
	loc.GET("person/:id", local.GetPersonByID)
	loc.POST("person", local.CreatePerson)
	loc.DELETE("person/:id", local.DeletePerson)

	rem := r.Group("person-remote")
	rem.POST("person", remPg.CreatePerson)
	rem.GET("persons", remPg.GetAllPerson)
	rem.GET("person/:id", remPg.GetPersonByID)
	rem.DELETE("person/:id", remPg.DeletePersonByID)
	r.Run()
}
