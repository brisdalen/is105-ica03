
package pipe

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func HexReturn(s string) string {
	txt := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(txt)))
	hex.Encode(dst, txt)
	return fmt.Sprintf("%s", dst)
}

func Base64Return(s string) string {
	b64 := base64.StdEncoding.EncodeToString([]byte(s))
	txt := []byte(b64)
	dst := make([]byte, hex.EncodedLen(len(txt)))
	hex.Encode(txt, dst)
	return fmt.Sprintf("%s", dst)
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