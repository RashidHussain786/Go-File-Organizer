# Go File Organizer

A simple and efficient command-line tool written in Go to organize files in a directory based on their extensions.

## Getting Started

### Prerequisites

- Go (version 1.18 or later)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/RashidHussain786/Go-File-Organizer.git
   ```
2. Navigate to the project directory:
   ```sh
   cd Go-File_Organizer
   ```
3. Build the executable:
   ```sh
   go build
   ```

## Usage

The tool provides several flags to customize its behavior:

```
./go-file-organizer [flags]
```

### Flags

| Flag        | Description                                       | Default       |
|-------------|---------------------------------------------------|---------------|
| `-src`      | The source directory to organize                  | `.`           |
| `-dest`     | The destination directory for organized files     | `./organized` |
| `-dry-run`  | Perform a dry run without moving files            | `false`       |
| `-verbose`  | Enable verbose output                             | `false`       |
| `-recursive`| Recursively organize files in subdirectories      | `false`       |

### Examples

#### Basic Organization

Organize files in the current directory and move them to the default `./organized` directory.

```sh
./go-file-organizer
```

#### Specify Source and Destination

Organize files from `~/Downloads` and move them to `~/Documents/Organized`.

```sh
./go-file-organizer -src ~/Downloads -dest ~/Documents/Organized
```

#### Recursive Organization

Recursively organize all files in the `test-data` directory, including those in subfolders.

```sh
./go-file-organizer -src test-data -recursive -verbose
```

#### Dry Run

See what changes would be made without actually moving any files.

```sh
./go-file-organizer -src test-data -dry-run -verbose
```

