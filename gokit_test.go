package bench

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
)

// Go kit's logger has no concept of dynamically mutable levels. The idiom is to
// predeclare your desired level during construction. This is an example helper
// constructor that performs that work. If positive is true, both info and error
// are logged. Otherwise, only error is logged.
func newLeveledLogger(logger log.Logger, positive bool) *leveledLogger {
	infoLogger := log.NewNopLogger()
	if positive {
		infoLogger = log.With(logger, "level", "info")
	}
	return &leveledLogger{
		Info:  infoLogger,
		Error: log.With(logger, "level", "error"),
	}
}

type leveledLogger struct {
	Info  log.Logger
	Error log.Logger
}

func BenchmarkGokitTextFile(b *testing.B) {

	tmpfile, err := ioutil.TempFile("", "benchmark-gokit")
	if err != nil {
		b.Fatal(err)
	}

	logger := log.With(log.NewLogfmtLogger(tmpfile))

	//logger := log.With(log.NewLogfmtLogger(stream), "ts", log.DefaultTimestampUTC)
	lvllog := newLeveledLogger(logger, true)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkGokitTextStd(b *testing.B) {
	logger := log.With(log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr)))

	lvllog := newLeveledLogger(logger, true)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkGokitJSONFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-gokit")
	if err != nil {
		b.Fatal(err)
	}

	logger := log.With(log.NewLogfmtLogger(tmpfile))
	lvllog := newLeveledLogger(logger, true)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})
}

func BenchmarkGokitJSONStd(b *testing.B) {
	logger := log.With(log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr)))
	lvllog := newLeveledLogger(logger, true)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})
}
