package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Humberto",
		"sobrenome": "Chiesi",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Humberto",
			"ultimo":   "Chiesi",
		},
		"curso": {
			"faculdade": "PUC",
			"curso":     "Ciencia da Computacao",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["nascimento"] = map[string]string{
		"dia": "12",
		"mes": "08",
		"ano": "2000",
	}
	fmt.Println(usuario2)
}
