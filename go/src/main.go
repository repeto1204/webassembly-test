package main

import (
   "bytes"
   "encoding/hex"
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "fmt"
   "syscall/js"
)

var key = []byte("1q2w3edfgtyhjuiklop09oikjuyhgtrf")
var iv = []byte("9875632654856956")

func main() {

    c := make(chan struct{}, 0)
    println("WASM Go Initialized")
    js.Global().Set("WebassemblyEncrypt", js.FuncOf(Encrypt))
    js.Global().Set("WebassemblyDecrypt", js.FuncOf(Decrypt))
    js.Global().Set("WebassemblyInjection", js.FuncOf(Injection))
    <-c

}

func Injection(this js.Value, i []js.Value) interface{} {
    doc := js.Global().Get("document")
    body := doc.Call("getElementById", "testDiv")
    body.Set("innerHTML", i[0].String())

    return ""
}

func Encrypt(this js.Value, i []js.Value) interface{} {
   data, err:= AesCBCEncrypt([]byte(i[0].String()))
   if err != nil {
       return err
   }
   fmt.Println("ENCRYPT BASE64")
   fmt.Println(base64.StdEncoding.EncodeToString(data));
   fmt.Println("ENCRYPT HEX")
   fmt.Println(hex.EncodeToString(data));
   return js.ValueOf(base64.StdEncoding.EncodeToString(data));
//    return js.ValueOf(hex.EncodeToString(data))
}

func AesCBCEncrypt(plainText []byte) ([]byte, error) {

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    //fill the original
    blockSize := block.BlockSize()
    plainText = PKCS5Padding(plainText, blockSize)

    cipherText := make([]byte,blockSize+len(plainText))


    //block size and initial vector size must be the same
    mode := cipher.NewCBCEncrypter(block,iv)
    mode.CryptBlocks(cipherText[blockSize:],plainText)

    return cipherText, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
//    padding := blockSize - len(ciphertext) % blockSize
//    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//    return append(ciphertext, padtext...)
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
   length := len(origData)
   unpadding := int(origData[length-1])
   return origData[:(length - unpadding)]
}

func PKCS5UnPadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func AesCBCDecrypt(cipherText []byte) ([]byte, error) {
   block, err := aes.NewCipher(key)
   if err != nil {
       panic(err)
   }

   blockSize := block.BlockSize()

   if len(cipherText) < blockSize {
       panic("ciphertext too short")
   }

   cipherText = cipherText[blockSize:]

   if len(cipherText)%blockSize != 0 {
       panic("ciphertext is not a multiple of the block size")
   }

   mode := cipher.NewCBCDecrypter(block, iv)

   // CryptBlocks can work in-place if the two arguments are the same.
   mode.CryptBlocks(cipherText, cipherText)

   cipherText = PKCS5UnPadding(cipherText)
   return cipherText,nil
}

func Decrypt(this js.Value, i []js.Value) interface{} {
    cipherText, _ := base64.StdEncoding.DecodeString(i[0].String())
    data,err := AesCBCDecrypt(cipherText);
    if err != nil {
        return err
    }

    return string(data)
}