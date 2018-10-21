package main

import (
	"log"
	"time"

	"github.com/spf13/afero"
)

func removeDirectories(filesystem afero.Fs, path string, timespan time.Duration, keep int) {

}

func getCandidateDirectories(filesystem afero.Fs, path string, duration time.Duration, keep int) []string {
	var candidateDirs []string
	now := time.Now()
	directory, err := afero.ReadDir(filesystem, path)
	checkErr(err)

	for _, file := range directory {
		if age := now.Sub(file.ModTime()); age > duration {
			candidateDirs = append(candidateDirs, file.Name())
		}
	}

	return candidateDirs
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {

}
