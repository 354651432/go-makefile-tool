SRC=$(shell find . -name "*.go")

main: $(SRC)
	@go build -o exmake
