package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

const Blocksize = 16

func keyGen() []byte {
	aesKey := make([]byte, Blocksize)
	_, err := rand.Read(aesKey)
	if err != nil {
		fmt.Println("error: reading rand.Read", err)
		return []byte(err.Error())
	}

	return aesKey
}

func readFromFile(filename string) ([]byte, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []byte(err.Error()), 0
	}
	defer file.Close()

	// Create a byte slice to hold the file contents
	buffer := make([]byte, 1024) // 1KB buffer size
	var numberOfBytesReadG int
	numberOfBytesReadG = 0

	for {
		// Read bytes into the buffer
		numberOfBytesRead, err := file.Read(buffer)

		// Check for end of file or error
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			return []byte(err.Error()), 0
		}

		// If we reach EOF, break out of the loop
		if err == io.EOF {
			break
		}

		// Process the read bytes
		fmt.Printf("Read %d bytes: %s\n", numberOfBytesRead, buffer[:numberOfBytesRead])
		numberOfBytesReadG = numberOfBytesRead

	}
	return buffer[:numberOfBytesReadG], numberOfBytesReadG

}

func encryptToFile(aesKey []byte, plaintextFile string) string {
	return ""
}

func decryptToFile(aesKey []byte, ciphertextFile string) string {
	return ""
}

func main() {
	fmt.Println("Hello, World!")
	// Open the file for reading

	pt_from_file, numberOfBytesRead := readFromFile("example.txt")
	fmt.Println(pt_from_file, numberOfBytesRead)

	fmt.Println("------------------------------")
	aeskey := keyGen()
	fmt.Println(aeskey)

	pt := make([]byte, Blocksize)
	pt[0] = 1
	pt[1] = 2
	pt[15] = 15

	fmt.Println("plaintext:", pt)

	cipher1, err := aes.NewCipher(aeskey)
	if err != nil {
		fmt.Println("err: with aes.NewCipher", err)
	}

	ct := make([]byte, Blocksize)
	temp_ct := make([]byte, Blocksize*6)

	for i := 0; i <= 16; i += Blocksize {
		//fmt.Println(i, i+Blocksize)
		temp_pt := pt_from_file[i : i+Blocksize]
		cipher1.Encrypt(temp_ct[i:i+Blocksize], temp_pt)

		fmt.Println("temp_pt: ", temp_pt, len(temp_pt))
		fmt.Println("temp_ct:", temp_ct[i:i+Blocksize], len(temp_ct))
	}

	// for decryption

	recoverdPt2 := make([]byte, Blocksize*6)

	fmt.Println(recoverdPt2)

	for i := 0; i <= 16; i += Blocksize {
		//fmt.Println(i, i+Blocksize)
		cipher1.Decrypt(recoverdPt2[i:i+Blocksize], temp_ct[i:i+Blocksize])

		fmt.Println("recovered_pt2:", recoverdPt2[i:i+Blocksize])
	}

	cipher1.Encrypt(ct, pt)

	fmt.Println(ct)

	recoverdPt := make([]byte, Blocksize)
	cipher1.Decrypt(recoverdPt, ct)
	fmt.Println(recoverdPt)
	print("-------------------------------------")
	fmt.Println(string(recoverdPt2))
}
