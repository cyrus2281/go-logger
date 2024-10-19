package logger

import (
	"io"
	"os"
)

// Default Logger instance
var loggerInstance logger

func init() {
	loggerInstance = NewLogger(INFO, os.Stdout, os.Stderr)
}

// SetPrefixFormatter sets the prefix formatter for the log message
func SetPrefixFormatter(prefixFormatter func(level int) string) {
	loggerInstance.SetPrefixFormatter(prefixFormatter)
}

// SetErrorPrefixFormatter sets the prefix for the error log message
func SetErrorPrefixFormatter(errorPrefixFormatter func(level int) string) {
	loggerInstance.SetErrorPrefixFormatter(errorPrefixFormatter)
}

// SetPrefixFormatters sets the prefix formatters for the log message and the error log message
func SetPrefixFormatters(prefixFormatter func(level int) string, errorPrefixFormatter func(level int) string) {
	loggerInstance.SetPrefixFormatters(prefixFormatter, errorPrefixFormatter)
}

// SetLogLevel sets the log level
func SetLogLevel(level int) {
	loggerInstance.SetLogLevel(level)
}

// GetLogLevel gets the log level
func GetLogLevel() int {
	return loggerInstance.level
}

// SetOutputWriters sets the output writers for standard output and error output
func SetOutputWriters(writer io.Writer, errWriter io.Writer) {
	loggerInstance.SetOutputWriters(writer, errWriter)
}

// SetOutputWriter sets the output writer for standard output
func SetOutputWriter(writer io.Writer) {
	loggerInstance.SetOutputWriter(writer)
}

// SetErrorOutputWriter sets the output writer for error output
func SetErrorOutputWriter(writer io.Writer) {
	loggerInstance.SetErrorOutputWriter(writer)
}

// Debug logs with DEBUG level
func Debug(message ...any) {
	loggerInstance.Debug(message...)
}

// Info logs with INFO level
func Info(message ...any) {
	loggerInstance.Info(message...)
}

// Warning logs with WARNING level
func Warning(message ...any) {
	loggerInstance.Warning(message...)
}

// Error logs with ERROR level to error output
func Error(message ...any) {
	loggerInstance.Error(message...)
}

// Fatal logs with FATAL level to error output and exits with 1
func Fatal(message ...any) {
	loggerInstance.Fatal(message...)
}

// Debugln logs with DEBUG level and a newline
func Debugln(message ...any) {
	loggerInstance.Debugln(message...)
}

// Infoln logs with INFO level and a newline
func Infoln(message ...any) {
	loggerInstance.Infoln(message...)
}

// Warningln logs with WARNING level and a newline
func Warningln(message ...any) {
	loggerInstance.Warningln(message...)
}

// Errorln logs with ERROR level to error output and a newline
func Errorln(message ...any) {
	loggerInstance.Errorln(message...)
}

// Fatalln logs with FATAL level to error output and exits with 1 and a newline
func Fatalln(message ...any) {
	loggerInstance.Fatalln(message...)
}

// DebugF logs formatted message with DEBUG level
func DebugF(format string, message ...any) {
	loggerInstance.DebugF(format, message...)
}

// InfoF logs formatted message with INFO level
func InfoF(format string, message ...any) {
	loggerInstance.InfoF(format, message...)
}

// WarningF logs formatted message with WARNING level
func WarningF(format string, message ...any) {
	loggerInstance.WarningF(format, message...)
}

// ErrorF logs formatted message with ERROR level to error output
func ErrorF(format string, message ...any) {
	loggerInstance.ErrorF(format, message...)
}

// FatalF logs formatted message with FATAL level to error output and exits with 1
func FatalF(format string, message ...any) {
	loggerInstance.FatalF(format, message...)
}

// CheckError checks error and logs with ERROR level to error output if not nil
func CheckError(err error) {
	loggerInstance.CheckError(err)
}

// CheckErrorln checks error and logs with ERROR level to error output if not nil and a newline
func CheckErrorln(err error) {
	loggerInstance.CheckErrorln(err)
}

// CheckFatal checks error and logs with FATAL level to error output and exits with 1 if not nil
func CheckFatal(err error) {
	loggerInstance.CheckFatal(err)
}

// CheckFatalln checks error and logs with FATAL level to error output and exits with 1 if not nil and a newline
func CheckFatalln(err error) {
	loggerInstance.CheckFatalln(err)
}

// CheckErrorF checks error and logs formatted message with ERROR level to error output if not nil
func CheckErrorF(err error, format string, message ...any) {
	loggerInstance.CheckErrorF(err, format, message...)
}

// CheckFatalF checks error and logs formatted message with FATAL level to error output and exits with 1 if not nil
func CheckFatalF(err error, format string, message ...any) {
	loggerInstance.CheckFatalF(err, format, message...)
}
