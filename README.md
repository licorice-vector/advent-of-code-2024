# Advent of Code 2024

Welcome to my solutions for [Advent of Code 2024](https://adventofcode.com/2024)! This repository contains solutions written in **Go**.

## 🎄 About Advent of Code
Advent of Code is an annual coding event where participants solve daily programming puzzles during the month of December. Each day provides a new challenge, ranging in difficulty, often requiring creative problem-solving and algorithmic thinking.

## 🚀 Getting Started

### Prerequisites
- Go: Ensure you have Go installed. You can download it from golang.org.
- Git: Clone the repository using Git.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/licorice-vector/advent-of-code-2024.git
   cd advent-of-code-2024
   ```

2. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

### Running a Solution
Each day's solution is in its own directory under ./days/dayXX/. To run a solution:
   ```bash
   go run days/day01/solution.go days/day01/main.go
   ```

### Testing
To run the tests for a specific day:
   ```bash
   go test ./days/day01/
   ```

## 📂 Project Structure
```
advent-of-code-2024/
├── days/
│   ├── day01/
│   │   ├── main.go           # Main code
│   │   ├── example_test.go   # Test of the example
│   │   ├── solution.go       # Solution for Day 1
│   │   ├── input.txt         # Input data for Day 1
│   │   └── example.txt       # Example input for testing
│   └── dayXX/                # Solutions for subsequent days
├── .gitignore                # Ignored files and directories
├── go.mod                    # Go module file
├── go.sum                    # Dependencies
└── README.md                 # Project documentation
```

## 🛠 Tools & Libraries
- Go: The primary programming language.
- VS Code: Recommended editor with Go extensions for debugging and linting.

## 🎯 Goals
- Solve all 25 puzzles (or as many as possible).
- Write clean, maintainable, and efficient Go code.
- Explore new algorithms and data structures.

## 📅 Progress
| Day   | Part 1 | Part 2 |
|-------|--------|--------|
| Day 1 | ✅     | ✅     |
| Day 2 | ✅     | ✅     |
| Day 3 | ✅     | ✅     |
| Day 4 | ✅     | ✅     |
| Day 5 | ✅     | ✅     |
| Day 6 | ✅     | ✅     |
| Day 7 | ✅     | ✅     |
| Day 8 | ✅     | ✅     |
| Day 9 | ✅     | ✅     |
| Day 10| ✅     | ✅     |
| Day 11| ✅     | ✅     |
| Day 12| ✅     | ✅     |
| Day 13| ✅     | ✅     |
| Day 14| ✅     | ✅     |
| Day 15| ✅     | ✅     |
| Day 16| ✅     | ✅     |
| Day 17| ✅     | ✅     |
| Day 18| ✅     | ✅     |
| Day 19| ✅     | ✅     |
| Day 20| ✅     | ✅     |
| Day 21| ✅     | ✅     |
| Day 22| ✅     | ✅     |
| Day 23| ✅     | ✅     |
| Day 24| ✅     | ✅     |
| Day 25| ✅     | ⭐     |

## 🌟 Acknowledgments
Thanks to Eric Wastl for creating Advent of Code and providing an amazing way to challenge ourselves each year.

## 📜 License
This project is licensed under the MIT License. See the LICENSE file for details.
