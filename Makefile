default:
		go build

test:
		/usr/local/go/bin/go test -timeout 30s -coverprofile=coverage.txt -covermode=atomic ./... 

.phony: test