package gotdv2

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

const stagingMeta = "stS3wBIZiHef+k8g48JUOaClN5N5VsgNlMBkMHMt4XLsatqk8otidq6B6XKbynunzpBVfbFWSdRNyq0Ngf6sPP5EKHIK1sSbrLyRQxrJcyesoAuT7/G2FdCKSCXJbqqSwizkN1vqks1LK4N/UT1JKZDnZtBRUoKBU80YsSrWsLoUx23SlI5+jq3Xakd58NWbC8PUe2WfpZ0ipBFkLnOUMhZSvM5PjpaXNs9K8sNDcSUru/J7JTkfJs+5Tye2lytKCVWDG3xoZw2NDwgrrHJjj0ET2oHGSk0iWAwhEFAK0y5vBiEnaNdCkBxdrVRC49t1T63fJ8Nf5n6yHlFyF2YXJ5CQlAN7pq8uMB4gnMvaDBVBQmKGlvadNmF3q6vs6t9uGtVrJgTVsMryMwYqABu1H7/BSTkwiMTZuyHto3K5IlkQBkCa2J8PDuQ22F29/UIPza73C0v5jJWnQA4r61K/lp6MCn1Vv8MGfQ4dtP4Qkc8IoCu5X/NGTU4bI0K0/9nWzpVVEgPUs3ZuHIzxXEyve4pRQ0AXSS9FtG8p93alMgXDv6Ig7OkUtTlgHLQbLOhy+sXFrD6SZYRAsGgZP5kruy+znL0GYXgd4v0zbZTvoWTolMYt41Y19vC7MxpFeOuvKvx90p19XAVvEQVzmLLwUmrKZW/16c1o4vS/152DjyLS/sLOAVs8tkdfeb6y48lmBJNB8JpEoLtg2Z1xNwMyz2JbxtMLZfAicOGn+ifAqI+S+FmOrPO0o96A4LxuXou7rKG1ljS34cv3h0ANloLa/36U/Z/XJU3kl5nGuReyb2d9qhKzNepdMx0mBFD88jNno/OBIylVHSQhrDINZ7g4aMosZg6LG+N9db9HSV35HiNIep6uLdEtVRa4Wh6qXUBVFlRyaIsUrDAlWS9/wf9Y0thekd3wn5SpADt+PVDw2QRtMaTSKfobetvUHLKgwX8N+4ErS2BkfBUHKIFYRtH6qz66d1t9ZMPbgE176o4ffEoJHtkSF/OoiqiqQCDNwGwGLRZUTyhG7qzBpV4NduO/2pa3vH+e1T6TKjZGju9YqAQir6DdPW3revfHqtkKucZbp1bgbb0r5xS7lOVQ1yrT2V7+vQJ7Ot6ZdD/UcFc5f5J4uLhLd06DcSJpadMuALrJRLqGzOTtouV92iMLj4UHB3AzpvpnX9LHayj8UtGV7PqjOVw4WndbTwJI9wFrjawh0PGzG/jC+05Ope6uq0rlPnDIV4R1Lk7Zf0xBS79p5LSUiJCkcC7AXKYoIZPcXBAe8l1rmLvmGHYpc8RABpGBN3o5R5vyIoJYEbFVsg9wYN8F06mGGLi6caMiDqICbyGn5ZZMKDzja+1ib+HmZQkwVg=="

func TestInitialize(t *testing.T) {

	t.Run("parameter empty", func(t *testing.T) {
		_, err := Initialize("")
		assert.Error(t, err, "parameter cannot be empty")
	})

	t.Run("error file", func(t *testing.T) {
		_, err := Initialize("./key")
		assert.Error(t, err, "file not found")
	})
	t.Run("success", func(t *testing.T) {
		_, err := Initialize("./sample/private.key")
		assert.Equal(t, err, nil)
	})
}

func TestExtract(t *testing.T) {
	t.Run("decrypt failed", func(t *testing.T) {
		f, _ := Initialize("./sample/private.key")
		_, err := f.Extract(stagingMeta)
		assert.Equal(t, err.Error(), "invalid meta or key")
	})

	t.Run("error unmarshal", func(t *testing.T) {
		f, err := Initialize("./sample/private_1.key")
		meta := encrypt()
		_, err = f.Extract(meta)
		if err != nil {
			log.Println(err)
		}
		assert.Equal(t, err.Error(), "invalid character 'k' looking for beginning of value")
	})

	t.Run("decrypt success", func(t *testing.T) {
		f, err := Initialize("./sample/app_private_key.key")
		if err != nil {
			log.Println(err)
		}
		d, err := f.Extract(stagingMeta)
		if err != nil {
			log.Println(err)
		}
		assert.Equal(t, d.ClientIp, "127.0.1")
	})
}

func encrypt() string {
	keyP, _ := os.ReadFile("./sample/public_1.pub")
	pubKey := bytesToPublicKey(keyP)
	encResponseData := encryptWithPublicKey([]byte("koala"), pubKey)
	return base64.StdEncoding.EncodeToString(encResponseData)
}
