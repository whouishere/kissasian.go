# kissasiandb

A simple webscraper for [kissasiandb](https://kissasiandb.com), meant for educational purposes.

For now the program will only look for a hard-coded show on the site. In which you can easily change on `scraper/scraper.go`.

## Building
### Using GNU's `make`
You can do a simple build with  `make` using

```
make build
```

and a `kissasiandb` binary will be built on the `bin` folder.

### Using Go directly
If you don't have `make`, or it simply doesn't work for you, you may use the Go command-line it directly.

Run it simply as

```
go build -o main.go bin/kissasiandb
```

__DISCLAIMER__: If you're on Windows, you need to insert `.exe` at the end of command in order to execute it.