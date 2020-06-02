package check

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Matt-Gleich/texsep/status"
	"github.com/Matt-Gleich/texsep/util"
	"github.com/fatih/color"
)

// ProjectRoot ... Check if the current directory is the project root
func ProjectRoot() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	if !util.Contains(fileNames, ".texsep.conf") {
		status.Step("Project root not detected\nIs thi the project root? (y or n)")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if input == "n\n" {
			fmt.Println("Please make sure you are at the project root then run again")
			os.Exit(0)
		}
		if input == "y\n" {
			err := ioutil.WriteFile(".texsep.conf", []byte("DO NOT DELETE. USED FOR texsep ROOT PROJECT DETECTION"), 0755)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Created .texsep file for automatic project detection")
		}
	} else {
		color.Green("Project root detected automatically")
	}
}
