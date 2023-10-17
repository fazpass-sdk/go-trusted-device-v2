package gotdv2

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"os"
)

type Fazpass struct {
	PrivateKey *rsa.PrivateKey
}

func (f *Fazpass) Extract(meta string) (Meta, error) {
	m := Meta{}
	unwrapBase64, err := base64.StdEncoding.DecodeString(meta)
	decrypted, err := decryptWithPrivateKey(unwrapBase64, f.PrivateKey)
	if err != nil {
		return m, errors.New("invalid meta or key")
	}
	jsonString := string(decrypted)
	err = json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func Initialize(privateKeyPath string) (Fazpass, error) {
	var privateKey *rsa.PrivateKey
	f := Fazpass{}
	if privateKeyPath == "" {
		return f, errors.New("parameter cannot be empty")
	}
	private, errFile := os.ReadFile(privateKeyPath)
	if errFile != nil {
		return f, errors.New("file not found")
	}
	privateKey, _ = bytesToPrivateKey(private)
	f.PrivateKey = privateKey
	return f, nil
}

func bytesToPrivateKey(priv []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	b := block.Bytes
	var err error
	key, err := x509.ParsePKCS1PrivateKey(b)

	return key, err
}

func decryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	return plaintext, err
}

func bytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	b := block.Bytes
	ifc, _ := x509.ParsePKIXPublicKey(b)
	key, _ := ifc.(*rsa.PublicKey)
	return key
}

func encryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	ciphertext, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	return ciphertext

}
