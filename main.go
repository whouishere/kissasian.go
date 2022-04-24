package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/whouishere/kissasiandb/scraper"
)

var version string = "dev-snapshot"

func main() {
	fmt.Printf("Version %s\n\n", version)

	scraper.ConnectToEpisodeList()

	fmt.Print("\nPress any key to exit.")
	bufio.NewScanner(os.Stdin).Scan()
}
