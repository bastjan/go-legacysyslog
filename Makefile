build: build-protocolchooser
	ragel -s -Z -G2 -o machine.go machine.go.rl
	go fmt machine.go

build-protocolchooser:
	cd utils/protocolchooser && ragel -s -Z -G2 -o protocolchooser.go protocolchooser.go.rl
	cd utils/protocolchooser && go fmt protocolchooser.go

test: build
	go test ./...

graph: graph-protocolchooser
	ragel -ZVps machine.go.rl -o machine.dot
	dot -Tsvg machine.dot -o machine.svg

graph-protocolchooser:
	cd utils/protocolchooser && ragel -ZVps protocolchooser.go.rl -o protocolchooser.dot
	cd utils/protocolchooser && dot -Tsvg protocolchooser.dot -o protocolchooser.svg
