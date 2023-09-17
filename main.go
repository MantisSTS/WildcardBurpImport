package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type ScopeImport struct {
	Target struct {
		Scope struct {
			AdvancedMode bool          `json:"advanced_mode"`
			Exclude      []interface{} `json:"exclude,omitempty"`
			Include      []struct {
				Enabled  bool   `json:"enabled"`
				Host     string `json:"host"`
				Protocol string `json:"protocol"`
			} `json:"include"`
		} `json:"scope"`
	} `json:"target"`
}

func main() {
	input := flag.String("f", "", "The wildcard file to parse")
	output := flag.String("o", "", "The output file to write to")
	flag.Parse()

	if *input == "" || *output == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var scope ScopeImport
	scope.Target.Scope.AdvancedMode = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		wildcard := scanner.Text()
		wildcard = regexp.QuoteMeta(strings.TrimSpace(wildcard))

		scope.Target.Scope.Include = append(scope.Target.Scope.Include, struct {
			Enabled  bool   `json:"enabled"`
			Host     string `json:"host"`
			Protocol string `json:"protocol"`
		}{Enabled: true, Host: wildcard, Protocol: "any"})
	}

	b, err := json.MarshalIndent(scope, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(*output, b, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
