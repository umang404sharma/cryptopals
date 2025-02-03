package utils

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexStr string) (string, error) {

	rawBytes, err := hex.DecodeString(hexStr)

	if err != nil {
		// fmt.Println("Error decoding the hex:", err)
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(rawBytes)

	return base64String, nil
}
