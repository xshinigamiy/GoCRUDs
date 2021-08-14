#!make
export $(shell cat config/commons.conf > .env ||  cat config/dev.conf > .env)
export $(shell echo "exporting env variables from .env")
include .env
export $(shell sed 's/=.*//' .env)

setup-env:
	include .env
	export $(shell sed 's/=.*//' .env)

build:
	go build -o app cmd/GoCRUDs/main.go

run:
	go run cmd/GoCRUDs/main.go

dev:
	reflex -s -r '\.go$$' -R '(vendor)\/*' -d fancy make run
