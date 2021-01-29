package widgets

import (
	"path/filepath"
	"regexp"
	"strings"
)

//getGroup returns a simplified process name based on the full command with arguments
func getGroup(fullCmd string) (group string) {
	// separate the executable from the arguments
	command := strings.SplitN(fullCmd, " ", 2)[0]

	if strings.HasPrefix(command, "/") {
		// for absolute paths, use the base name only
		group = filepath.Base(command)
	} else if strings.HasPrefix(command, "[") {
		// simplify the process names in brackets
		re := regexp.MustCompile(`\[|]|/.+$`)
		group = re.ReplaceAllString(command, "")
	} else {
		// strip the colon, that may remain in some executable names
		group = strings.TrimRight(command, ":")
	}

	return
}
