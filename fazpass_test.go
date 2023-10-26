package gotdv2

import (
	"encoding/base64"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const stagingMeta = "stS3wBIZiHef+k8g48JUOaClN5N5VsgNlMBkMHMt4XLsatqk8otidq6B6XKbynunzpBVfbFWSdRNyq0Ngf6sPP5EKHIK1sSbrLyRQxrJcyesoAuT7/G2FdCKSCXJbqqSwizkN1vqks1LK4N/UT1JKZDnZtBRUoKBU80YsSrWsLoUx23SlI5+jq3Xakd58NWbC8PUe2WfpZ0ipBFkLnOUMhZSvM5PjpaXNs9K8sNDcSUru/J7JTkfJs+5Tye2lytKCVWDG3xoZw2NDwgrrHJjj0ET2oHGSk0iWAwhEFAK0y5vBiEnaNdCkBxdrVRC49t1T63fJ8Nf5n6yHlFyF2YXJ5CQlAN7pq8uMB4gnMvaDBVBQmKGlvadNmF3q6vs6t9uGtVrJgTVsMryMwYqABu1H7/BSTkwiMTZuyHto3K5IlkQBkCa2J8PDuQ22F29/UIPza73C0v5jJWnQA4r61K/lp6MCn1Vv8MGfQ4dtP4Qkc8IoCu5X/NGTU4bI0K0/9nWzpVVEgPUs3ZuHIzxXEyve4pRQ0AXSS9FtG8p93alMgXDv6Ig7OkUtTlgHLQbLOhy+sXFrD6SZYRAsGgZP5kruy+znL0GYXgd4v0zbZTvoWTolMYt41Y19vC7MxpFeOuvKvx90p19XAVvEQVzmLLwUmrKZW/16c1o4vS/152DjyLS/sLOAVs8tkdfeb6y48lmBJNB8JpEoLtg2Z1xNwMyz2JbxtMLZfAicOGn+ifAqI+S+FmOrPO0o96A4LxuXou7rKG1ljS34cv3h0ANloLa/36U/Z/XJU3kl5nGuReyb2d9qhKzNepdMx0mBFD88jNno/OBIylVHSQhrDINZ7g4aMosZg6LG+N9db9HSV35HiNIep6uLdEtVRa4Wh6qXUBVFlRyaIsUrDAlWS9/wf9Y0thekd3wn5SpADt+PVDw2QRtMaTSKfobetvUHLKgwX8N+4ErS2BkfBUHKIFYRtH6qz66d1t9ZMPbgE176o4ffEoJHtkSF/OoiqiqQCDNwGwGLRZUTyhG7qzBpV4NduO/2pa3vH+e1T6TKjZGju9YqAQir6DdPW3revfHqtkKucZbp1bgbb0r5xS7lOVQ1yrT2V7+vQJ7Ot6ZdD/UcFc5f5J4uLhLd06DcSJpadMuALrJRLqGzOTtouV92iMLj4UHB3AzpvpnX9LHayj8UtGV7PqjOVw4WndbTwJI9wFrjawh0PGzG/jC+05Ope6uq0rlPnDIV4R1Lk7Zf0xBS79p5LSUiJCkcC7AXKYoIZPcXBAe8l1rmLvmGHYpc8RABpGBN3o5R5vyIoJYEbFVsg9wYN8F06mGGLi6caMiDqICbyGn5ZZMKDzja+1ib+HmZQkwVg=="
const newMeta = "NlvZpR44O4TOMCYTRsBLHKMVZLucupg2NJVI4Kc2/4EP/Ml9ckHLZXi/XAmIOdgChU3rJhEf8JB5jAXA5I6j8+ls+jIpF13sU+IeP2PezBX5zYYMhWI5NVIdYSefGFFBOKEceLNqgMybpTneRm98o3mdbJrKg4cAJJDAhbeNLRBHD1XZFj8SCMaupfuxY/5I09oia4ms3/eikIwLO6w1lUZT0Wo9JHr2bZM14JhWVm3qHMNVMp8dceS+SMvRa7dnLH1Ca7dD+Zxq9gfIWO73ULAWHMHtWLc3d+6Kfip7eYmHborpFBG3NU4ZBC3NC7L5FPlndsU8m2ODybzkNWAIPbBRz5q3gKVGorJUs4fIK71myEcoTOwHPvHpcpGeKXSKQgyb9uu3ZOXMKZVDXGDMXUH5UqcRz/jWz937Erabo08L8AFtzNkG8AQwaLzthApNtlPOfpajYfgKGydqv/x4dBHvZ3YmPhQl1uAuaRXMUi/ceHzqn8r0X2iT20I8r3ilauujQgFRK+RVY86n/xi/fVpDem9KvpFXSoaJBJM3WxqrFIc2ZvRGItR7AraEoCE404nPgC9DltOCq+50X5cNYIyeqZJO7UkiW0NjmX1DySPmvcTev6M7aW8QCpoXLDehaEGyAdaHiMK+fh/r+p5ELnth3zztpAanCFBo3r6OoPgA14QA5myUqVW1ppcH21Ukc3hEZBWFQMMeQPpjIjG3PVSHshpcE7gQiiYPI3DuHodmf5JvXDsB0AbaMbeKpiNTsBvzp9ZWiJ+ucBfjS0nmOrtVbkw1lLc6gvkm6vl71riS5yPZ7yK1JbfNHMvMZfDhOF8ZDraSZdg8uVj1BKcF2fSqayCmnnLbWE54jIdS6GlNP9o79qGOH/QfVLLYfZbS30Q9eaJr4WqZVewFe/0LoSKEQj4r20+C37dv1T21AuzQ19kMu77GM+GC5GeXYJnykSTKA5qRrrqJ24JsDFe/MucxVFomdinuqqTjmMq/yUVxxOWf7H/Z7Hb1jB913UjEvEwAklRRv0WPTsmkcS78nLHVVq+jKBNNJS3Ek/IwHAkb7+Mz57g5fLWG7kPA6fjLM9sCnQWn+RtbL4+5tEYpeQYdYIMv2S6CRgHndI/oq/l8hzGXaUd2Eo/zUMXPVOFbfKDgcx4XZWBiTBAyU7BqSW59HQ2bRElfOvkNpnl7I+xI6OiXIJ1yuYnF1HqqP3NTXCRrO3vZKyOw9QpY/+Y+5wQr81sXMeccBNnxNe0M4dLMEEIRXwpBm3Cbbbf117auDFaViaWDwUsO13YfO7KPgEOx+MUhWPT3nh8G2KiwadquhgexTAVgxZJfdrPQAV+e80yDyIX0P+DPh+RB48+xFw=="

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
		f, err := Initialize("./sample/private_2.key")
		if err != nil {
			log.Println(err)
		}
		d, err := f.Extract(newMeta)
		valueOf := reflect.ValueOf(d)
		typeOf := valueOf.Type()

		for i := 0; i < valueOf.NumField(); i++ {
			field := valueOf.Field(i)
			log.Printf("Field Name: %s, Field Value: %v\n", typeOf.Field(i).Name, field.Interface())
		}
		if err != nil {
			log.Println(err)
		}
		assert.Equal(t, d.ClientIp, "114.10.25.88")
	})
}

func encrypt() string {
	keyP, _ := os.ReadFile("./sample/public_1.pub")
	pubKey := bytesToPublicKey(keyP)
	encResponseData := encryptWithPublicKey([]byte("koala"), pubKey)
	return base64.StdEncoding.EncodeToString(encResponseData)
}
