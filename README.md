# Talking clock

This project provides two ways of returning time in a human-friendly way.

## Development
To test the project, call `make test`

## Build
Build by calling `make`

## Usage
The time is expected to be in twenty-four hours format.

### CLI
Call `./talking-clock-cli` or `./talking-clock-cli 17:30`.

### Server
Make a get request to `http://localhost:8080/` or `http://localhost:8080/?time=15:45`.

## Cleanup
To remove the created executables, call `make clean`.