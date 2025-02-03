package main

import (
	"fmt"

	"github.com/umang404sharma/cryptopals/internal/encoding"
)

func main() {
	base64String, err := encoding.HexToBase64("a49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		fmt.Println("Error converting the Hex code to Base64:\n", err)
	}
	fmt.Println("Base64 String: ", base64String)
}
