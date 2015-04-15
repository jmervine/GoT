# tests without -tabs for go tip
travis: .PHONY
	# Run Test Suite
	go test -test.v

test: .PHONY
	# Run Test Suite
	go test -test.v

quiet/test: .PHONY
	go test

functional: .PHONY
	cd _example; go test -test.v

quiet/functional: .PHONY
	cd _example; go test

build: quiet/test quiet/functional .PHONY
	go build -o 'pkg/got' -v -a -race

readme: quiet/test quiet/functional .PHONY
	# generating readme
	godoc -ex -v -templates "$(PWD)/_support" . > README.md

.PHONY:

