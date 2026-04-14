# Monstercat Pokemon Sequence

A small Go CLI that builds a Pokemon word chain from a starting letter.

## What It Does

- Loads Pokemon names from `sequence/pokemon.txt`.
- Prompts you for a letter in a loop.
- Builds a sequence where:
  - the first Pokemon starts with your input letter
  - each next Pokemon starts with the last letter of the previous one
  - each Pokemon is used at most once per generated sequence
- Prints the result as a comma-separated list.

## Project Structure

- `main.go`: CLI loop, input handling, and printing output.
- `sequence/sequence.go`: sequence generation logic.
- `sequence/pokemon.txt`: source dataset.

## Requirements

- Go (matching `go.mod`)

## Run

```bash
go run .
```

Then type a starting letter and press Enter.

## Example

Input:

```text
l
```

Output format:

```text
lumineon, nosepass, ...
```

## Notes

- If no Pokemon starts with the entered letter, the program prints `no matching words`.
- The app runs continuously; stop it with `Ctrl+C`.
