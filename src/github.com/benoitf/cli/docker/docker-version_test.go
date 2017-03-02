package docker

import "testing"

// result of the command
// 	$ docker version
// used as mock data
var dummyContent = `Client:
 Version:      1.11.0
 API version:  1.26
 Go version:   go1.7.5
 Git commit:   092cba3
 Built:        Wed Feb  8 08:47:51 2017
 OS/Arch:      darwin/amd64

Server:
 Version:      1.10.0
 API version:  1.26 (minimum version 1.12)
 Go version:   go1.7.5
 Git commit:   092cba3
 Built:        Wed Feb  8 08:47:51 2017
 OS/Arch:      linux/amd64
 Experimental: true`


var dummyInvalidContent = `No version output`

// dummy implementation
type FooDockerVersionExecuter struct {
	Content string

}
// implementation of the execute method for the test
func (fooDockerVersionExecuter FooDockerVersionExecuter)  Execute() (string, error) {
	return fooDockerVersionExecuter.Content, nil
}

// Check that parsing version if working correctly
func TestGetVersionValidOutput(t *testing.T) {
	version, err := getVersion(FooDockerVersionExecuter{Content: dummyContent})
	if version != "1.11.0" {
		t.Error("Invalid version")
	}
	if (err != nil) {
		t.Error("No error should be reported")
	}

}


// Check that parsing version if working correctly
func TestGetVersionInvalidOutput(t *testing.T) {
	version, err := getVersion(FooDockerVersionExecuter{Content: dummyInvalidContent})
	if (err == nil) {
		t.Error("An error should be reported")
	}
	if version != "" {
		t.Error("No version can't be returned")
	}

}