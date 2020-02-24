package main

import "errors"

// create struct
type Person struct {
	id int
	firstname string
	lastname string
}

// Create register a person

func PersonCreate(p Person) error{
	q := `INSERT INTO 
			people (id, firstname, lastname) 
			VALUES ($1 , $2 , $3 )`
	db := getConnection()
	defer db.Close()

	stmt , err := db.Prepare(q)
	if err != nil {
		return  err
	}
	defer stmt.Close()

	r, err := stmt.Exec(p.id, p.firstname, p.lastname)
	if err != nil {
		return err
	}

	i , _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba una fila afectada")
	}

	return nil
}