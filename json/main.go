package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("learning json")
	person := Person{Name: "fahad", Age: 23, IsAdult: true}
	fmt.Println(person)

	//convert to json marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("errror", err)
		return
	}
	fmt.Println("person json data", string(jsonData))

	//decodng //unmarshalling

	var persondata Person
	err = json.Unmarshal(jsonData, &persondata)
	if err != nil {
		fmt.Println("unmarshaling data err", err)
		return
	}
	fmt.Println("person data is", persondata)

}
