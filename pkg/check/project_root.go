package check

import (
	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsep/pkg/status"
	"github.com/Matt-Gleich/texsep/pkg/util"
)

// Check if the current directory is the project root
func ProjectRoot() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		statuser.Error("Failed to read current directory", err, 1)
	}

	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	if !util.Contains(fileNames, ".texsep.conf") {
		statuser.Warning("Project root not detected")
		var isRoot bool
		prompt := &survey.Confirm{
			Message: "Are you currently in the root of your project?",
		}
		err := survey.AskOne(prompt, &isRoot)
		if err != nil {
			statuser.Error("Failed to ask if current directory is project root", err, 1)
		}
		if isRoot {
			err := ioutil.WriteFile(".texsep.conf", []byte("DO NOT DELETE. USED FOR TEXSEP ROOT PROJECT DETECTION"), 0755)
			if err != nil {
				statuser.Error("Failed to create & write to .texsep.conf", err, 1)
			}
			status.Success("Created the .texsep.conf file for automatic project detection")
		} else {
			statuser.ErrorMsg("Please make sure you are at the project root then run again", 1)
		}
	} else {
		status.Success("Project root detected automatically")
	}
}
