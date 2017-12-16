SHELL := /bin/bash

TARGET := main
VERSION := 0.0.1
BUILD := `git rev-parse HEAD`

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)

.PHONY: all build clean install fmt uninstall simplify check run

all: check install

$(TARGET): $(SRC)
				@go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
				@true

clean: 
				@rm -f $(TARGET)

