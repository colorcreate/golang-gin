package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.com/go-gin-curd/helper"

	"github.com/gin-gonic/gin"
)

// Person is
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CreatePerson is
func (r *PersonConnection) CreatePerson(c *gin.Context) {
	var person Person
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &person)
	fmt.Println(person)
	err = r.createPerson(&person)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, person)
}

// GetAllPerson is
func (r *PersonConnection) GetAllPerson(c *gin.Context) {
	persons := r.getAllPerson()
	c.JSON(http.StatusOK, persons)
}

// GetPersonByID is
func (r *PersonConnection) GetPersonByID(c *gin.Context) {
	strID := c.Param("id")
	id := helper.CUint(strID)
	person, err := r.getPersonByID(id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, person)
}

// DeletePersonByID is
func (r *PersonConnection) DeletePersonByID(c *gin.Context) {
	strID := c.Param("id")
	id := helper.CUint(strID)
	err := r.deletePersonByID(id)
	if err != nil {
		panic(err)
	}
	c.String(http.StatusOK, "Delete Person with ID = "+helper.CStr(id)+" success!")
}
