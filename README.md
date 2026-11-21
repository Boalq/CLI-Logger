# CLI-Logger for Obsidian Daily Notes

A minimal Go CLI tool that appends a timestamped log entry to today's Obsidian Daily Note.
If the note doesn't exist yet, it is created.

```
$ log "fixed the build bug"
'fixed the build bug' successfully written to 21-11-2026.md
```

## Setup

### 1. Adjust two lines in `cmd/main.go`

- **Line 12:** path to the folder of your daily notes:
  ```go
  const note_dir = "/home/boalq/Documents/BoranPersonal/09 Calender/"   // <- your path
  ```
- **Line 25:** file name pattern — must match your daily note names:
  ```go
  date := fmt.Sprintf("%d-%d-%d.md", day, month, year)   // 21-11-2026.md
  ```
  e.g. for Obsidian's default `2026-11-21.md` use `year, month, day` (and `%02d` for zero-padded names like `2026-11-05.md`).

### 2. Build

```sh
go build -o build/cli-logger ./cmd
```

### 3. Bind it to your terminal

Use the absolute path to the binary.

**Fish** — create `~/.config/fish/functions/log.fish` (autoloaded on new shells):

```fish
function log
    /home/YOURUSERNAME/CLI-Logger/build/cli-logger $argv
end
```

**Bash** — add to `~/.bashrc`:

```bash
alias log='/home/YOURUSERNAME/CLI-Logger/build/cli-logger'
```

## Usage

```sh
log "my note text"
```

Takes exactly one argument — quote texts with spaces.
