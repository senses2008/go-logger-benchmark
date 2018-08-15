package bench

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func BenchmarkLogrusTextFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-logrus")
	if err != nil {
		b.Fatal(err)
	}

	logger := log.New()
	logger.Formatter = &log.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	}
	logger.Out = tmpfile
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkLogrusTextStd(b *testing.B) {

	logger := log.New()
	logger.Formatter = &log.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	}
	logger.Out = os.Stderr
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkLogrusJSONFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-logrus")
	if err != nil {
		b.Fatal(err)
	}

	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}
	logger.Out = tmpfile
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.WithFields(log.Fields{
				"rate": "15",
				"low":  16,
				"high": 123.2,
			}).Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkLogrusJSONStd(b *testing.B) {
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}
	logger.Out = os.Stderr
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.WithFields(log.Fields{
				"rate": "15",
				"low":  16,
				"high": 123.2,
			}).Info("The quick brown fox jumps over the lazy dog")
		}
	})
}
