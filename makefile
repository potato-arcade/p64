all: install

install: 
	go install ./... 

test:
	go vet ./...
	golint ./...
	go test ./...

run: 
	$(HOME)/go/bin/potato64

generate:
	fyne bundle -package p64 -name font kongtext.ttf > bundle.go
