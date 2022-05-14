package scraper

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const statusFile = "kissasiandb.status"

var statusFilePath string

func UpdateEpisode(episode int) {
	file, fileErr := os.OpenFile(statusFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()

	_, popErr := popLine(file)
	if popErr != nil {
		log.Fatal(popErr)
	}

	_, writeErr := file.WriteString(fmt.Sprint(episode))
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}

func GetWatchedEpisode() int {
	statusFilePath = getFileDir()
	epStr, err := readFileLines(statusFilePath)

	// if previously there wasn't a watched cache and it was just created,
	// the file will return nil, which then we need to put something first
	if err != nil {
		log.Fatal(err)
	} else if epStr == nil {
		UpdateEpisode(0)
		return 0
	}

	watchedEp, err := strconv.Atoi(epStr[0])
	if err != nil {
		log.Fatal(err)
	}
	return watchedEp
}

// straight up stolen from StackOverflow lol didn't even modified it
func popLine(f *os.File) ([]byte, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	nw, err := io.Copy(f, buf)
	if err != nil {
		return nil, err
	}
	err = f.Truncate(nw)
	if err != nil {
		return nil, err
	}
	err = f.Sync()
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return line, nil
}

func readFileLines(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content []string

	// create a file scanner
	scanner := bufio.NewScanner(file)
	// loop through it
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("\n'%s' scanner got an error. Error:\n%s\n", filename, err)
	}

	return content, err
}

// checks where does status file exists
func getFileDir() string {
	cwDir, err := os.Getwd() // current working directory
	if err != nil {
		log.Fatal(err)
	}
	homeDir, err := os.UserHomeDir() // user home directory
	if err != nil {
		log.Fatal(err)
	}
	cacheDir, err := os.UserCacheDir() // user cache directory
	if err != nil {
		log.Fatal(err)
	}
	execDir, err := filepath.Abs(filepath.Dir(os.Args[0])) // executable directory
	if err != nil {
		log.Fatal(err)
	}

	possibleDirs := []string{
		cwDir,
		homeDir,
		cacheDir,
		execDir,
	}

	var absoluteFilePath string

	// search for the status file on possibleDirs directories (in order)
	// first one to be found is returned
	for _, dir := range possibleDirs {
		_, err := os.Stat(filepath.Join(dir, statusFile))
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				log.Fatal(err)
			}
		}
		absoluteFilePath = filepath.Join(dir, statusFile)
	}

	return absoluteFilePath
}
