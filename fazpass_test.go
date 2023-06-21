package gotdv2

import (
	"encoding/base64"
	"encoding/json"
	. "github.com/fazpass-sdk/go-trusted-device-v2/dto"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// Membuat mock HTTP server untuk pengujian
func createMockServer(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}

const finalMeta = "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw=="

func TestCheckDevice(t *testing.T) {
	// Membuat mock server
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/check" {
			t.Errorf("expected '/check' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody CheckRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.PicId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.PicId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: finalMeta},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	mockServerError := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/check" {
			t.Errorf("expected '/check' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody CheckRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.PicId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.PicId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: finalMeta},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	// Mengatur base URL dan merchant key untuk pengujian
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)

	// Menjalankan pengujian
	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	t.Run("param empty", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServer.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		_, err := fazpass.CheckDevice("", "", "")
		assert.EqualError(t, err, "pic id, meta or app id cannot be empty")
	})

	t.Run("failed to create HTTP request", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     "failed url",
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		picID := "mockPicID"
		meta := "mockMeta"
		appID := "mockAppID"

		// Memeriksa error saat melakukan CheckDevice
		_, err := fazpass.CheckDevice(picID, meta, appID)
		assert.Error(t, err, "expected error when creating HTTP request")
	})

	t.Run("failed", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServerError.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		_, err := fazpass.CheckDevice(picID, meta, appID)
		assert.Error(t, err, "server returned status code 406")

	})

	t.Run("success", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServer.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		device, err := fazpass.CheckDevice(picID, meta, appID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		// Memeriksa hasil pengujian
		expectedDevice := &Device{}
		expectedUnwrapMeta, _ := base64.StdEncoding.DecodeString("c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==")
		expectedDecrypted, _ := decryptWithPrivateKey(expectedUnwrapMeta, fazpass.PrivateKey)
		_ = json.Unmarshal(expectedDecrypted, &expectedDevice)
		if !isEqualDevice(device, expectedDevice) {
			t.Errorf("expected device %+v, got %+v", expectedDevice, device)
		}
	})
}

func TestCheckAsyncDevice(t *testing.T) {
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/check" {
			t.Errorf("expected '/check' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody CheckRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.PicId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.PicId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==", // base64 encoded string
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)
	fazpass := Fazpass{
		BaseUrl:     mockServer.URL,
		MerchantKey: "mockMerchantKey",
		PrivateKey:  privateKey,
	}

	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	deviceChan, errChan := fazpass.CheckAsyncDevice(picID, meta, appID)

	select {
	case device := <-deviceChan:
		// Verifikasi hasil device
		expectedDevice := &Device{FazpassId: "fazpass_id"} // Ganti dengan nilai yang diharapkan
		if !isEqualDevice(device, expectedDevice) {
			t.Errorf("expected device %+v, got %+v", expectedDevice, device)
		}
	case err := <-errChan:
		// Verifikasi error
		t.Errorf("unexpected error: %v", err)
	case <-time.After(5 * time.Second):
		t.Error("timeout occurred")
	}
}

func TestErrorCheckAsyncDevice(t *testing.T) {
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/check" {
			t.Errorf("expected '/check' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody CheckRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.PicId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.PicId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==", // base64 encoded string
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)
	fazpass := Fazpass{
		BaseUrl:     mockServer.URL,
		MerchantKey: "mockMerchantKey",
		PrivateKey:  privateKey,
	}

	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	_, errChan := fazpass.CheckAsyncDevice(picID, meta, appID)

	select {
	case err := <-errChan:
		// Verifikasi error
		assert.Error(t, err, "server returned status code 406")
	case <-time.After(5 * time.Second):
		t.Error("timeout occurred")
	}
}

func TestEnrollDevice(t *testing.T) {
	// Membuat mock server
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/enroll" {
			t.Errorf("expected '/enroll' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody EnrollRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.PicId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.PicId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: finalMeta},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	mockServerError := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/enroll" {
			t.Errorf("expected '/enroll' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody EnrollRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.PicId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.PicId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: finalMeta},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	// Mengatur base URL dan merchant key untuk pengujian
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)

	// Menjalankan pengujian
	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	t.Run("param empty", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServer.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		_, err := fazpass.EnrollDevice("", "", "")
		assert.EqualError(t, err, "pic id, meta or app id cannot be empty")
	})

	t.Run("failed to create HTTP request", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     "failed url",
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		picID := "mockPicID"
		meta := "mockMeta"
		appID := "mockAppID"

		// Memeriksa error saat melakukan CheckDevice
		_, err := fazpass.EnrollDevice(picID, meta, appID)
		assert.Error(t, err, "expected error when creating HTTP request")
	})

	t.Run("failed", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServerError.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		_, err := fazpass.EnrollDevice(picID, meta, appID)
		assert.Error(t, err, "server returned status code 406")

	})

	t.Run("success", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServer.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		device, err := fazpass.EnrollDevice(picID, meta, appID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		// Memeriksa hasil pengujian
		expectedDevice := &Device{}
		expectedUnwrapMeta, _ := base64.StdEncoding.DecodeString("c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==")
		expectedDecrypted, _ := decryptWithPrivateKey(expectedUnwrapMeta, fazpass.PrivateKey)
		_ = json.Unmarshal(expectedDecrypted, &expectedDevice)
		if !isEqualDevice(device, expectedDevice) {
			t.Errorf("expected device %+v, got %+v", expectedDevice, device)
		}
	})
}

func TestEnrollAsyncDevice(t *testing.T) {
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/enroll" {
			t.Errorf("expected '/enroll' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody EnrollRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.PicId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.PicId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==", // base64 encoded string
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)
	fazpass := Fazpass{
		BaseUrl:     mockServer.URL,
		MerchantKey: "mockMerchantKey",
		PrivateKey:  privateKey,
	}

	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	deviceChan, errChan := fazpass.EnrollAsyncDevice(picID, meta, appID)

	select {
	case device := <-deviceChan:
		// Verifikasi hasil device
		expectedDevice := &Device{FazpassId: "fazpass_id"} // Ganti dengan nilai yang diharapkan
		if !isEqualDevice(device, expectedDevice) {
			t.Errorf("expected device %+v, got %+v", expectedDevice, device)
		}
	case err := <-errChan:
		// Verifikasi error
		t.Errorf("unexpected error: %v", err)
	case <-time.After(5 * time.Second):
		t.Error("timeout occurred")
	}
}

func TestErrorEnrollAsyncDevice(t *testing.T) {
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/enroll" {
			t.Errorf("expected '/enroll' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody EnrollRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.PicId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.PicId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==", // base64 encoded string
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)
	fazpass := Fazpass{
		BaseUrl:     mockServer.URL,
		MerchantKey: "mockMerchantKey",
		PrivateKey:  privateKey,
	}

	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	_, errChan := fazpass.EnrollAsyncDevice(picID, meta, appID)

	select {
	case err := <-errChan:
		// Verifikasi error
		assert.Error(t, err, "server returned status code 406")
	case <-time.After(5 * time.Second):
		t.Error("timeout occurred")
	}
}

func TestValidateDevice(t *testing.T) {
	// Membuat mock server
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/validate" {
			t.Errorf("expected '/validate' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody ValidateRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.FazpassId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.FazpassId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: finalMeta},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	mockServerError := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/validate" {
			t.Errorf("expected '/validate' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody ValidateRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.FazpassId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.FazpassId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: finalMeta},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	// Mengatur base URL dan merchant key untuk pengujian
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)

	// Menjalankan pengujian
	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	t.Run("param empty", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServer.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		_, err := fazpass.ValidateDevice("", "", "")
		assert.EqualError(t, err, "fazpass id, meta or app id cannot be empty")
	})

	t.Run("failed to create HTTP request", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     "failed url",
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		picID := "mockPicID"
		meta := "mockMeta"
		appID := "mockAppID"

		// Memeriksa error saat melakukan CheckDevice
		_, err := fazpass.ValidateDevice(picID, meta, appID)
		assert.Error(t, err, "expected error when creating HTTP request")
	})

	t.Run("failed", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServerError.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		_, err := fazpass.ValidateDevice(picID, meta, appID)
		assert.Error(t, err, "server returned status code 406")

	})

	t.Run("success", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServer.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		device, err := fazpass.ValidateDevice(picID, meta, appID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		// Memeriksa hasil pengujian
		expectedDevice := &Device{}
		expectedUnwrapMeta, _ := base64.StdEncoding.DecodeString("c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==")
		expectedDecrypted, _ := decryptWithPrivateKey(expectedUnwrapMeta, fazpass.PrivateKey)
		_ = json.Unmarshal(expectedDecrypted, &expectedDevice)
		if !isEqualDevice(device, expectedDevice) {
			t.Errorf("expected device %+v, got %+v", expectedDevice, device)
		}
	})
}

func TestValidateAsyncDevice(t *testing.T) {
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/validate" {
			t.Errorf("expected '/validate' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody ValidateRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.FazpassId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.FazpassId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==", // base64 encoded string
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)
	fazpass := Fazpass{
		BaseUrl:     mockServer.URL,
		MerchantKey: "mockMerchantKey",
		PrivateKey:  privateKey,
	}

	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	deviceChan, errChan := fazpass.ValidateAsyncDevice(picID, meta, appID)

	select {
	case device := <-deviceChan:
		// Verifikasi hasil device
		expectedDevice := &Device{FazpassId: "fazpass_id"} // Ganti dengan nilai yang diharapkan
		if !isEqualDevice(device, expectedDevice) {
			t.Errorf("expected device %+v, got %+v", expectedDevice, device)
		}
	case err := <-errChan:
		// Verifikasi error
		t.Errorf("unexpected error: %v", err)
	case <-time.After(5 * time.Second):
		t.Error("timeout occurred")
	}
}

func TestErrorValidateAsyncDevice(t *testing.T) {
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/validate" {
			t.Errorf("expected '/validate' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody ValidateRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.FazpassId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.FazpassId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==", // base64 encoded string
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)
	fazpass := Fazpass{
		BaseUrl:     mockServer.URL,
		MerchantKey: "mockMerchantKey",
		PrivateKey:  privateKey,
	}

	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	_, errChan := fazpass.ValidateAsyncDevice(picID, meta, appID)

	select {
	case err := <-errChan:
		// Verifikasi error
		assert.Error(t, err, "server returned status code 406")
	case <-time.After(5 * time.Second):
		t.Error("timeout occurred")
	}
}

func TestRemoveDevice(t *testing.T) {
	// Membuat mock server
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/remove" {
			t.Errorf("expected '/remove' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody RemoveRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.FazpassId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.FazpassId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: finalMeta},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	mockServerError := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/remove" {
			t.Errorf("expected '/remove' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody RemoveRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.FazpassId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.FazpassId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: finalMeta},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	// Mengatur base URL dan merchant key untuk pengujian
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)

	// Menjalankan pengujian
	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	t.Run("param empty", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServer.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		_, err := fazpass.RemoveDevice("", "", "")
		assert.EqualError(t, err, "fazpass id, meta or app id cannot be empty")
	})

	t.Run("failed to create HTTP request", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     "failed url",
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		picID := "mockPicID"
		meta := "mockMeta"
		appID := "mockAppID"

		// Memeriksa error saat melakukan CheckDevice
		_, err := fazpass.RemoveDevice(picID, meta, appID)
		assert.Error(t, err, "expected error when creating HTTP request")
	})

	t.Run("failed", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServerError.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		_, err := fazpass.RemoveDevice(picID, meta, appID)
		assert.Error(t, err, "server returned status code 406")

	})

	t.Run("success", func(t *testing.T) {
		fazpass := Fazpass{
			BaseUrl:     mockServer.URL,
			MerchantKey: "mockMerchantKey",
			PrivateKey:  privateKey,
		}
		device, err := fazpass.RemoveDevice(picID, meta, appID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		// Memeriksa hasil pengujian
		expectedDevice := &Device{}
		expectedUnwrapMeta, _ := base64.StdEncoding.DecodeString("c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==")
		expectedDecrypted, _ := decryptWithPrivateKey(expectedUnwrapMeta, fazpass.PrivateKey)
		_ = json.Unmarshal(expectedDecrypted, &expectedDevice)
		if !isEqualDevice(device, expectedDevice) {
			t.Errorf("expected device %+v, got %+v", expectedDevice, device)
		}
	})
}

func TestRemoveAsyncDevice(t *testing.T) {
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/remove" {
			t.Errorf("expected '/remove' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody RemoveRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.FazpassId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.FazpassId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==", // base64 encoded string
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)
	fazpass := Fazpass{
		BaseUrl:     mockServer.URL,
		MerchantKey: "mockMerchantKey",
		PrivateKey:  privateKey,
	}

	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	deviceChan, errChan := fazpass.RemoveAsyncDevice(picID, meta, appID)

	select {
	case device := <-deviceChan:
		// Verifikasi hasil device
		expectedDevice := &Device{FazpassId: "fazpass_id"} // Ganti dengan nilai yang diharapkan
		if !isEqualDevice(device, expectedDevice) {
			t.Errorf("expected device %+v, got %+v", expectedDevice, device)
		}
	case err := <-errChan:
		// Verifikasi error
		t.Errorf("unexpected error: %v", err)
	case <-time.After(5 * time.Second):
		t.Error("timeout occurred")
	}
}

func TestErrorRemoveAsyncDevice(t *testing.T) {
	mockServer := createMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Menangani permintaan POST ke endpoint mock
		if r.Method != "POST" {
			t.Errorf("expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.Path != "/remove" {
			t.Errorf("expected '/remove' path, got '%s'", r.URL.Path)
		}
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer mockMerchantKey"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected Authorization header '%s', got '%s'", expectedAuthHeader, authHeader)
		}
		// Memeriksa request body
		var requestBody RemoveRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		expectedPicID := "mockPicID"
		expectedMeta := "mockMeta"
		expectedAppID := "mockAppID"
		if requestBody.FazpassId != expectedPicID {
			t.Errorf("expected PicID '%s', got '%s'", expectedPicID, requestBody.FazpassId)
		}
		if requestBody.Meta != expectedMeta {
			t.Errorf("expected Meta '%s', got '%s'", expectedMeta, requestBody.Meta)
		}
		if requestBody.MerchantAppId != expectedAppID {
			t.Errorf("expected MerchantAppID '%s', got '%s'", expectedAppID, requestBody.MerchantAppId)
		}

		// Mengirimkan respons palsu
		response := Response{
			Data: Data{
				Meta: "c0L1/z4JhlsUxoVeyiXzYulAFuokf01iGIR2w0QXMHSCjfxHgz8McwKr6ksMNSmUJciqE5wUUuiKerqOuTdluqqmLR9QUXBvLAZzLOxrRR4z1CyMiPFVuLeQZLv65Da+uthTmPSYjv5/jVJ/0r+0/REmF56xVW4pIxCpASmEO8AQXJ//dMw4zKxwvy1lQvxRCiNO+CAuDDrB16effJgWF+yE+eGVb05+v2PUZdgYXHV9pn3l5SoyZYjtL8+qrkzGj9SOISACnWQrRQrC7eSDlS5cK4IT5n9GiG07fg7lB8+xDjdGdtVSprwEILih+xBLLHOKlhmVKJtYG++IsrBmH1BI4RugDzP65IPh1ms3FBJ3gOrHMhW8wkX26WLBa28ujarb+CaZ3Cl1qcTivmPaZouhuUFFXcELw/KFe73Tt0RUQ52j5HuhXtjqRTj4XIlCi4Rfcn3hontmdRopwAlqGd6yvlKB/YUhe6r12Enh31xR+Am+5l1NJ6s/l7V8rSAfC/imhvkwyoxNPNIvJFt9+t1aPUjophWh+oaXtoNdYqFPb1bztQ/N9vjUG1J4cua235woEpFiBKf/6Ply8UQEFW1DNt0muQltuwwnk5yuyUHlsOsd/zA7/UKuOF3Gk8g9bAOkS3B7TSllSuVKvaP6DzrMWj0VCDxZnDOFQDsWW1IgduqDQxO2c5o6oAatBy9Uk0V+xtXt9N36OwoyDGVvm1+hmtzdEzBRhPZxNTG1Odq2tozpfcThLVHyiiJPI9wi/mjVxINW/TJy/R72cNO7H/wdgjoFpb29w2TFJoiN8NjjU5zrSjcTA7V3hIUkVbTGPcokdEVP7TxhutK+FTCFk8Jdw3U/PLhd5J7D0DlZYosRAOsKkyJh8DqMjK9CQTJAnP4VFvjZcZSOr3s6PvjNTQVCJjrK1cICe+b2ezaHoNQW2IJJaoPC0MvxWOFK3TwcFuh5BIN1javoo8PkN2656kixGS09p+bqsSCFyiFyNzkY7yFn7znR6S/+IBTNTTL4TSE9EKiJN18qyUm3JijCLf+TfO9hYOLUPg+ROI04EWQf7PpDNTsqiyxQa+gRGp3koCa5e0yRBxQOlRRsNJlZ/n60Hu9YRLBDM29m0NuWxkEBb3cZrxyPLvwveBtXoZH3kcq1FXMb4ChJF9WJHZTb/6zoK0RiC4XiidqZgzlF2yJTrGEkatBoSZYhOKzS8IP+bJmjS1jxMfmbxR4deericY7G+h9m+0Yyp919/cu2VukmCd/5JUzFssRYBtDWLdDHXrl48en2hX8BKq6CP4nnTIgFVrCQwA9WKJhBwnUWIFkIg2psBXAD/DTGCe3t2daEAwQ6eUt0VEBfOeCyqb3KDw==", // base64 encoded string
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	file, _ := os.ReadFile("./sample/private.key")
	privateKey, _ := bytesToPrivateKey(file)
	fazpass := Fazpass{
		BaseUrl:     mockServer.URL,
		MerchantKey: "mockMerchantKey",
		PrivateKey:  privateKey,
	}

	picID := "mockPicID"
	meta := "mockMeta"
	appID := "mockAppID"

	_, errChan := fazpass.RemoveAsyncDevice(picID, meta, appID)

	select {
	case err := <-errChan:
		// Verifikasi error
		assert.Error(t, err, "server returned status code 406")
	case <-time.After(5 * time.Second):
		t.Error("timeout occurred")
	}
}

func TestInitialize(t *testing.T) {
	t.Run("error parameter", func(t *testing.T) {
		_, err := Initialize("", "", "")
		assert.Error(t, err, "parameter cannot be empty")
	})
	t.Run("error file", func(t *testing.T) {
		_, err := Initialize("./key", "http://localhost.com", "merchant_KEY")
		assert.Error(t, err, "file not found")
	})
	t.Run("success", func(t *testing.T) {
		_, err := Initialize("./sample/private.key", "http://localhost.com", "merchant_KEY")
		assert.Equal(t, err, nil)
	})
}

// Fungsi untuk memeriksa kesamaan dua objek Device
func isEqualDevice(d1, d2 *Device) bool {

	return d1.FazpassId == d2.FazpassId
}
