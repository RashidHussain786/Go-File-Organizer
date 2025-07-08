# Concurrent URL Checker

A simple and efficient command-line tool written in Go to check the status of multiple URLs concurrently.

## Getting Started

### Prerequisites

- Go (version 1.18 or later)

### Installation

1. Navigate to the project directory:
   ```sh
   cd concurrentUrlChecker
   ```
2. Build the executable:
   ```sh
   go build
   ```

## Usage

The tool reads a list of URLs from a file named `urls.txt` and checks the status of each.

Create a file named `urls.txt` in the same directory and add one URL per line. For example:

```
https://www.google.com
https://www.github.com
https://www.a-non-existent-site.com
```

Then, run the tool:

```sh
./concurrentUrlChecker
```

### Example Output

```
https://www.google.com is UP
https://www.github.com is UP
https://www.a-non-existent-site.com is DOWN
```
