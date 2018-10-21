package main

import (
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func mockFilesystem() afero.Fs {
	fs := afero.NewMemMapFs()
	fs.Mkdir("today", 0644)
	fs.Mkdir("lastweek", 0644)
	fs.Mkdir("lastmonth", 0644)
	fs.Mkdir("lastyear", 0644)

	day := 24 * time.Hour
	today := time.Now()
	lastweek := today.Add(-7 * day)
	lastmonth := today.Add(-30 * day)
	lastyear := today.Add(-365 * day)

	fs.Chtimes("today", today, today)
	fs.Chtimes("lastweek", lastweek, lastweek)
	fs.Chtimes("lastmonth", lastmonth, lastmonth)
	fs.Chtimes("lastyear", lastyear, lastyear)

	return fs
}

func assertDirExists(t *testing.T, fs afero.Fs, dir string) {
	exists, _ := afero.DirExists(fs, dir)
	assert.True(t, exists, "The directory does not exist")
}

func assertDirNotExists(t *testing.T, fs afero.Fs, dir string) {
	exists, _ := afero.DirExists(fs, dir)
	assert.False(t, exists, "The directory exists")
}

func TestGetCandidateDirectoriesOlderThanTwoWeeks(t *testing.T) {
	fs := mockFilesystem()

	candidateDirs := getCandidateDirectories(fs, "/", time.Hour*24*14, 0)
	assert.Equal(t, []string{"lastmonth", "lastyear"}, candidateDirs)
}

func TestGetCandidateDirectoriesOlderThanTwoWeeksButKeepTheLastThree(t *testing.T) {
	fs := mockFilesystem()

	candidateDirs := getCandidateDirectories(fs, "/", time.Hour*24*14, 3)
	assert.Equal(t, []string{"lastmonth", "lastyear"}, candidateDirs)
}

// func TestDeletesDirsOlderThanTwoWeeks(t *testing.T) {
// 	fs := mockFilesystem()

// 	timespan, _ := time.ParseDuration("14d")

// 	removeDirectories(fs, "/", timespan)

// 	assertDirExists(t, fs, "/today")
// 	assertDirExists(t, fs, "/lastweek")
// 	assertDirNotExists(t, fs, "/lastmonth")
// 	assertDirNotExists(t, fs, "/lastyear")
// }
