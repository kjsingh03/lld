# LLD

My hands-on **Low-Level Design & Interview Preparation** repository :)

Focused on:

* SOLID & OOP
* UML / Class Diagrams
* Design Patterns
* Real-world LLD Problems
* C++ / Go / Python implementations

## Structure

```text
01-design-patterns/
02-lld-problems/
03-advanced-lld/
04-interview-drills/
05-cheatsheets/
```

Each problem contains:

```text
problem/
├── diagram.drawio
├── README.md
├── cpp/
├── go/
└── python/
```

## Run

Pass the problem directory to Make:

```bash
make cpp ./01-design-patterns/creational/factory
make go ./01-design-patterns/creational/factory
make python ./01-design-patterns/creational/factory
```

This automatically runs:

```text
cpp/main.cpp
go/main.go
python/main.py
```

### Stack

* C++17
* Go
* Python + uv
* Make
* diagrams.net

**Built by Techentia · Karanjot Singh**