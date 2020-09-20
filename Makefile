#!make

.PHONY: deps

deps: 
	@echo Installing developer dependencies
	@GO111MODULE=off go get github.com/cosmtrek/air
