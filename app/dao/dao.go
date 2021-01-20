package dao

import (
	"database/sql"
	"fmt"
	"webapi/app/entity"
)

//StudentDAO is ..
type StudentDAO struct {
	Name string
}

//GetAll is ...
func (d *StudentDAO) GetAll() []entity.Student {
	var students []entity.Student
	students = append(students, entity.Student{Name: "demo", Age: 10})
	db, _ := sql.Open("sqlite3", d.Name)
	rows, _ := db.Query("SELECT Name, Age FROM Student")
	for rows.Next() {
		student := entity.Student{Name: "", Age: 0}
		rows.Scan(&student.Name, &student.Age)
		students = append(students, student)
	}
	return students
}

//Get is ..
func (d *StudentDAO) Get(id int) entity.Student {
	db, _ := sql.Open("sqlite3", d.Name)
	rows, _ := db.Query("SELECT Name, Age FROM Student WHERE id = ?", id)
	student := entity.Student{Name: "", Age: 0}
	for rows.Next() {
		rows.Scan(&student.Name, &student.Age)
		return student
	}
	rows.Close()
	db.Close()
	return student
}

//Add is ..
func (d *StudentDAO) Add(student entity.Student) error {
	query := fmt.Sprintf("INSERT INTO Student (Name, Age) VALUES (?, ?)")
	db, err := sql.Open("sqlite3", d.Name)
	statement, err := db.Prepare(query)
	_, err = statement.Exec(student.Name, student.Age)
	statement.Close()
	db.Close()
	return err
}

//Del is ..
func (d *StudentDAO) Del(id int) error {
	query := fmt.Sprintf("DELETE FROM Student WHERE id = ?")
	db, err := sql.Open("sqlite3", d.Name)
	statement, err := db.Prepare(query)
	_, err = statement.Exec(id)
	statement.Close()
	db.Close()
	return err
}
