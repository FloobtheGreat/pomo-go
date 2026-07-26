# Pomo-Go

Pomo-Go is a small terminal-based Pomodoro timer written in Go. It guides you through selecting a timer preset, setting the number of cycles, and then runs focused work and break intervals from the command line.

## Features

- Interactive terminal setup
- Preset timer modes
- Repeating work/break cycles
- Minimal dependency footprint

## Requirements

- Go 1.25 or newer

## Run

```bash
go run .
```

## Notes

- Cycle count is limited to 1-4.
- Timer presets currently map to standard, short, and long intervals.
