go-ethereum-wallet:
	env GO111MODULE=on go build

clean:
	rm go-ethereum-wallet

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	go-ethereum-wallet \
	clean \
	test \
	lint