
// The docker package manages all the docker related stuff like invoking commands
package docker

import (
	"os/exec"
)


// Checks if docker is installed or not in the current user PATH
func IsInstalled() (bool, error) {
	_, err := exec.Command("docker").CombinedOutput()
	if err != nil {
		return false, err
	}
	return true, nil
}

