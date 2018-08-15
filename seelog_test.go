package bench

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/cihub/seelog"
)

func BenchmarkSeelogTextFile(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark-seelog")
	if err != nil {
		b.Fatal(err)
	}

	logger, err := log.LoggerFromWriterWithMinLevelAndFormat(tmpfile, log.TraceLvl, "%Time %Level %Msg")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		defer logger.Flush()
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}

func BenchmarkSeelogTextStd(b *testing.B) {
	logger, err := log.LoggerFromWriterWithMinLevelAndFormat(os.Stderr, log.TraceLvl, "%Time %Level %Msg")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		defer logger.Flush()
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
}
