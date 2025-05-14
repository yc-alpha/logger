package benchmark

import (
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/yc-alpha/logger"
	"github.com/yc-alpha/logger/backend"
)

// BenchmarkInfo-12		24850		58835 ns/op		0 B/op		0 allocs/op
func BenchmarkInfo(b *testing.B) {
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
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		log.Info("This is an info message, we use it to test the performance of the logger")
	}
	b.ReportAllocs()

}

// BenchmarkInfos-12
// 27985             51170 ns/op               0 B/op          0 allocs/op
func BenchmarkInfos(b *testing.B) {
	log := logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithBackends(logger.AnyLevel, backend.OSBackend().Build()),
		logger.WithSeparator(logger.AnyLevel, " "),
		logger.WithFields(logger.AnyLevel,
			logger.DatetimeField("2006/01/02 15:04:05"),
			logger.LevelField().Upper(),
			logger.MessageField(),
		),
		logger.WithEncoders(logger.AnyLevel, logger.LogFmtEncoder),
	)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		log.Infos("hello world",
			logger.Group("TEST",
				logger.Float64("float64", 6.1),
				logger.String("string", "test"),
				logger.Int64("int64", 7),
			),
		)
	}
	b.ReportAllocs()
}

// BenchmarkDefaultLogger-12
// 26377             54866 ns/op             232 B/op          2 allocs/op
func BenchmarkDefaultLogger(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Info("hello world!")
	}
	b.ReportAllocs()
}

// BenchmarkSlog-12
// 25784             49782 ns/op             533 B/op          8 allocs/op
func BenchmarkSlog(b *testing.B) {
	s := slog.New(slog.NewTextHandler(os.Stderr, nil))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Info("hello world",
			slog.Group("TEST",
				slog.String("test", "6"),
				slog.String("test", "7"),
				slog.String("test", "7"),
			),
			slog.String("test", strconv.Itoa(n)),
		)
	}
}
