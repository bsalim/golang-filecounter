# Directory File Counter

This Go program counts the number of files in each folder within the user's home directory, processing directories concurrently using goroutines. 
It demonstrates the use of goroutines for parallel processing and the `time` package for measuring execution time.

## Features

- Concurrently counts files in each folder using goroutines.
- Processes directories only one level deep.
- Measures and prints the total execution time.


## Requirements

- Go 1.18 or later

## Usage

1. **Clone the repository:**

   ```sh
   git clone <repository-url>
   cd <repository-directory>

2. **Build the program:**
go build -o directory-file-counter

3. **Run it**
./directory-file-counter