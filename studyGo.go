package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	ID     int
	Name   string
	Tell   string
	Gender string
}

func main() {
	p := `{"ID":101,"Name":"Fah hahaa","Tell":"0989956241","Gender":"Female"}`

	var data person
	err := json.Unmarshal([]byte(p), &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID = ", data.ID, " Name =", data.Name, " Tell =", data.Tell, " Gender = ", data.Gender)
}
