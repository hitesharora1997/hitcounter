# Request Counter Server

Request Counter Server is a simple HTTP server written in Go. It responds to each request with a counter of the total number of requests received during the previous 60 seconds (moving window). The server persists data to a file, allowing it to continue returning the correct numbers after a restart, using only the standard library.

## Usage

### Build and Run

Follow these steps to build and run the server:

1. **Install Go**: Ensure Go is installed on your system. Download it from [https://golang.org/dl/](https://golang.org/dl/).

2. **Clone Repository**:
   ```bash
   git clone https://github.com/hitesharora1997/hitcounter.git
   ```

3. **Navigate to Directory**:
   ```bash
   cd hitcounter
   ```

4. **Build the Server**:
   ```bash
   make build
   ```

5. **Run the Server**:
   ```bash
   make run
   ```

### API Endpoint

To get the total number of requests made in the last 60 seconds, make a GET request to:
```
http://localhost:8090/
```

## Development Commands

- **Build**: Compile the server with `make build`.
- **Run**: Start the server with `make run`.
- **Test**: Run all tests with `make test`.
- **Race Condition Test**: Test for race conditions with `make test-race`.
- **Clean Test Cache**: Clean the test cache and run tests with `make test-clean-testcache`.
- **Benchmarks**: Run benchmarks with `make benchmarks`.
- **Coverage**: Generate coverage files with `make coverage`.
- **Clean**: Clean up the project directory with `make clean`.

