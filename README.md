# Logger

A high-performance, configurable logging library for Go applications.

## Features

- Multiple log levels (Debug, Info, Warn, Error, Fatal, Panic)
- Flexible output formatting (Plain text, JSON, LogFmt)
- Customizable log fields
- Color support for console output
- Log file rotation
- Zero allocation logging
- Context support
- Thread-safe

## Installation

```bash
go get github.com/yc-alpha/logger
```

## Quick Start

```go
package main

import "github.com/yc-alpha/logger"

func main() {
    // Use default logger
    logger.Info("hello world")
    
    // Create custom logger
	log := logger.NewLogger(
		// Won't output logs below Info level
		logger.WithLevel(logger.InfoLevel),
		// Set the output to stdout for all log levels
		logger.WithBackends(logger.AnyLevel, backend.OSBackend().Build()),
		logger.WithFields(logger.AnyLevel,
			logger.DatetimeField(time.DateTime),
			logger.LevelField().Key("level").Upper(),
			logger.MessageField(),
		),
		// Set the encoder to LogFmtEncoder for all log levels
		logger.WithEncoders(logger.AnyLevel, logger.LogFmtEncoder),
	)
    
    log.Info("hello logger!")
}
```

## Configuration

### Log Levels

- `DebugLevel`
- `InfoLevel` (default)
- `WarnLevel`
- `ErrorLevel` 
- `FatalLevel`
- `PanicLevel`

### Backends

- Console output (`OSBackend`)
- File output with rotation (`DefaultFileBackend`)

### Fields

Built-in fields:
- `LevelField` - Log level
- `MessageField` - Log message
- `DatetimeField` - Formated time
- `TimeField` - Timestamp
- `CallerField` - File and line number
- `FuncNameField` - Function name
- `CustomField` - User-defined field

### Formatting

Supported encoders:
- `PlainEncoder` - Simple text format
- `JSONEncoder` - JSON format  
- `LogFmtEncoder` - LogFmt format

## Examples

### Adding Custom Fields

```go
log.Infos("hello",
    logger.Group("metadata",
        logger.String("user", "admin"),
        logger.Int("id", 123),
    ),
)
```

### Using Colors

```go
log := logger.NewLogger(
    logger.WithFields(logger.ErrorLevel,
        logger.LevelField().Color(logger.Red),
        logger.MessageField().Color(logger.Yellow),
    ),
)
```

### Customize different fields for different log levels
- Add a group field "sys_info" to info and warn level logs
- Add `file` and `func` fields to warn level logs and above
```go
log := NewLogger(
	WithLevel(InfoLevel),
	WithBackends(AnyLevel, backend.OSBackend().Build()),
	WithFields(AnyLevel,
		DatetimeField(time.DateTime).Key("datetime"),
		LevelField().Key("level").Upper().Prefix("[").Suffix("]"),
		MessageField().Key("msg"),
	),
	WithFields(InfoLevel|WarnLevel,
		Group("sys_info",
			CustomField(func(buf *buffer.Buffer) {
				buildInfo, _ := debug.ReadBuildInfo()
				buf.WriteString(buildInfo.GoVersion)
			}).Key("go_version"),
			Group("sys", CustomField(func(buf *buffer.Buffer) {
				buf.WriteInt(int64(os.Getpid()))
			}).Key("pid")),
		)),
	WithFields(ErrorLevel|FatalLevel|PanicLevel,
		CallerField(true, false).Key("file"),
		FuncNameField(true).Key("func"),
	),
	WithEncoders(AnyLevel, LogFmtEncoder),
)
log.Debug("hello debug")
log.Info("hello info")
log.Warn("hello warn")
log.Error("hello error")
```

### File Rotation

```go
log := logger.NewLogger(
    logger.WithBackends(logger.InfoLevel,
        backend.DefaultFileBackend().
            Filename("app.log").
            MaxSize(100).  // MB
            MaxBackups(5).
            Compress(true).
            Build(),
    ),
)
```

## Performance
View test case [Here](./benchmark/benchmark_test.go)
```text
BenchmarkInfo-12			24850		58835 ns/op		0 B/op		0 allocs/op
BenchmarkInfos-12       	27985       51170 ns/op     0 B/op      0 allocs/op
BenchmarkDefaultLogger-12	26377		54866 ns/op		232 B/op	2 allocs/op
```


## Reference Projects
- [zap](https://github.com/uber-go/zap):Refer to the implementation of buffer
- [lumberjack](https://github.com/natefinch/lumberjack):Refer to lumberjack for file rotation implementation

## License

MIT License - see [LICENSE](LICENSE) for details