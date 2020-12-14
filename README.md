# Advent of Code Solutions
This are _my_ solutions to the [AdventOfCode](https://adventofcode.com) 2020 programming puzzles written in the Go programming language.

I would strongly suggest that you try to solve the puzzles before checking the solutions. 

# Instructions

Each user gets a different set if input data for each puzzle. To run the solutions yourself you will need first to get your session token exported.

You can get yours by logging into [AdventOfCode](https://adventofcode.com) and inspecting your cookies contents. You should have a cookie for the `.adventofcode.com` domain. Export the cookie value as:

```bash
export SESSION=536.........
```

Once you have your shell populated with your SESSION env you can run

```
$ make dayN
```

to get the solutions for Puzzle N ( where N is an integer )

For example, to get the solutions for Day 1 Puzzle run:

```
$ make day1
```

# Structure

I tried to keep the structure for each puzzle as similar as possible, that way checking the code for different puzzles should be easier for the reader.

In addition, I structured the tests to be as close as possible to the Puzzle instructions.

