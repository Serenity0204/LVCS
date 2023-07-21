.PHONY: build test clean

build:
	go build

test:
    ## utils, manager, then models
	cd internal/utils && go test -v && cd .. && go test -v && cd models && go test -v

ifeq ($(OS),Windows_NT)
clean:
	del /f LVCS.exe && rd /s /q ".lvcs" && rd /s /q ".lvcs_test"
else
clean:
	rm -f LVCS && rm -rf ".lvcs" && rm -rf ".lvcs_test"
endif
