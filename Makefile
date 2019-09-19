start:
	GO111MODULE=on go run main.go

vet:
	GO111MODULE=on go vet -v ./...

build: vet
	GO111MODULE=on go build -o ${PWD}/bin/vayne ./

test:
	GO111MODULE=on go test -v ./...