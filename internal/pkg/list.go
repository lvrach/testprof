package pkg

import (
	"os/exec"
	"strings"
)

func List() ([]string, error) {
	raw, err := exec.Command("go", "list", "-f", "{{.ImportPath}}", "./...").Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.Trim(string(raw), "\n"), "\n"), nil
}
