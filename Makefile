.PHONY: cpp go python FORCE

cpp:
	@g++ -std=c++17 "$(word 2,$(MAKECMDGOALS))/cpp/main.cpp" -o /tmp/lld_cpp
	@/tmp/lld_cpp

go:
	@go run "$(word 2,$(MAKECMDGOALS))/go/main.go"

python:
	@uv run python "$(word 2,$(MAKECMDGOALS))/python/main.py"

%: FORCE
	@:

FORCE:

# .PHONY: python cpp go

# python:
# 	uv run python 01-design-patterns/creational/factory/python/main.py

# cpp:
# 	g++ -std=c++17 01-design-patterns/creational/factory/cpp/main.cpp -o /tmp/lld_cpp
# 	/tmp/lld_cpp

# go:
# 	go run ./01-design-patterns/creational/factory/go