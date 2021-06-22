build: talking-clock-cli talking-clock-server

talking-clock-cli:
	go build -o ./talking-clock-cli ./cli/main.go
talking-clock-server:
	go build -o ./talking-clock-server ./server/main.go

test:
	go test ./clock

clean:
	rm -f ./talking-clock-cli
	rm -f ./talking-clock-server
