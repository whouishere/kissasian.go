GC := go build
CFLAGS := -v 
LDFLAGS := 

BINARY := kissasiandb
BIN_DIR := bin

ifeq ($(OS),Windows_NT)
	BINARY+=.exe
endif

VERSION := 0.1

LDFLAGS += -X 'main.version=$(VERSION)'

SRC := main.go

.PHONY: build run test clean cleandata

all: build test

build: 
	$(GC) $(CFLAGS) -ldflags="$(LDFLAGS)" -o $@ $<

run: build
	./$(BIN_DIR)/$(BINARY)

test:
	go test -v $(SRC)

clean: 
	rm -f $(BIN_DIR)/$(BINARY)

cleandata: 
	rm -f $(BIN_DIR)/watched.txt