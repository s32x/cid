init:
	-rm -rf ./vendor go.mod go.sum
	GO111MODULE=on go mod init

deps:
	-rm -rf ./vendor go.sum
	GO111MODULE=on go mod vendor
	
test:
	go test ./...

prune:
	up prune -r 5

deploy: deps test
	-up stack plan
	-up stack apply
	up deploy production