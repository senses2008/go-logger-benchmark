package bench

import (
	"io/ioutil"
	"os"
	"testing"

	log "gopkg.in/inconshreveable/log15.v2"
)

func BenchmarkLog15TextFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-log15")
	if err != nil {
		b.Fatal(err)
	}

	logger := log.New()
	logger.SetHandler(log.StreamHandler(tmpfile, log.LogfmtFormat()))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkLog15TextStd(b *testing.B) {
	logger := log.New()
	logger.SetHandler(log.StreamHandler(os.Stderr, log.LogfmtFormat()))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkLog15JSONFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-log15")
	if err != nil {
		b.Fatal(err)
	}

	logger := log.New()
	logger.SetHandler(log.StreamHandler(tmpfile, log.JsonFormat()))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})
}

func BenchmarkLog15JSONStd(b *testing.B) {

	logger := log.New()
	logger.SetHandler(log.StreamHandler(os.Stderr, log.JsonFormat()))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})
}
