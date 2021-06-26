package main

import (
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
	fmt.Println("JSON - Unmarshal")
	fmt.Println("-----------------------------------------------")
	// Convertendo um JSON em Struct
	cachorroJSON := `{"nome":"Toddy","raca":"York","idade":7}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	fmt.Println("-----------------------------------------------")
	// Convertendo um JSON em Map
	cachorro2JSON := `{"nome":"Fred","raca":"Maltes"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2JSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
