build:
	@go build -o ./bin/pokedex .
	
run: build
	@./bin/pokedex