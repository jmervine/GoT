# tests without -tabs for go tip
travis: get .PHONY
	# Run Test Suite
	go test -test.v=true

test: format .PHONY
	# Run Test Suite
	-go test -test.v=true

functional: .PHONY
	cd _example; go test -test.v

build: test .PHONY
	go build -o 'pkg/got' -v -a -race

readme: test
	# generating readme
	godoc -ex -v -templates "$(PWD)/_support" . > README.md

format: .PHONY
	# Gofmt Source
	gofmt -tabs=false -tabwidth=4 -w=true -l=true *.go

.PHONY:

