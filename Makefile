.PHONY: build run

build:
	go build -o build main.go

test:
	cd helper && go test

# ## without argument
# run:
# 	./build

## with argument
run:
	./build $(filter-out $@,$(MAKECMDGOALS))
%:
	@: