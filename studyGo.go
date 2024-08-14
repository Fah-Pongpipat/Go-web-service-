package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	ID     int
	Name   string
	Tell   string
	Gender string
}

func main() {
	data, error := json.Marshal(&person{ID: 001, Name: "Pongpipat siangsai", Tell: "0661320025", Gender: "Male"})
	fmt.Println(string(data), "Error :", error)
}
