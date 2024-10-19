# Go Logger Package

[![GitHub Tag](https://img.shields.io/github/v/tag/cyrus2281/go-logger?label=Version)](https://pkg.go.dev/github.com/cyrus2281/go-logger)
[![License](https://img.shields.io/github/license/cyrus2281/go-logger)](https://github.com/cyrus2281/go-logger/blob/main/LICENSE)
[![buyMeACoffee](https://img.shields.io/badge/BuyMeACoffee-cyrus2281-yellow?logo=buymeacoffee)](https://www.buymeacoffee.com/cyrus2281)
[![GitHub issues](https://img.shields.io/github/issues/cyrus2281/go-logger?color=red)](https://github.com/cyrus2281/go-logger/issues)
[![GitHub stars](https://img.shields.io/github/stars/cyrus2281/go-logger?style=social)](https://github.com/cyrus2281/go-logger/stargazers)

- [Go Logger Package](#go-logger-package)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Using the Default Logger](#using-the-default-logger)
    - [Basic Logger Setup](#basic-logger-setup)
    - [Customizing Log Prefixes](#customizing-log-prefixes)
    - [Log Levels](#log-levels)
    - [Formatted Logs](#formatted-logs)
    - [Error Checking](#error-checking)
  - [Log Level Behavior](#log-level-behavior)
  - [Advanced Usage](#advanced-usage)
    - [Custom Output Writers](#custom-output-writers)
  - [Contributing](#contributing)
    - [Running Tests](#running-tests)
    - [License](#license)

A lightweight and customizable logging package for Go. This package provides support for different log levels, output streams, formatted logs, and error handling, making it easy to integrate logging into your application.

## Features

- Support for log levels: `DEBUG`, `INFO`, `WARNING`, `ERROR`, `FATAL`, `OFF`
- Customizable log message and error prefixes
- Separate output writers for standard logs and errors
- Easy-to-use functions for logging formatted messages
- Error checking with built-in logging for `ERROR` and `FATAL` levels
- Fatal logs that exit the program

## Installation

```bash
go get github.com/cyrus2281/go-logger
```

## Usage

### Using the Default Logger

The default logger writes to `os.Stdout` and `os.Stderr` and has the log level set to `INFO`.
All of which can be customized using the `SetOutputWriters`, `SetLogLevel`, and `SetPrefixFormatters` methods.


```go
package main

import (
	"os"
	"github.com/cyrus2281/go-logger/logger"
)

func main() {
	// Logging examples
	logger.Infoln("This is an info message")
	logger.Warningln("This is a warning message")
	logger.Errorln("This is an error message")
}
```
### Basic Logger Setup

You can create a new logger using the `NewLogger` function by specifying the log level and output writers:

```go
package main

import (
	"os"
	"github.com/cyrus2281/go-logger/logger"
)

func main() {
	// Create a logger with INFO level
	log := logger.NewLogger(logger.INFO, os.Stdout, os.Stderr)

	// Logging examples
	log.Infoln("This is an info message")
	log.Warningln("This is a warning message")
	log.Errorln("This is an error message")
}
```

### Customizing Log Prefixes

You can set prefixes for standard log messages and error messages:

```go
func formatter(level int) string {
    currentDate := time.Now().Format("2006-01-02")
	return fmt.Sprintf("[%s] [%d] ", currentDate, level)
}

log.SetPrefixFormatter(formatter)
log.SetErrorPrefixFormatter(formatter)

log.Infoln("Application started")
log.Errorln("Failed to connect to database")
```

### Log Levels

The package supports the following log levels, allowing you to filter messages based on importance:

- `DEBUG`
- `INFO`
- `WARNING`
- `ERROR`
- `FATAL`
- `OFF`

```go
log.SetLogLevel(logger.WARNING)

log.Debugln("This will not be logged")
log.Infoln("This will not be logged")
log.Warningln("This will be logged")
log.Errorln("This will be logged")
```

### Formatted Logs

You can log formatted messages using the `DebugF`, `InfoF`, `WarningF`, `ErrorF`, and `FatalF` methods.

```go
log.InfoF("User %s logged in at %s", "john_doe", "10:00 AM")
log.ErrorF("Failed to load file: %s", "config.yaml")
```

### Error Checking

The logger provides utility methods for checking errors and logging them automatically:

```go
err := errors.New("connection failed")
log.CheckError(err)  // Logs if `err` is not nil

fatalErr := errors.New("fatal error")
log.CheckFatal(fatalErr)  // Logs and exits the program if `fatalErr` is not nil
```

## Log Level Behavior

The log level controls which messages will be printed based on their importance:

| Log Level | Logs Debug | Logs Info | Logs Warning | Logs Error | Logs Fatal |
| --------- | ---------- | --------- | ------------ | ---------- | ---------- |
| `DEBUG`   | ✔          | ✔         | ✔            | ✔          | ✔          |
| `INFO`    | ✖          | ✔         | ✔            | ✔          | ✔          |
| `WARNING` | ✖          | ✖         | ✔            | ✔          | ✔          |
| `ERROR`   | ✖          | ✖         | ✖            | ✔          | ✔          |
| `FATAL`   | ✖          | ✖         | ✖            | ✖          | ✔          |
| `OFF`     | ✖          | ✖         | ✖            | ✖          | ✖          |

## Advanced Usage

### Custom Output Writers

You can direct standard and error logs to different output writers (e.g., files, network streams, etc.).

```go
var outBuf, errBuf bytes.Buffer

log.SetOutputWriters(&outBuf, &errBuf)

log.Infoln("This will be written to outBuf")
log.Errorln("This will be written to errBuf")
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to help improve this package.

### Running Tests

To run the tests for this package, use:

```bash
go test ./...
```

### License

This project is licensed under the Apache v2.0 License - see the [LICENSE](LICENSE) file for details.
