
package pipe

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"log"
)

func HexReturn(s string) string {
	return fmt.Sprintf("%X", s)
}

func Base64Return(s string) string {
	return fmt.Sprintln(base64.StdEncoding.EncodeToString([]byte(s)))
}

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


	return fmt.Sprintln(buffer)
}