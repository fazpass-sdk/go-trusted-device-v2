package gotdv2

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"

	. "github.com/fazpass-sdk/go-trusted-device-v2/dto"
	"os"
)

type Fazpass struct {
	PrivateKey  *rsa.PrivateKey
	MerchantKey string
	BaseUrl     string
}

func (f Fazpass) CheckDevice(picId string, meta string, appId string) (*Device, error) {
	device := &Device{}
	if picId == "" || meta == "" || appId == "" {
		return device, errors.New("pic id, meta or app id cannot be empty")
	}
	url := fmt.Sprintf("%s/check", f.BaseUrl)
	// Create request body
	requestBody := CheckRequest{
		PicId:         picId,
		Meta:          meta,
		MerchantAppId: appId,
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	// Add authorization header to the req
	bearerToken := fmt.Sprintf("Bearer %s", f.MerchantKey) // assuming MerchantKey is the bearer token
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status code %d", resp.StatusCode)
	}
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	unwrapBase64, err := base64.StdEncoding.DecodeString(response.Data.Meta)
	decrypted, err := decryptWithPrivateKey(unwrapBase64, f.PrivateKey)
	err = json.Unmarshal(decrypted, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (f Fazpass) CheckAsyncDevice(picId string, meta string, appId string) (<-chan *Device, <-chan error) {
	deviceChan := make(chan *Device, 1)
	errChan := make(chan error, 1)
	go func() {
		device, err := f.CheckDevice(picId, meta, appId)
		if err != nil {
			errChan <- err
			return
		}
		deviceChan <- device
	}()
	return deviceChan, errChan
}

func (f Fazpass) EnrollDevice(picId string, meta string, appId string) (*Device, error) {
	device := &Device{}
	if picId == "" || meta == "" || appId == "" {
		return device, errors.New("pic id, meta or app id cannot be empty")
	}
	url := fmt.Sprintf("%s/enroll", f.BaseUrl)
	// Create request body
	requestBody := EnrollRequest{
		PicId:         picId,
		Meta:          meta,
		MerchantAppId: appId,
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	// Add authorization header to the req
	bearerToken := fmt.Sprintf("Bearer %s", f.MerchantKey) // assuming MerchantKey is the bearer token
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status code %d", resp.StatusCode)
	}
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	unwrapBase64, err := base64.StdEncoding.DecodeString(response.Data.Meta)
	decrypted, err := decryptWithPrivateKey(unwrapBase64, f.PrivateKey)
	err = json.Unmarshal(decrypted, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (f Fazpass) EnrollAsyncDevice(picId string, meta string, appId string) (<-chan *Device, <-chan error) {
	deviceChan := make(chan *Device, 1)
	errChan := make(chan error, 1)
	go func() {
		device, err := f.EnrollDevice(picId, meta, appId)
		if err != nil {
			errChan <- err
			return
		}
		deviceChan <- device
	}()

	return deviceChan, errChan
}

func (f Fazpass) ValidateDevice(fazpassId string, meta string, appId string) (*Device, error) {
	device := &Device{}
	if fazpassId == "" || meta == "" || appId == "" {
		return device, errors.New("fazpass id, meta or app id cannot be empty")
	}
	url := fmt.Sprintf("%s/validate", f.BaseUrl)
	// Create request body
	requestBody := ValidateRequest{
		FazpassId:     fazpassId,
		Meta:          meta,
		MerchantAppId: appId,
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	// Add authorization header to the req
	bearerToken := fmt.Sprintf("Bearer %s", f.MerchantKey) // assuming MerchantKey is the bearer token
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status code %d", resp.StatusCode)
	}
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	unwrapBase64, err := base64.StdEncoding.DecodeString(response.Data.Meta)
	decrypted, err := decryptWithPrivateKey(unwrapBase64, f.PrivateKey)
	err = json.Unmarshal(decrypted, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (f Fazpass) ValidateAsyncDevice(fazpassId string, meta string, appId string) (<-chan *Device, <-chan error) {
	deviceChan := make(chan *Device, 1)
	errChan := make(chan error, 1)
	go func() {
		device, err := f.ValidateDevice(fazpassId, meta, appId)
		if err != nil {
			errChan <- err
			return
		}
		deviceChan <- device
	}()

	return deviceChan, errChan
}

func (f Fazpass) RemoveDevice(fazpassId string, meta string, appId string) (*Device, error) {
	device := &Device{}
	if fazpassId == "" || meta == "" || appId == "" {
		return device, errors.New("fazpass id, meta or app id cannot be empty")
	}
	url := fmt.Sprintf("%s/remove", f.BaseUrl)
	// Create request body
	requestBody := ValidateRequest{
		FazpassId:     fazpassId,
		Meta:          meta,
		MerchantAppId: appId,
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	// Add authorization header to the req
	bearerToken := fmt.Sprintf("Bearer %s", f.MerchantKey) // assuming MerchantKey is the bearer token
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status code %d", resp.StatusCode)
	}
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	unwrapBase64, err := base64.StdEncoding.DecodeString(response.Data.Meta)
	decrypted, err := decryptWithPrivateKey(unwrapBase64, f.PrivateKey)
	err = json.Unmarshal(decrypted, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (f Fazpass) RemoveAsyncDevice(fazpassId string, meta string, appId string) (<-chan *Device, <-chan error) {
	deviceChan := make(chan *Device, 1)
	errChan := make(chan error, 1)
	go func() {
		device, err := f.RemoveDevice(fazpassId, meta, appId)
		if err != nil {
			errChan <- err
			return
		}
		deviceChan <- device
	}()

	return deviceChan, errChan
}

func Initialize(privateKeyPath string, baseUrl string, merchantKey string) (TrustedDevice, error) {
	var privateKey *rsa.PrivateKey
	f := Fazpass{}
	if privateKeyPath == "" || baseUrl == "" || merchantKey == "" {
		return f, errors.New("parameter cannot be empty")
	}
	private, errFile := os.ReadFile(privateKeyPath)
	if errFile != nil {
		return f, errors.New("file not found")
	}
	privateKey, _ = bytesToPrivateKey(private)
	f.BaseUrl = baseUrl
	f.MerchantKey = merchantKey
	f.PrivateKey = privateKey
	return f, nil
}

func bytesToPrivateKey(priv []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	b := block.Bytes
	key, err := x509.ParsePKCS1PrivateKey(b)
	return key, err
}

func decryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	return plaintext, err
}
