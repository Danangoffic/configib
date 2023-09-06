package Tools

import (
	"crypto/des"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image"
	"io"
	"log"
	"strings"
)

const (
	KEY_STRING string = "3816c875da2f2540"
)

type appFactory struct {
}

type AppFactory interface {
	IsEmpty(v interface{}) bool
	GetDeviceLockToken(data, salt string) map[string]string
	EncryptData(source string) string
	DecryptData(source string) string
}

func NewAppFactory() AppFactory {
	return &appFactory{}
}

func (a *appFactory) IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	if str, ok := v.(string); ok {
		return str == "" || str == "null" || str == "undefined"
	}

	return false
}

func (a *appFactory) GetDeviceLockToken(data, salt string) map[string]string {
	hash := sha512.Sum512([]byte(data + salt))
	// hash2:=sha512.Sum([]byte(salt))
	newmap := make(map[string]string)
	newmap["hash"] = hex.EncodeToString(hash[:])
	newmap["salt"] = salt

	return newmap
}

func (a *appFactory) EncryptData(source string) string {
	return encryptData(source)
}

func encryptData(source string) string {
	// Get our secret key
	key := getKey()

	// Create the cipher
	des, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Initialize the des for encryption
	var iv = []byte(source)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	var dest []byte
	des.Encrypt(dest, iv)

	// Encrypted the cleartext
	return getString(dest)
}

func (a *appFactory) DecryptData(source string) string {
	return decryptData(source)
}

func decryptData(source string) string {
	// Get our secret key
	key := getKey()

	// Create the cipher
	des, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Initialize the des for encryption
	var iv = []byte(source)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	var dest []byte
	des.Decrypt(dest, iv)

	// Encrypted the cleartext
	return getString(dest)
}

func getKey() (key []byte) {
	// Get the key bytes from the string
	bytes := getBytes(KEY_STRING)

	// Create a DESKeySpec object from the key bytes

	desKeySpec, err := des.NewCipher(bytes)
	if err != nil {
		log.Println("Error generating encryption key ", err)
	}
	keys := make([]byte, desKeySpec.BlockSize())
	rand.Read(keys)

	// Create a SecretKeyFactory object
	// secretKeyFactory := keys

	// Generate the secret key
	// key, err = secretKeyFactory.GenerateSecret(desKeySpec)
	// if err != nil {
	// 	return nil
	// }

	return keys
}

func getBytes(str string) []byte {
	return []byte(str)
}
func getString(bytes []byte) string {
	sb := strings.Builder{}
	for i := 0; i < len(bytes); i++ {
		s := fmt.Sprintf("%02x", bytes[i])
		sb.WriteString(s)
	}
	return sb.String()
}

func flipImage(imageSource string, dimension string) string {
	// Decode the base64 image
	bytes, err := base64.StdEncoding.DecodeString(imageSource)

	// Create a new image
	img, err := image.New(bytes)
	if err != nil {
		panic(err)
	}

	// Flip the image
	if dimension == "vertical" {
		img = image.FlipV(img)
	} else if dimension == "horizontal" {
		img = image.FlipH(img)
	}

	// Encode the image as base64
	encodedImage := base64.StdEncoding.EncodeToString(img.Bytes())

	// Remove any newline characters from the base64 string
	encodedImage = strings.Replace(encodedImage, "\n", "", -1)
	encodedImage = strings.Replace(encodedImage, "\r", "", -1)

	return encodedImage
}
