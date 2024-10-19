package logger

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger_SetPrefix(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(DEBUG, &buf, &buf)

	log.SetPrefix("[TEST] ")
	log.Debugln("Debug message")

	expected := "[TEST] DEBUG: Debug message\n"
	assert.Equal(t, expected, buf.String(), "The output should match the expected debug message with prefix.")
}

func TestLogger_SetErrorPrefix(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(DEBUG, &buf, &buf)

	log.SetErrorPrefix("[ERROR TEST] ")
	log.Errorln("Error message")

	expected := "[ERROR TEST] ERROR: Error message\n"
	assert.Equal(t, expected, buf.String(), "The output should match the expected error message with error prefix.")
}

func TestLogger_SetPrefixes(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(DEBUG, &buf, &buf)

	log.SetPrefixes("[LOG] ", "[ERR_LOG] ")
	log.Warningln("Warning message")
	log.Errorln("Error message")

	expected := "[LOG] WARNING: Warning message\n[ERR_LOG] ERROR: Error message\n"
	println(buf.String())
	assert.Equal(t, expected, buf.String(), "The output should match the expected log and error prefixes.")
}

func TestLogger_SetLogLevel(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(WARNING, &buf, &buf)

	log.Infoln("Info message")       // Should not log since level is set to WARNING
	log.Warningln("Warning message") // Should log
	log.Errorln("Error message")     // Should log

	expected := "WARNING: Warning message\nERROR: Error message\n"
	assert.Equal(t, expected, buf.String(), "The output should only include warning and error messages.")
}

func TestLogger_DebugLevelLogging(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(DEBUG, &buf, &buf)

	log.Debugln("Debug message")
	expected := "DEBUG: Debug message\n"

	assert.Equal(t, expected, buf.String(), "The output should match the expected debug message.")
}

func TestLogger_InfoLevelLogging(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(INFO, &buf, &buf)

	log.Infoln("Info message")
	expected := "Info message\n"

	assert.Equal(t, expected, buf.String(), "The output should match the expected info message.")
}

func TestLogger_WarningLevelLogging(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(WARNING, &buf, &buf)

	log.Warningln("Warning message")
	expected := "WARNING: Warning message\n"

	assert.Equal(t, expected, buf.String(), "The output should match the expected warning message.")
}

func TestLogger_ErrorLevelLogging(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(ERROR, &buf, &buf)

	log.Errorln("Error message")
	expected := "ERROR: Error message\n"

	assert.Equal(t, expected, buf.String(), "The output should match the expected error message.")
}

func TestLogger_SetOutputWriters(t *testing.T) {
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	log := NewLogger(DEBUG, &outBuf, &errBuf)

	log.SetOutputWriters(&outBuf, &errBuf)
	log.Infoln("Info message")
	log.Errorln("Error message")

	assert.Equal(t, "Info message\n", outBuf.String(), "Info message should be logged to the standard output.")
	assert.Equal(t, "ERROR: Error message\n", errBuf.String(), "Error message should be logged to the error output.")
}

func TestLogger_SetOutputWriter(t *testing.T) {
	var outBuf bytes.Buffer
	log := NewLogger(DEBUG, &outBuf, &outBuf)

	log.SetOutputWriter(&outBuf)
	log.Infoln("Info message")

	assert.Equal(t, "Info message\n", outBuf.String(), "Info message should be logged to the standard output.")
}

func TestLogger_SetErrorOutputWriter(t *testing.T) {
	var errBuf bytes.Buffer
	log := NewLogger(DEBUG, &errBuf, &errBuf)

	log.SetErrorOutputWriter(&errBuf)
	log.Errorln("Error message")

	assert.Equal(t, "ERROR: Error message\n", errBuf.String(), "Error message should be logged to the error output.")
}

func TestLogger_DebugF(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(DEBUG, &buf, &buf)

	log.DebugF("Debug formatted message: %d", 42)
	expected := "DEBUG: Debug formatted message: 42"

	assert.Equal(t, expected, buf.String(), "The output should match the expected formatted debug message.")
}

func TestLogger_InfoF(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(INFO, &buf, &buf)

	log.InfoF("Info formatted message: %d", 42)
	expected := "Info formatted message: 42"

	assert.Equal(t, expected, buf.String(), "The output should match the expected formatted info message.")
}

func TestLogger_WarningF(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(WARNING, &buf, &buf)

	log.WarningF("Warning formatted message: %d", 42)
	expected := "WARNING: Warning formatted message: 42"

	assert.Equal(t, expected, buf.String(), "The output should match the expected formatted warning message.")
}

func TestLogger_ErrorF(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(ERROR, &buf, &buf)

	log.ErrorF("Error formatted message: %d", 42)
	expected := "ERROR: Error formatted message: 42"

	assert.Equal(t, expected, buf.String(), "The output should match the expected formatted error message.")
}

func TestLogger_CheckError(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(ERROR, &buf, &buf)

	err := errors.New("some error")
	log.CheckError(err)

	expected := "ERROR: some error"
	assert.Equal(t, expected, buf.String(), "The output should match the expected error message.")
}

func TestLogger_CheckFatal(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(FATAL, &buf, &buf)

	defer func() {
		if r := recover(); r != nil {
			expected := ""
			assert.Equal(t, expected, buf.String(), "The output should be empty.")
		}
	}()

	log.CheckFatal(nil)
}

func TestLogger_CheckErrorln(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(ERROR, &buf, &buf)

	err := errors.New("some error")
	log.CheckErrorln(err)

	expected := "ERROR: some error\n"
	assert.Equal(t, expected, buf.String(), "The output should match the expected error message with newline.")
}

func TestLogger_CheckFatalln(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(FATAL, &buf, &buf)
	defer func() {
		if r := recover(); r != nil {
			expected := ""
			assert.Equal(t, expected, buf.String(), "The output should be empty.")
		}
	}()

	log.CheckFatalln(nil)
}

func TestLogger_CheckErrorF(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(ERROR, &buf, &buf)

	err := errors.New("some error")
	log.CheckErrorF(err, "Error occurred: %s", "file not found")

	expected := "ERROR: Error occurred: file not foundsome error"
	assert.Equal(t, expected, buf.String(), "The output should match the expected formatted error message.")
}
