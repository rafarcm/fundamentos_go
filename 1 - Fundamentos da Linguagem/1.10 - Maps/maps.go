package main

import "fmt"

func main() {
	fmt.Println("Maps")

	fmt.Println("---------------------------------------------")
	usuario1 := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])

	usuario1["profissao"] = "Engenheiro"
	fmt.Println(usuario1)

	fmt.Println("---------------------------------------------")
	usuario2 := map[int]string{
		1: "Pedro",
		2: "Silva",
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2[2])

	usuario2[50] = "Médico"
	fmt.Println(usuario2)

	delete(usuario2, 2)
	fmt.Println(usuario2)

	fmt.Println("---------------------------------------------")
	// Map aninhado
	usuario3 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Rafael",
			"sobrenome": "Moraes",
		},
		"curso": {
			"nome":      "Engenhria",
			"faculdade": "USP",
		},
	}

	fmt.Println(usuario3)

}
