build:
	@go build -o ./bin/pokedex ./main.go

run: build
	@./bin/pokedex