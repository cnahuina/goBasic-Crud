package main

import (
	"fmt"
	"log"
)

func main (){
	p := Person{
		id: 1,
		firstname: "Carlos",
		lastname: "Ñahuiña",
	}

	err := PersonCreate(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creado exitosamente")
}