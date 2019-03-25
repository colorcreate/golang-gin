package local

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/go-gin-curd/helper"
)

var _persons []Person

// Person is
type Person struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func init() {
	_persons = []Person{
		Person{ID: 1, Name: "me", Age: 25},
		Person{ID: 2, Name: "me2", Age: 26},
		Person{ID: 3, Name: "me3", Age: 27},
	}
}

// GetAllPerson is
func GetAllPerson(c *gin.Context) {
	c.JSON(http.StatusOK, _persons)
}

// GetPersonByID is
func GetPersonByID(c *gin.Context) {
	strID := c.Param("id")
	id := helper.CUint(strID)
	person, err := getPersonByID(_persons, id)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, person)
}

// CreatePerson is
func CreatePerson(c *gin.Context) {
	var person Person
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &person)
	person = createOrUpdate(&_persons, person)
	c.JSON(http.StatusOK, person)
}

// DeletePerson is
func DeletePerson(c *gin.Context) {
	strID := c.Param("id")
	id := helper.CUint(strID)
	idx, err := getIndexByID(_persons, id)
	if err != nil {
		panic(err.Error())
	}
	_persons = append(_persons[:idx], _persons[idx+1:]...)
	c.String(http.StatusOK, "Delete Person with ID "+helper.CStr(id)+" success!")
}
