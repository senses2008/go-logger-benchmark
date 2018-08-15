package bench

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/op/go-logging"
)

func BenchmarkGologgingTextFile(b *testing.B) {

	tmpfile, err := ioutil.TempFile("", "benchmark-gologging")
	if err != nil {
		b.Fatal(err)
	}

	logger := log.MustGetLogger("")
	subBackend := log.NewLogBackend(tmpfile, "", 0)
	formatter := log.MustStringFormatter("%{time:2006-01-02T15:04:05Z07:00} %{level} %{message}")
	backend := log.NewBackendFormatter(subBackend, formatter)
	leveled := log.AddModuleLevel(backend)
	logger.SetBackend(leveled)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkGologgingTextStd(b *testing.B) {
	logger := log.MustGetLogger("")
	subBackend := log.NewLogBackend(os.Stderr, "", 0)
	formatter := log.MustStringFormatter("%{time:2006-01-02T15:04:05Z07:00} %{level} %{message}")
	backend := log.NewBackendFormatter(subBackend, formatter)
	leveled := log.AddModuleLevel(backend)
	logger.SetBackend(leveled)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}
