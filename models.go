package main

import (
	"os/exec"
)

const dataDir = "/var/kiwix/data"
const libraryFile = dataDir + "/library.xml"

func listZimFiles() ([]string, error) {
	cmd := exec.Command("kiwix-manage", libraryFile, "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseZimFiles(output), nil
}

func addZimFile(zimFile string) error {
	cmd := exec.Command("kiwix-manage", libraryFile, "add", dataDir+"/"+zimFile)
	return cmd.Run()
}

func deleteZimFile(zimFile string) error {
	cmd := exec.Command("kiwix-manage", libraryFile, "remove", dataDir+"/"+zimFile)
	return cmd.Run()
}

func parseZimFiles(output []byte) []string {
	// Parse the output and return the list of .zim files
	// Placeholder for actual parsing logic
	return []string{}
}
