default:
	@echo "Nothing to do"

build:
	ragel -s -Z -G2 -o machine.go machine.go.rl
	go fmt machine.go

test: build
	go test

graph:
	ragel -ZVps machine.go.rl -o machine.dot
	dot -Tsvg machine.dot -o machine.svg
	google-chrome machine.svg
