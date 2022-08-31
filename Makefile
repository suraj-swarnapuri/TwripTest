

.PHONY: clean
	clean:
		go clean -cache -testcache -modcache

.PHONY: gen
gen:
	# Auto-generate code
	protoc --proto_path=. --twirp_out=. --go_out=. rpc/twirptest/twirptest.proto

.PHONY: server
server:
	go run cmd/twirptest/main.go

.PHONY: client
client:
	go build -o bin/twirptestclient ./cmd/client/main.go