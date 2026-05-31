.PHONY: generate build run

generate:
	oapi-codegen -package public -generate types,server spec/public.yaml > gen/public/public.gen.go
	oapi-codegen -package protected -generate types,server spec/protected.yaml > gen/protected/protected.gen.go

build:
	go build -o server ./cmd/api

run:
	go run ./cmd/api
