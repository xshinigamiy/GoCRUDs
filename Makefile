#!make
include config/commons.conf
export $(shell sed 's/=.*//' config/commons.conf)
include config/dev.conf
export $(shell sed 's/=.*//' config/dev.conf)

build:
	go build -o app cmd/GoCRUDs/main.go

run:
	go run cmd/GoCRUDs/main.go

dev:
	reflex -s -r '\.go$$' -R '(vendor)\/*' -d fancy make run

prod:
	include config/prod.conf
	export $(shell sed 's/=.*//' config/prod.conf)
	reflex -s -r '\.go$$' -R '(vendor)\/*' -d fancy make run