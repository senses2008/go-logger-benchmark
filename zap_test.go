package bench

import (
	"io/ioutil"
	"testing"

	"go.uber.org/zap"
)

func BenchmarkZapTextFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-zap")
	if err != nil {
		b.Fatal(err)
	}

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()

	defer logger.Sync()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkZapTextStd(b *testing.B) {

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stderr"}
	logger, err := cfg.Build()

	defer logger.Sync()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkZapJSONFile(b *testing.B) {

	tmpfile, err := ioutil.TempFile("", "benchmark-zap")
	if err != nil {
		b.Fatal(err)
	}

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog",
				zap.String("rate", "15"),
				zap.Int("low", 16),
				zap.Float32("high", 123.2),
			)
		}
	})
}

func BenchmarkZapJSONStd(b *testing.B) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stderr"}
	logger, err := cfg.Build()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog",
				zap.String("rate", "15"),
				zap.Int("low", 16),
				zap.Float32("high", 123.2),
			)
		}
	})
}
