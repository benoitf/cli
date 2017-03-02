package docker

import (
	"regexp"
	"strings"
	"os/exec"
	"errors"
)

// Version executer interface that is executing a command to get the output
// usually the $docker version output
type VersionExecuter interface {
	Execute() (string, error)
}

// struct for docker version executer
type DockerVersionExecuter struct {

}

// parsing of the result of the version executer
func parse(versionExecuter VersionExecuter) (string, error) {
	re := regexp.MustCompile("\n.*Version:(.*)\n")
	content, err := versionExecuter.Execute();
	matching := re.FindStringSubmatch(content);
	if len(matching) != 2 {
		return "", errors.New("No matching version in the resulting docker version output" + content)
	} else {
		return strings.TrimSpace(re.FindStringSubmatch(content)[1]), err
	}
}

// Default implementation of version executer
// based on $(docker version) output
func (DockerVersionExecuter) Execute() (string, error) {
	cmd := exec.Command(GetDockerBinary(), "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// internal method allowing to pickup the version executer
func getVersion(versionExecuter VersionExecuter) (string, error) {
	version, err := parse(versionExecuter);
	return version, err
}

// Use "docker version" for getting version
func GetVersion() (string, error) {
	dockerVersionExecuter := DockerVersionExecuter{}
	return getVersion(dockerVersionExecuter)
}
