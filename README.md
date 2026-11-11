# Advent of Code

My solutions to [Advent of Code](https://adventofcode.com). Each year uses a different language. The solutions for a year are in `<year>-<language>/`, for example `2015-go/`.

## Inputs

This repo does not include the puzzle inputs or the answers. The Advent of Code author asks people not to share them ([FAQ](https://adventofcode.com/about#faq_copying)). I keep mine in a private repo, linked here as a git submodule at `inputs/`.

To run a solution, you must add your own inputs and expected answers. Get your input for each puzzle on [adventofcode.com](https://adventofcode.com) (you must log in). The expected answers are the ones the site accepted for your input.

Put the files in `inputs/` in this format:

```
inputs/
└── 2015/                 year
    └── 01/               day, two digits (01 to 25)
        ├── input.txt     your puzzle input
        ├── part1.txt     expected answer for part 1
        └── part2.txt     expected answer for part 2
```

The answer files contain only the answer.

Git cannot get the `inputs` submodule when you clone this repo, because it is private. Clone the repo without `--recurse-submodules`, then put your files in the empty `inputs/` folder.
