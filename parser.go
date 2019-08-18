package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"strings"
	"encoding/json"
)

type Feature struct {
	Name string `json:"name"`
	Scenarios []string `json:"scenarios"`
}

func main() {
	feature := Feature{}

	body, err := ioutil.ReadFile("Guess.feature")
	if err != nil {
		log.Fatalf("unable to load file: %v", err)
	}

	// Split string by each of its lines
	lines := strings.Split(string(body), "\n")

	for index, line := range lines {
		// Extract "Feature"
		prefix := "Feature:"
		if strings.HasPrefix(line, prefix) {
			feature.Name = strings.TrimSpace(line[len(prefix):])
			//result["feature"] = feature
		}

		// Get scenarios
		prefix = "Scenario:"
		if strings.HasPrefix(strings.TrimSpace(line), prefix) {
			fmt.Println("scenario prefix found")
			scenario := strings.TrimSpace(line[(strings.Index(line,prefix)+len(prefix)):])

			for i := index+1; ; i++ {
				fmt.Println(i)
				
				// Identifying the end of the scenario
				if i >= len(lines) || len(lines[i]) == 0 {
					break;
				} else {
					fmt.Println("appending", strings.TrimSpace(lines[i]))
					scenario = fmt.Sprintf("%s\n%s ",scenario, strings.TrimSpace(lines[i]))
				}
			}
			feature.Scenarios = append(feature.Scenarios, scenario)
		}
	}
	resultJSON, err := json.Marshal(feature)

	if err != nil {
		fmt.Println("error while parsing to json :v",err)
	}

	fmt.Println(string(resultJSON))
}