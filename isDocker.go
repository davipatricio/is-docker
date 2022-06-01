package isdocker

import (
	"io/ioutil"
	"os"
	"strings"
)

var IsDockerCached bool

func DockerEnvExists() bool {
	// Check if the file ./.docker exists
	_, err := os.Stat("./.dockerenv")
	return err == nil
}

func HasDockerCGroup() bool {
	// Read the file /proc/self/cgroup
	content, err := ioutil.ReadFile("/proc/self/cgroup")
	if err != nil {
		return false
	}

	// Check if the content contains the string "docker"
	return strings.Contains(string(content), "docker")

}

func Check() bool {
	if IsDockerCached {
		return IsDockerCached
	}

	if DockerEnvExists() {
		IsDockerCached = true
		return true
	}

	if HasDockerCGroup() {
		IsDockerCached = true
		return true
	}

	return false
}
