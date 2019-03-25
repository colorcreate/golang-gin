package remote

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func init() {
	orm.SetTableNameInflector(func(s string) string {
		return "person"
	})
}

// PersonConnection is
type PersonConnection struct {
	db *pg.DB
}

// ConnectionManager is
func (r *PersonConnection) ConnectionManager() {
	r.db = pg.Connect(&pg.Options{
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASS"),
		Database: os.Getenv("PGDB"),
	})

	err := r.db.CreateTable((*Person)(nil), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (r *PersonConnection) createPerson(person *Person) error {
	var err error
	if person.ID != 0 {
		err = r.db.Update(person)
	} else {
		err = r.db.Insert(person)
	}
	return err
}

func (r *PersonConnection) getAllPerson() []Person {
	var persons []Person
	err := r.db.Model(&persons).Select()
	if err != nil {
		fmt.Println(err.Error())
	}
	return persons
}

func (r *PersonConnection) getPersonByID(ID uint) (Person, error) {
	var person Person
	err := r.db.Model(&person).Where("id = ?", ID).Select()
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *PersonConnection) deletePersonByID(ID uint) error {
	var person Person
	_, err := r.db.Model(&person).Where("id = ?", ID).Delete()
	return err
}
