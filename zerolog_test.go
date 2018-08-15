package bench

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func BenchmarkZerologTextFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-zerolog")
	if err != nil {
		b.Fatal(err)
	}

	logger := zerolog.New(tmpfile).With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkZerologTextStd(b *testing.B) {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkZerologJSONFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-zerolog")
	if err != nil {
		b.Fatal(err)
	}

	logger := zerolog.New(tmpfile).With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkZerologJSONStd(b *testing.B) {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")
		}
	})
}
