package b64

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Base64Return(s string) string {
	data, _ := hex.DecodeString(s)
	b64 := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf(b64)
}

func Base64ReturnBytes(b []byte) []byte {
	sEnc := base64.StdEncoding.EncodeToString(b)
	return []byte(sEnc)
}