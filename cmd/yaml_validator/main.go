package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func main() {

	fmt.Println(`
*****************************************
* YAML-schema validation tool           *
* based on GoLang 1.19.5                *
* maintained by: vitvasilyuk@gmail.com  *
*****************************************`)

	schemaPath := flag.String("schema", "", "a string path")
	featureDir := flag.String("dir", "", "a string path")
	yamlFiles := flag.String("yaml", "", "a string path")

	flag.Parse()

	fmt.Println("\nPaths passed:\n- feature_dir:", *featureDir)
	fmt.Println("- schema:", *schemaPath)
	fmt.Println("- yaml:", *yamlFiles)
	fmt.Println()

	hasErrors := false
	fileNames := strings.Split(*yamlFiles, ",")

	for _, name := range fileNames {

		var document map[string]interface{}

		data, err := os.ReadFile(*featureDir + "/" + name)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		if err := yaml.Unmarshal(data, &document); err != nil {
			log.Fatalf("unable to parse yaml: %v", err)
		}

		schemaLoader := gojsonschema.NewReferenceLoader("file://" + *schemaPath)
		documentLoader := gojsonschema.NewGoLoader(document)
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)

		if err != nil {
			panic(err.Error())
		}

		if !result.Valid() {
			hasErrors = true
			fmt.Printf("%s is failed, see errors :\n", name)
			for _, desc := range result.Errors() {
				fmt.Printf("    - %s\n", desc)
			}
		} else {
			fmt.Printf("%s is valid\n", name)
		}
	}
	if hasErrors {
		fmt.Printf("\nVALIDATION FAILED !!!\n\n")
		return
	}
	fmt.Printf("\nVALIDATION SUCCESSFUL !!!\n\n")
}
