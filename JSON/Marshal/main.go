package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	d := dog{"boltz", "labrador", 4}

	dogInJSON, erro := json.Marshal(d)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dogInJSON)                  //bityes
	fmt.Println(bytes.NewBuffer(dogInJSON)) //json

	d2 := map[string]string{
		"name":  "mel",
		"breed": "poodle",
		"age":   "9",
	}

	dog2InJSON, erro := json.Marshal(d2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dog2InJSON)                  //bityes
	fmt.Println(bytes.NewBuffer(dog2InJSON)) //json
}
