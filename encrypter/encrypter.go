package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Переменная окружения KEY пуста")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	// Инициализация блочного шифра AES из ключа enc.Key
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	// Включение режиме GCM
	// Обеспечивает не только шифрование но аутентификацию данных через криптографический тег(защита от подмены)
	aesGСM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	// Number used once - криптографический параметр, который должен быть уникальным для каждой операции шифрования одним и тем же ключем
	nonce := make([]byte, aesGСM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce) // Криптостойкая случайность из crypto/rand
	if err != nil {
		panic(err.Error())
	}
	return aesGСM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encrypter) Decrypt(encryptedStr []byte) []byte {
	// Инициализация блочного шифра AES из ключа enc.Key
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	// Включение режиме GCM
	// Обеспечивает не только шифрование но аутентификацию данных через криптографический тег(защита от подмены)
	aesGСM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGСM.NonceSize()
	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGСM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
