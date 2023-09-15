package test

import (
	"testing"

	"github.com/Arian-p1/ddocs-go/internal"
)

func TestComprssion(t *testing.T) {
  compressed, err := internal.Compress("test string to compress")
  if err != nil {
    t.Fatalf("failed to compress: %v", err)
  }
  t.Logf("compressed data: %v", compressed)
  decompressed, err := internal.Decompress(compressed)
  if err != nil {
    t.Fatalf("failed to decompress: %v", err)
  }
  t.Logf("decompressed data: %v", decompressed)
}

func TestFile(t *testing.T) {
  test_map := map[string]string{
    "test key": "test value",
  }
  err := internal.WriteToFile(test_map)
  if err != nil {
    t.Fatalf("failed to write: %v", err)
  }
  data, err := internal.ReadFile()
  if err != nil {
    t.Fatalf("failed to read: %v", err)
  }
  t.Logf("data: %v", data)
}
