.PHONY: build test init add design.txt command.txt clean

build:
	go build

test:
    ## utils, manager, then models
	cd internal/utils_test && go test -v && cd .. && go test -v && cd models && go test -v

ifeq ($(OS),Windows_NT)
clean:
	del /f LVCS.exe
else
clean:
	rm -f LVCS
endif
