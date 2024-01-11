package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type update struct {
	Status string
	Class  string
}

func main() {
	// Example of single command and piped command.
	strArr := [2]string{"<command to execute>", "<command> | <command>"}

	var rtn string
	for _, str := range strArr {
		stdout, err := exec.Command("bash", "-c", str).Output()
		if err != nil {
			log.Fatal(err)
		}

		rtn += string(stdout)

	}

	json, err := json.Marshal(rtn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}
