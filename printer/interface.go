package printer

import (
	"errors"
	"strings"
)

type (
	// Printer represents a logging interface.
	Printer interface {
		// DebugF sets the color to DebugColor, prints the string as given and resets the color, provided the LogLevel is LevelDebug or higher.
		DebugF(format string, args ...interface{})
		// InfoF sets the color to InfoColor, prints the string as given and resets the color, provided the LogLevel is LevelInfo or higher.
		InfoF(format string, args ...interface{})
		// WarnF sets the color to WarnColor, prints the string as given and resets the color, provided the LogLevel is LevelWarn or higher.
		WarnF(format string, args ...interface{})
		// ErrorF sets the color to ErrorColor, prints the string as given and resets the color.
		ErrorF(format string, args ...interface{})
		// PrintF formats the given format string and prints it to stdout without alterations.
		PrintF(format string, args ...interface{})
		// LogF formats the given format string and prints it to stdout using the color set by UseColor, defaulting to InfoColor if not set.
		LogF(format string, args ...interface{})

		// UseColor will internally save the given color and use it when LogF is invoked.
		UseColor(color Color) Printer
		// MapColorToLevel sets the color for the given log level.
		MapColorToLevel(color Color, level LogLevel) Printer

		// SetColor sets the given color without printing newline char and returns itself.
		// Best used with PrintF as other methods may reset the color before returning.
		SetColor(color Color)
		// ResetColor will reset the color to default.
		// Best used in a defer statement right after setting a color to not forget about resetting.
		ResetColor()
		// SetLevel sets the logging level.
		SetLevel(level LogLevel) Printer
		// SetName sets a prefix for DebugF, InfoF, WarnF
		SetName(name string) Printer
		// CheckIfError will print the error in ErrorColor to stderr if it is non-nil and exit with exit code 1.
		CheckIfError(err error)
	}
	LogLevel int
	Color    string
)

const (
	// LevelError is the error logging level.
	LevelError = 0
	// LevelWarn is the warning logging level.
	LevelWarn = 1
	// LevelInfo is the info logging level.
	LevelInfo = 2
	// LevelDebug is the debug logging level.
	LevelDebug = 3
)

var (
	// DefaultPrinter is the default logger instance.
	DefaultPrinter = New()
	// DefaultLevel is the default logging level for new logger instances.
	DefaultLevel LogLevel = LevelInfo
)

// New returns a new colored logger.
func New() Printer {
	return &colorPrinter{
		level: DefaultLevel,
		color: InfoColor,
		name:  "",
		colorMap: map[LogLevel]Color{
			LevelError: ErrorColor,
			LevelWarn:  WarnColor,
			LevelInfo:  InfoColor,
			LevelDebug: DebugColor,
		},
	}
}

// DebugF uses DefaultPrinter to log the given format string and its arguments.
func DebugF(format string, args ...interface{}) {
	DefaultPrinter.DebugF(format, args...)
}

// InfoF uses DefaultPrinter to log the given format string and its arguments.
func InfoF(format string, args ...interface{}) {
	DefaultPrinter.InfoF(format, args...)
}

// WarnF uses DefaultPrinter to log the given format string and its arguments.
func WarnF(format string, args ...interface{}) {
	DefaultPrinter.WarnF(format, args...)
}

// CheckIfError uses DefaultPrinter to print the given error and exit the program.
func CheckIfError(err error) {
	DefaultPrinter.CheckIfError(err)
}

// ParseLogLevel returns the parsed log level from a given lower-cased string.
// If it cannot be parsed, the LevelInfo is returned along with an error.
func ParseLogLevel(level string) (LogLevel, error) {
	levelMap := map[string]LogLevel{
		"error": LevelError,
		"warn":  LevelWarn,
		"info":  LevelInfo,
		"debug": LevelDebug,
	}
	lvl, found := levelMap[strings.ToLower(level)]
	if found {
		return lvl, nil
	}
	return LevelInfo, errors.New("log level invalid: " + level)
}
