package cli

import (
	"github.com/benoitf/cli/docker"
	"fmt"
	"os"
	"path"
	"flag"
)

// Host folder that will be transformed into container /data storage
var dataDirectory string

// initialize the paths, like current
func initPath() {
	// Current folder
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	currentPathFolder := path.Dir(ex) + "/.che"
	fmt.Println("No data folder given, using", currentPathFolder)
	flag.StringVar(&dataDirectory, "data", currentPathFolder, "Data directory of CLI")
}

func Execute() {

	version, err := docker.GetVersion()
	if (err != nil) {
		fmt.Fprintln(os.Stdout, err);
		os.Exit(1);
	}
	fmt.Fprintln(os.Stdout, "Native CLI. Found Docker Version", version);


	initPath()

	// set flags
	flag.Parse()
	// Give as arguments all arguments not parsed by flag package
	dockerRunComand := &docker.DockerRunCommand{Options: []string{"-v", dataDirectory + ":/data"}, ImageName : "eclipse/che-cli:nightly", Command:"", Args:flag.Args()}
	docker.Exec(dockerRunComand)

}