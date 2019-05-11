package gzip

import (
	"bytes"
	"compress/gzip"
	"log"
)

func GZipReturn(s string) string {
	var buffer bytes.Buffer
	gZipWriter := gzip.NewWriter(&buffer)

	_, err := gZipWriter.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	err = gZipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	return buffer.String()
}

func GZipReturnBytes(b []byte) []byte {
	var buffer bytes.Buffer
	gZipWriter := gzip.NewWriter(&buffer)

	_, err := gZipWriter.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	err = gZipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	return buffer.Bytes()
}