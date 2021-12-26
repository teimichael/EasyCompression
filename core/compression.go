package core

import (
	"fmt"
	"github.com/klauspost/compress/zstd"
	"io"
	"log"
	"os"
)

func Compress(path, outputName string) {
	data, err := os.ReadFile(path)

	// File reading failed
	if err != nil {
		log.Fatal(err)
		return
	}

	compressed := compress(data)

	if err = os.WriteFile(outputName, compressed, 0666); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully compressed ", path)
		showCompressionRatio(len(data), len(compressed))
	}
}

func Decompress(path, outputName string) {
	data, err := os.ReadFile(path)

	// File reading failed
	if err != nil {
		log.Fatal(err)
		return
	}

	decompressed, err := decompress(data)

	// Decompression failed
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = os.WriteFile(outputName, decompressed, 0666); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully decompressed ", path)
		showCompressionRatio(len(decompressed), len(data))
	}
}

func compress(src []byte) []byte {
	var encoder, _ = zstd.NewWriter(nil)
	return encoder.EncodeAll(src, make([]byte, 0, len(src)))
}

func decompress(src []byte) ([]byte, error) {
	var decoder, _ = zstd.NewReader(nil)
	return decoder.DecodeAll(src, nil)
}

func compressStream(in io.Reader, out io.Writer) error {
	enc, err := zstd.NewWriter(out)
	if err != nil {
		return err
	}
	_, err = io.Copy(enc, in)
	if err != nil {
		enc.Close()
		return err
	}
	return enc.Close()
}

func decompressStream(in io.Reader, out io.Writer) error {
	d, err := zstd.NewReader(in)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = io.Copy(out, d)
	return err
}

func showCompressionRatio(uncompressed, compressed int) {
	fmt.Println("Uncompressed size: ", uncompressed)
	fmt.Println("Compressed size: ", compressed)
	fmt.Println("Compression ratio: ", float32(uncompressed)/float32(compressed))
}
