#
# A simple Makefile containing basic targets to build and run the program
#
# NOTE: Run make depends to get the testing package.
#

build:
	go build -v picoyplaca.go

clean:
	go clean -x

test:
	go test -cover -check.f "picoyplacaSuite.*" -check.vv .

depends:
	go get -u -v gopkg.in/check.v1
