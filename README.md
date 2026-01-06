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
