package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
)

type SqlcJson struct {
	Version  string     `json:"version"`
	Packages []Packages `json:"packages"`
}

type Packages struct {
	Name                   string   `json:"name"`
	Path                   string   `json:"path"`
	Queries                []string `json:"queries"`
	Schema                 []string `json:"schema"`
	Engine                 string   `json:"engine"`
	Emit_json_tags         bool     `json:"emit_json_tags"`
	Emit_prepared_queries  bool     `json:"emit_prepared_queries"`
	Emit_interface         bool     `json:"emit_interface"`
	Emit_exact_table_names bool     `json:"emit_exact_table_names"`
	Emit_empty_slices      bool     `json:"emit_empty_slices"`
}

func generateModels() []string {
	var models []string
	for _, dbTables := range mysqldata.AppliDbTables {
		tablePath := strings.ReplaceAll(dbTables, "_", "/")
		models = append(models, "./"+tablePath)
	}
	return models
}

func main() {
	models := generateModels()
	sqlcJsonData := SqlcJson{
		Version: "1",
		Packages: []Packages{
			{Name: "mysql",
				Path:                  "../datasources/mysql",
				Queries:               models,
				Schema:                models,
				Engine:                "mysql",
				Emit_json_tags:        true,
				Emit_prepared_queries: true,
				Emit_interface:        true,
				Emit_empty_slices:     true,
			},
		},
	}
	jsonData, err := json.Marshal(sqlcJsonData)
	if err != nil {
		fmt.Println("Error while encoding json data :", err)
		return
	}

	file, err := os.Create("sqlc.json")
	if err != nil {
		fmt.Println("Erreur while creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error while writing json data in file :", err)
		return
	}

	cmd := exec.Command("sqlc", "generate")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error while executing sqlc generate :", err)
		return
	}
	fmt.Println(string(out))
	fmt.Println("Models generated successfully !")
	fmt.Println("Cleaning models...")
	cleanModels()
}

func cleanModels() {
	script := "cleanModels.py"

	cmd := exec.Command("python", script)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
