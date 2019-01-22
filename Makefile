gofmt :
	$(info Reformatting all source files...)
	go fmt ./...

build :
	go build

test :
	mkdir -p coverage
	go test -covermode=count -coverprofile=coverage/count.out ./...

coverage : test
	go tool cover -html=coverage/count.out
