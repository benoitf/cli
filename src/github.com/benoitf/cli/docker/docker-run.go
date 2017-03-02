package docker

import (
	"os/exec"
	"fmt"
	"os"
)



// Defines the DockerRunCommand structure
// 	Options are the options to docker run command like volume mounting
// 	ImageName is the docker image to execute
// 	Command is an optional command to execute on the image
//	Args are the arguments to the provided docker run command
type DockerRunCommand struct {
	Options []string
	ImageName string
	Command string
	Args []string
}


// Executes the given docker command by streaming the output to stdout
func Exec(dockerRunCommand *DockerRunCommand) error {

	dockerArgs :=[]string{}
	// it's a run command
	dockerArgs = append(dockerArgs, "run")

	//TODO: only if run inside a terminal ?
	dockerArgs = append(dockerArgs, "-ti")

	// cleanup
	dockerArgs = append(dockerArgs, "--rm")


	// always mount the docker.socket for now
	dockerArgs = append(dockerArgs, "-v", "/var/run/docker.sock:/var/run/docker.sock")

	if len(dockerRunCommand.Options) > 0 {
		dockerArgs = append(dockerArgs, dockerRunCommand.Options...)
	}
	dockerArgs = append(dockerArgs, dockerRunCommand.ImageName)
	if len (dockerRunCommand.Command) > 0 {
		dockerArgs = append(dockerArgs, dockerRunCommand.Command)
	}
	if len(dockerRunCommand.Args) > 0 {
		dockerArgs = append(dockerArgs, dockerRunCommand.Args...)
	}

	// Create docker command
	cmd := exec.Command(GetDockerBinary(), dockerArgs...)

	// pipe the output/error streams
	_, err := cmd.StdoutPipe()
	/*	_, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return err
	}*/

	/*scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()*/
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}

	/*slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)*/
	err = cmd.Wait()
	return nil
}