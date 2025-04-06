package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	colorRed     color = "\033[31m"
	colorYellow  color = "\033[33m"
	colorBlue    color = "\033[34m"
	colorDefault color = "\033[0m"
)

func paint(color color, message string) string {
	return color + message + colorDefault
}

type config struct {
	out           *os.File
	err           *os.File
	in            *os.File
	lineSeparator byte
	timestamp     bool
	timeFormat    string
	color         bool
}

type Console struct {
	conf *config
}

type color = string

func defaultConfig() *config {
	return &config{
		in:            os.Stdin,
		out:           os.Stdout,
		err:           os.Stderr,
		lineSeparator: '\n',
	}
}

var DEFAULT = NewConsole()

func NewConsole() *Console {
	return &Console{
		conf: defaultConfig(),
	}
}

func (console *Console) SetTimestampFormat(format string) {
	console.conf.timestamp = true
	console.conf.timeFormat = format
}

func (console *Console) SetLogColor(yes bool) {
	console.conf.color = yes
}

func (console *Console) SetLineSeparator(separator byte) {
	console.conf.lineSeparator = separator
}

func (console *Console) SetInChannel(channel *os.File) {
	console.conf.in = channel
}

func (console *Console) SetOutChannel(channel *os.File) {
	console.conf.out = channel
}

func (console *Console) SetErrChannel(channel *os.File) {
	console.conf.err = channel
}

func (console *Console) GetInChannel() *os.File {
	return console.conf.in
}

func (console *Console) GetOutChannel() *os.File {
	return console.conf.out
}

func (console *Console) GetErrChannel() *os.File {
	return console.conf.err
}

func (console *Console) buildLogMessage(color color, prefix, source, message string) string {

	if strings.TrimSpace(source) == "" {
		source = "Application"
	}

	payload := "(" + source + ") "

	if console.conf.timestamp {
		payload += time.Now().Format(console.conf.timeFormat) + " "
	}

	payload += "-> " + prefix + ": "
	if console.conf.color {
		payload = paint(color, payload)
	}

	payload += message

	return payload
}

func (console *Console) Log(message string) (n int, err error) {
	return fmt.Fprintln(console.conf.out, message)
}

func (console *Console) Error(message, source string) (n int, err error) {
	return fmt.Fprintln(console.conf.out, console.buildLogMessage(colorRed, "ERROR", source, message))
}

func (console *Console) Debug(message, source string) (n int, err error) {
	return fmt.Fprintln(console.conf.out, console.buildLogMessage(colorYellow, "DEBUG", source, message))

}

func (console *Console) Trace(message, source string) (n int, err error) {
	return fmt.Fprintln(console.conf.out, console.buildLogMessage(colorBlue, "TRACE", source, message))

}

func (console *Console) ReadString() (string, error) {
	return bufio.NewReader(console.conf.in).ReadString(console.conf.lineSeparator)
}

func (console *Console) Close() (error, error, error) {
	var inErr, outErr, errErr error
	if console.conf.in != nil {
		inErr = console.conf.in.Close()
	}

	if console.conf.out != nil {
		outErr = console.conf.out.Close()
	}

	if console.conf.err != nil {
		errErr = console.conf.err.Close()
	}

	return inErr, outErr, errErr
}
