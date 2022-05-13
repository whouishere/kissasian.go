package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/whouishere/kissasiandb/scraper"
)

var Version string = "DEV"

func main() {
	fmt.Printf("Version %s\n\n", Version)

	scraper.ConnectToEpisodeList()

	fmt.Print("\nPress any key to exit.")
	bufio.NewScanner(os.Stdin).Scan()
}
