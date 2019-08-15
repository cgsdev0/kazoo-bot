package main

import (
	"encoding/json"
	"os"

	"samsaradev.io/team"
)

func main() {
	result, err := json.Marshal(team.AllTeams)
	if err == nil {
		os.Stdout.Write(result)
	}
}
