package local

import (
	"errors"

	"gitlab.com/go-gin-curd/helper"
)

func getPersonByID(arr []Person, ID uint) (Person, error) {
	for _, item := range arr {
		if item.ID == ID {
			return item, nil
		}
	}
	return Person{}, errors.New("Person with ID " + helper.CStr(ID) + " not found")
}

func getIndexByID(arr []Person, ID uint) (int, error) {
	for idx, item := range arr {
		if item.ID == ID {
			return idx, nil
		}
	}
	return 0, errors.New("Person with ID " + helper.CStr(ID) + " not found")
}

func getMaxID(arr []Person, ID uint) uint {
	if len(arr) != 0 {
		return arr[len(arr)-1].ID
	}
	return 0
}

func createOrUpdate(persons *[]Person, p Person) Person {
	if p.ID != 0 {
		idx, err := getIndexByID(*persons, p.ID)
		if err != nil {
			panic(err.Error())
		}
		(*persons)[idx] = p
	} else {
		p.ID = getMaxID(*persons, p.ID) + 1
		*persons = append(*persons, p)
	}
	return p
}
