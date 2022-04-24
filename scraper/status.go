package scraper

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const watchedFile = "watched.txt"

func UpdateEpisode(episode int) {
	file, err := os.OpenFile(watchedFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
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
	watchedEp, err := strconv.Atoi(readFileLines(watchedFile)[0])
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

func readFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("\nCouldn't open file '%s'. Error:\n%s\n", filename, err)
		return nil
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

	return content
}
