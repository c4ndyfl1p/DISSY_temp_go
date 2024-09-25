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
	// return []byte , no of bytes read
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

func encrypt(aesKey []byte, plaintext []byte) []byte {

	cipher1, err := aes.NewCipher(aesKey)
	if err != nil {
		fmt.Println("err: with aes.NewCipher", err)
	}

	// hardcoding the ct size for now, might need to change it later
	temp_ct := make([]byte, Blocksize*6)

	for i := 0; i <= 16; i += Blocksize {
		//fmt.Println(i, i+Blocksize)
		temp_pt := plaintext[i : i+Blocksize]
		cipher1.Encrypt(temp_ct[i:i+Blocksize], temp_pt)

		fmt.Println("temp_pt: ", temp_pt, len(temp_pt))
		fmt.Println("temp_ct:", temp_ct[i:i+Blocksize], len(temp_ct))
	}

	return temp_ct
}

func decrypt(aeskey []byte, ct2 []byte) []byte {
	// for decryption
	//temp_ct := make([]byte, Blocksize*6)
	cipher1, err := aes.NewCipher(aeskey)
	if err != nil {
		fmt.Println("err: with aes.NewCipher", err)
	}

	recoverdPt2 := make([]byte, Blocksize*6)

	fmt.Println(recoverdPt2)

	for i := 0; i <= 16; i += Blocksize {
		//fmt.Println(i, i+Blocksize)
		cipher1.Decrypt(recoverdPt2[i:i+Blocksize], ct2[i:i+Blocksize])

		fmt.Println("recovered_pt2:", recoverdPt2[i:i+Blocksize])
	}

	return recoverdPt2

}

func writeToFile(aesKey []byte, ciphertextFile string) string {
	return ""
}

func main() {
	fmt.Println("Hello, World!")

	// read plaintext from file and load into bytes
	pt_from_file, numberOfBytesRead := readFromFile("example.txt")
	fmt.Println(pt_from_file, numberOfBytesRead)

	fmt.Println("------------------------------")

	// generate aes key
	aeskey := keyGen()
	fmt.Println(aeskey)

	// encrypt the plaintext(byte)
	ct2 := encrypt(aeskey, pt_from_file)
	fmt.Println(ct2)

	// decrypt the ciphertext(byte)
	recoveredPt := decrypt(aeskey, ct2)
	print("-------------------------------------")
	fmt.Println(string(recoveredPt))

	//write to ciphertext to file
}
