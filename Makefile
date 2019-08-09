#
# A simple Makefile containing basic targets to build and run the program
#
# NOTE: Run make depends to get the testing package.
#

default:
	@make help

build:
	go build -v picoyplaca.go

clean:
	go clean -x

test:
	go test -cover -check.f "picoyplacaSuite.*" -check.vv .

depends:
	go get -u -v gopkg.in/check.v1

help:
	@echo "------------------ How to use this Makefile ------------------"
	@echo "make build   - Builds the executable."
	@echo "make clean   - Cleans the work directory."
	@echo "make depends - Downloads the check.v1 dependency framework"
	@echo "               needed for running unit tests."
	@echo "make help    - Show this help text."
	@echo "make test    - Runs the Unit tests and shows code coverage."
	@echo "--------------------------------------------------------------"
