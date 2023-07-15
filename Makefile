.PHONY: build run

build:
	go build -o build main.go

test:
	cd internal/utils_test && go test -v

# ## without argument
# run:
# 	./build

## with argument
run:
	./build $(filter-out $@,$(MAKECMDGOALS))
%:
	@: