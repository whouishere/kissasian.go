# kissasiandb

A simple webscraper for [kissasiandb](https://kissasiandb.com), meant for educational purposes.

For now the program will only look for a hard-coded show on the site. In which you can easily change on `scraper/scraper.go`.

## Building
If you want to build this the easy way, I recommend you to use [Task](https://github.com/go-task/task)It's a simple tool and it integrates very well with Go.

To build the project with it, just be sure to be inside the project directory an run

```
task build
```

and a `kissasiandb` binary will be built on the `bin` folder.

### Using Go directly
If you don't have or don't want to use Task, you can use the Go command-line it directly.

Run it simply as

```
go build -o main.go bin/kissasiandb
```

__DISCLAIMER__: If you're on Windows, you need to insert `.exe` at the end of command in order to execute it.