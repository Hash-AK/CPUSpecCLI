package main

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed cpus.json
var DB embed.FS

type CPUs struct {
	CPUs []CPU `json:"cpus"`
}
type CPU struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Brand              string `json:"brand"`
	Generation         int    `json:"generation"`
	GenerationCodename string `json:"generation_codename"`
	Series             string `json:"series"`
}

func main() {
	dbJson, err := DB.ReadFile("cpus.json")
	if err != nil {
		panic(err)
	}
	var cpus CPUs
	json.Unmarshal([]byte(dbJson), &cpus)
	for i := range cpus.CPUs {
		fmt.Println(cpus.CPUs[i].ID)
		fmt.Println(cpus.CPUs[i].Name)
		fmt.Println(cpus.CPUs[i].Brand)
		fmt.Println(cpus.CPUs[i].Generation)
		fmt.Println(cpus.CPUs[i].GenerationCodename)
		fmt.Println(cpus.CPUs[i].Series)

	}

}
