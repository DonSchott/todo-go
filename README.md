# todo

A minimal command-line todo manager written in Go.

## Installation
```bash
go install
```

Or manually:
```bash
go build
sudo cp todo /usr/local/bin/
```

## Usage

### Add a todo
```bash
todo add "Buy groceries"
```

### List all todos
```bash
todo list
```

### Mark as done
```bash
todo done 1
```

### Remove a todo
```bash
todo remove 1
```

## Storage

Todos are stored in `~/.todo.json` in human-readable JSON format.

## License

MIT

## About

This is a learning project to explore Go fundamentals. Built with assistance from Claude (Anthropic) as a practical introduction to:
- Go syntax and idioms
- Command-line argument parsing
- JSON serialization
- File I/O
- Struct types

The goal was to implement the same simple todo app previously written in Python, to compare languages and understand Go's approach.
