#! /usr/bin/make
ifeq ($(OS),Windows_NT)
	BUILD_TARGET_FILES = image-changer.exe main.go
else
	BUILD_TARGET_FILES ?= image-changer main.go
endif

all: cleandep depend build

prepare: cleandep depend

depend:
	@glide install

cleandep:
	@rm -rf vendor

build:
	@go build -o $(BUILD_TARGET_FILES)
