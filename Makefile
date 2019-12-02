default:
	@echo "Nothing to do"

test:
	ragel -s -Z -G2 -o machine.go machine.go.rl
	go test
