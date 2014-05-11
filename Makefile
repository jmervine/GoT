# tests without -tabs for go tip
travis: .PHONY
	# Run Test Suite
	go test -test.v=true

test: format .PHONY
	# Run Test Suite
	go test -test.v=true

quiet/test: .PHONY
	go test

functional: .PHONY
	cd _example; go test -test.v

quiet/functional: .PHONY
	cd _example; go test

build: format quiet/test quiet/functional .PHONY
	go build -o 'pkg/got' -v -a -race

readme: format quiet/test quiet/functional .PHONY
	# generating readme
	godoc -ex -v -templates "$(PWD)/_support" . > README.md

format: .PHONY
	# Gofmt Source
	gofmt -tabs=false -tabwidth=4 -w=true -l=true *.go

.PHONY:

