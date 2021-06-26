package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Struct com campos mapeados para JSON
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	fmt.Println("JSON - Marshal")
	fmt.Println("-----------------------------------------------")
	// Convertendo um Struct em JSON
	c := cachorro{"Toddy", "York", 7}
	fmt.Println(c)

	cachorroJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJSON)                  // Slice de bytes
	fmt.Println(bytes.NewBuffer(cachorroJSON)) // Converte um slice de bytes em algo legível como um JSON

	fmt.Println("-----------------------------------------------")
	// Convertendo um Map em JSON
	c2 := map[string]string{
		"nome": "Fred",
		"raca": "Maltes",
	}
	cachorro2JSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2JSON) // Slice de bytes
	fmt.Println(bytes.NewBuffer(cachorro2JSON))
}
