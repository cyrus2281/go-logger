package logger

import (
	"fmt"
	"io"
	"os"
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
	OFF
)

type logger struct {
	level           int
	outputWriter    io.Writer
	errOutputWriter io.Writer
	prefix          string
	errPrefix       string
}

// getPrefix gets the message prefix based on the log level
func (l *logger) getPrefix(level int) string {
	switch level {
	case DEBUG:
		return l.prefix + "DEBUG: "
	case INFO:
		return l.prefix + ""
	case WARNING:
		return l.prefix + "WARNING: "
	case ERROR:
		return l.errPrefix + "ERROR: "
	case FATAL:
		return l.errPrefix + "ERROR: "
	default:
		return l.prefix + ""
	}
}

// log logs the message with the specified level and line ending option
func (l *logger) log(level int, ln bool, message ...any) {
	prefix := l.getPrefix(level)
	if prefix != "" {
		message = append([]any{prefix}, message...)
	}
	if level >= l.level {
		if level >= ERROR {
			fmt.Fprint(l.errOutputWriter, message...)
		} else {
			fmt.Fprint(l.outputWriter, message...)
		}
		if ln {
			if level >= ERROR {
				fmt.Fprintln(l.errOutputWriter)
			} else {
				fmt.Fprintln(l.outputWriter)
			}
		}
	}
	if level == FATAL {
		os.Exit(1)
	}
}

// SetPrefix sets the prefix for the log message
func (l *logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

// SetErrorPrefix sets the prefix for the error log message
func (l *logger) SetErrorPrefix(prefix string) {
	l.errPrefix = prefix
}

// SetPrefixes sets the prefixes for the log message and the error log message
func (l *logger) SetPrefixes(prefix string, errPrefix string) {
	l.prefix = prefix
	l.errPrefix = errPrefix
}

// SetLogLevel sets the log level
func (l *logger) SetLogLevel(level int) {
	l.level = level
}

// GetLogLevel gets the log level
func (l *logger) GetLogLevel() int {
	return l.level
}

// SetOutputWriters sets the output writers for standard output and error output
func (l *logger) SetOutputWriters(writer io.Writer, errWriter io.Writer) {
	l.outputWriter = writer
	l.errOutputWriter = errWriter
}

// SetOutputWriter sets the output writer for standard output
func (l *logger) SetOutputWriter(writer io.Writer) {
	l.outputWriter = writer
}

// SetErrorOutputWriter sets the output writer for error output
func (l *logger) SetErrorOutputWriter(writer io.Writer) {
	l.errOutputWriter = writer
}

// Debug logs with DEBUG level
func (l *logger) Debug(message ...any) {
	l.log(DEBUG, false, message...)
}

// Info logs with INFO level
func (l *logger) Info(message ...any) {
	l.log(INFO, false, message...)
}

// Warning logs with WARNING level
func (l *logger) Warning(message ...any) {
	l.log(WARNING, false, message...)
}

// Error logs with ERROR level to error output
func (l *logger) Error(message ...any) {
	l.log(ERROR, false, message...)
}

// Fatal logs with FATAL level to error output and exits with 1
func (l *logger) Fatal(message ...any) {
	l.log(FATAL, false, message...)
}

// Debugln logs with DEBUG level and a newline
func (l *logger) Debugln(message ...any) {
	l.log(DEBUG, true, message...)
}

// Infoln logs with INFO level and a newline
func (l *logger) Infoln(message ...any) {
	l.log(INFO, true, message...)
}

// Warningln logs with WARNING level and a newline
func (l *logger) Warningln(message ...any) {
	l.log(WARNING, true, message...)
}

// Errorln logs with ERROR level to error output and a newline
func (l *logger) Errorln(message ...any) {
	l.log(ERROR, true, message...)
}

// Fatalln logs with FATAL level to error output and exits with 1 and a newline
func (l *logger) Fatalln(message ...any) {
	l.log(FATAL, true, message...)
}

// DebugF logs formatted message with DEBUG level
func (l *logger) DebugF(format string, message ...any) {
	l.log(DEBUG, false, fmt.Sprintf(format, message...))
}

// InfoF logs formatted message with INFO level
func (l *logger) InfoF(format string, message ...any) {
	l.log(INFO, false, fmt.Sprintf(format, message...))
}

// WarningF logs formatted message with WARNING level
func (l *logger) WarningF(format string, message ...any) {
	l.log(WARNING, false, fmt.Sprintf(format, message...))
}

// ErrorF logs formatted message with ERROR level to error output
func (l *logger) ErrorF(format string, message ...any) {
	l.log(ERROR, false, fmt.Sprintf(format, message...))
}

// FatalF logs formatted message with FATAL level to error output and exits with 1
func (l *logger) FatalF(format string, message ...any) {
	l.log(FATAL, false, fmt.Sprintf(format, message...))
}

// CheckError checks error and logs with ERROR level to error output if not nil
func (l *logger) CheckError(err error) {
	if err != nil {
		l.log(ERROR, false, err)
	}
}

// CheckErrorln checks error and logs with ERROR level to error output if not nil and a newline
func (l *logger) CheckErrorln(err error) {
	if err != nil {
		l.log(ERROR, true, err)
	}
}

// CheckFatal checks error and logs with FATAL level to error output and exits with 1 if not nil
func (l *logger) CheckFatal(err error) {
	if err != nil {
		l.log(FATAL, false, err)
	}
}

// CheckFatalln checks error and logs with FATAL level to error output and exits with 1 if not nil and a newline
func (l *logger) CheckFatalln(err error) {
	if err != nil {
		l.log(FATAL, true, err)
	}
}

// CheckErrorF checks error and logs formatted message with ERROR level to error output if not nil
func (l *logger) CheckErrorF(err error, format string, message ...any) {
	if err != nil {
		l.log(ERROR, false, fmt.Sprintf(format, message...), err)
	}
}

// CheckFatalF checks error and logs formatted message with FATAL level to error output and exits with 1 if not nil
func (l *logger) CheckFatalF(err error, format string, message ...any) {
	if err != nil {
		l.log(FATAL, false, fmt.Sprintf(format, message...), err)
	}
}

// NewLogger creates a new logger instance with the specified log level, output writer and error output writer
func NewLogger(level int, outputWriter io.Writer, errOutputWriter io.Writer) logger {
	return logger{
		level:           level,
		outputWriter:    outputWriter,
		errOutputWriter: errOutputWriter,
	}
}

// NewDefaultLogger creates a new logger instance with default log level, standard output and error output
func NewDefaultLogger() logger {
	return NewLogger(INFO, os.Stdout, os.Stderr)
}
