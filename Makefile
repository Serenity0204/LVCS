.PHONY: build run

build:
	go build -o build main.go

# ## without argument
# run:
# 	./build

## with argument
run:
	./build $(filter-out $@,$(MAKECMDGOALS))
%:
	@: