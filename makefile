SHELL := /bin/bash

# Path: makefile
# Name: makefile
# Description: Makefile for the project

run:
	go run cmd/main.go

build:
	go build -o bin/main cmd/main.go