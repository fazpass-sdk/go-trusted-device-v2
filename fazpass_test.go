package gotdv2

import (
	"encoding/base64"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const stagingMeta = "u6DetnC4wkfxgQ77MatnYMbraMZ5/sSw3DNVH3EZdRtcmvXM1qPt4K8KPI0F56MCXDqqK6RZx8ky5n1vSMK2wKRUyCtYFCqC+FDTIYo5caVmSYHYd/+Nuk2y462LMZvE40oNlahFkbl6BxXflMBgnzb13e1ZHMihJwLQbjXd/7qpVzhZd+ab96fJ2UPAC9LBQVg6hCX1ACQ2B2u340Oz8usBoPhr51YHB2bNJ4JvJ2v25gYTS0TNff8A32hYT15r3rh/l02CfC4WXUR2zxMxQh9rlTFg1szdW7KeW3jiixULSw2D959Pq+HgisZSZ7TDVjJHEsqERY8mgyDfS9R3eAJCp7kWbFaol0cGyUg2Vq6uAX+qt2ejQftlqEmJGbii/SkTHb/QbyJOdqnpElG+YhZDgzKZNal52+85yUypw4G9oodno+4wExFFGi4M+SbdUMKQSvsddZzZ9O2Yg0CKjPFU1jyoTYFdo7gxiMAe1aV/ggE187dVuXNNqceC8z1xqmByv1Gbs5Lwuit4J9sFVFNmhcYjnSgg2+oZXGKGGhIhXL+2SPWToH16R01JmxRT3v4fz0FOD+WpTrMLgm6GW6plgc01TGxy5NH6KYBm4ZlmYwFLBZBRrrHmLCEtugU0bfXvDyyMF/q9CcXMCnye1QmMo9Z4zyXIbn9gBAa2JW3VmvOgk18x9sm2Vwr5Q4RKZYmfnAzUDa6BNf4iD0Jh15o+xFiTJT/RSHg4oS4m2g/XOdjzsXu+jEF8FUbZGv2hR86QvEPdmC2hL6LKKfoG53UGu+pkRG3bPvXSRN3+qK53WMpd53V+NFTHig9PtWU1NVxe9Z7nZddDe4c/wxtl4lysEPRbLjtRB0zIkkg05OxOBBLnRfIK8WCHSKXNjUAMS5uPmi03PVMcccZR47E7DQN44LF/sRU34IXrqcWyzI2t6DQh3BTt1KfhcuMKwgpmPtTL2DFj1/YHjIz+c19KKznUm3NtFpxYBveQsqa6myM+QIyly7g4rWRb+vEqc3QSgsT99pAprKcttDvaXvjh0BqleEL6muYQ1ivjogUDi+da/5QosoePZ5cJwO50WokMRkVQWAtXwJMh8BUfK9BccrFghWfPI9J8Uehb87epo2T+GnCDdlpk2PJbUl3DygsCgNKEno7kVwprnT7TFToLfGIrA/J04FSHRZt+9fpfq85vhWEmNQsjE9bJirJZGs+2gCp5WEOLX1h2B2AjHNrVqq/WhIVjm9O3RGLFJTQtfptdTsUbLMss6OnNUDePHAk4Jy95VdGOtAXze1eqqYGexn+hfSPxi9JRwYPSM2aUEOY8yXXShJ+fP3DwMsQdroJQ4QCjCUuqNy1EmYg2b/ibFg=="

// const newMeta = "NlvZpR44O4TOMCYTRsBLHKMVZLucupg2NJVI4Kc2/4EP/Ml9ckHLZXi/XAmIOdgChU3rJhEf8JB5jAXA5I6j8+ls+jIpF13sU+IeP2PezBX5zYYMhWI5NVIdYSefGFFBOKEceLNqgMybpTneRm98o3mdbJrKg4cAJJDAhbeNLRBHD1XZFj8SCMaupfuxY/5I09oia4ms3/eikIwLO6w1lUZT0Wo9JHr2bZM14JhWVm3qHMNVMp8dceS+SMvRa7dnLH1Ca7dD+Zxq9gfIWO73ULAWHMHtWLc3d+6Kfip7eYmHborpFBG3NU4ZBC3NC7L5FPlndsU8m2ODybzkNWAIPbBRz5q3gKVGorJUs4fIK71myEcoTOwHPvHpcpGeKXSKQgyb9uu3ZOXMKZVDXGDMXUH5UqcRz/jWz937Erabo08L8AFtzNkG8AQwaLzthApNtlPOfpajYfgKGydqv/x4dBHvZ3YmPhQl1uAuaRXMUi/ceHzqn8r0X2iT20I8r3ilauujQgFRK+RVY86n/xi/fVpDem9KvpFXSoaJBJM3WxqrFIc2ZvRGItR7AraEoCE404nPgC9DltOCq+50X5cNYIyeqZJO7UkiW0NjmX1DySPmvcTev6M7aW8QCpoXLDehaEGyAdaHiMK+fh/r+p5ELnth3zztpAanCFBo3r6OoPgA14QA5myUqVW1ppcH21Ukc3hEZBWFQMMeQPpjIjG3PVSHshpcE7gQiiYPI3DuHodmf5JvXDsB0AbaMbeKpiNTsBvzp9ZWiJ+ucBfjS0nmOrtVbkw1lLc6gvkm6vl71riS5yPZ7yK1JbfNHMvMZfDhOF8ZDraSZdg8uVj1BKcF2fSqayCmnnLbWE54jIdS6GlNP9o79qGOH/QfVLLYfZbS30Q9eaJr4WqZVewFe/0LoSKEQj4r20+C37dv1T21AuzQ19kMu77GM+GC5GeXYJnykSTKA5qRrrqJ24JsDFe/MucxVFomdinuqqTjmMq/yUVxxOWf7H/Z7Hb1jB913UjEvEwAklRRv0WPTsmkcS78nLHVVq+jKBNNJS3Ek/IwHAkb7+Mz57g5fLWG7kPA6fjLM9sCnQWn+RtbL4+5tEYpeQYdYIMv2S6CRgHndI/oq/l8hzGXaUd2Eo/zUMXPVOFbfKDgcx4XZWBiTBAyU7BqSW59HQ2bRElfOvkNpnl7I+xI6OiXIJ1yuYnF1HqqP3NTXCRrO3vZKyOw9QpY/+Y+5wQr81sXMeccBNnxNe0M4dLMEEIRXwpBm3Cbbbf117auDFaViaWDwUsO13YfO7KPgEOx+MUhWPT3nh8G2KiwadquhgexTAVgxZJfdrPQAV+e80yDyIX0P+DPh+RB48+xFw=="
const responseMeta = "KsQvJaJSGYH56IeFzjw+6lZOGamb2i9uw1lUwQtWHNgLlONzp3AFtH2AAoChTA1PtkX6iJMLRSHJ2YV2/DgqyBqydiLdOT0y3Bk/oT509faWytilfm6A6s/86vaJz2zZTaziipAzsJ9eC9PwxXBJt1gOLdZ+WmVQWUeUBdC+mvbNm3lKrY3ukbIcRUbXBvdQT2c4/dQZmiS665TSfAcWjkO1zPCD5fmOAZcoXJxVhmEjkIH3Y9EXze4i2MUooXbEU6vjOKeBQSHSTGpeWMURPbKeuIYhRbWitQzNHwnP9e2WtiEggcKIteBeXJwxt3uKBZ2iVrYx8MjZe7g2LsXj70xrILqIO76gWMyMJoI/5Kp4KfsDr5qdrQvV6H0Idy/HA3/Js8fXuc6km2WOc28belON3p5CoWepNVcxCfyd7M5VYdhWpvfijKTPHnv522dgCb7Z24/cJ0BOeQH3RXCWXx4GaLFEL4sXkmtjALOjB7w+kt3sHoC0EaGWpPLMgCAdrKHd1fuTZzML+kyjMdaK2V+wKXvp7TEyTAIAHj+5Owte5oY8Nn6875py+iGQLXU4pm5pMMOnExVFebqYd8IznIPawIbwo/GphXrifxtf327/cUWPXLROUJjCILychKWYgnYggbjSjg/F8H+Jgp69cfZZ0n69/fm3mUpX1VghVX9Z+LJvD2h/JihH84lMgnYDPD4B4tMPZ9b0K3AXMVAihDG53EXrZdHYX6RQP/RYavfMBMpLKJKB4cCVngnnjxvUA3Kget+YC6b+0OJpDUNSDBc/ChkAMxWVcZE6eZmBbRgBZudmQk9XMCHn0hwcg8yYxInbddJKRJvLvnd+Sq4ePRY9nE3yyDAWhNOg4HUlSbw7oWzCR2CozimrDywow1W652UYoL3DEuZ7f7zChb+ZGNvOnI0mJs6j4YloJbGTX699OxFcEix79Rp5AnrdMq6gLwCxjWQ0st8U2vjq/ncApklqSHpMHVL26RkLBGwhskCtEGr/AspwMnC8oaXlMNxaWXqnvBHOc4AM3XpfVRXCP8Wm3HWpZANYkREDeItmA28DbaR50StDyPwbRJPaKYrYr/JiNfU0+LD+5l1UlmDun+LeGmDCtr1rFtw13LYVxCYz6nWzWAlHwqN49fbFv14ykTfOf/+Ji/XBZ4mN9XRMgaW84LMta6OeGUEZ5k+jFqxvRJJHUorQ7JAfdqPmGdUN3duu9ZUunyqcq8ugNihDp8WvdZb4gGKCFKrwfyDudPu+R2VItPvtBBWHHhmx7Ks0NvvBaN5EwHTFUKS21Q2CnGwfFln9KgHr3rO12G4XBwQ3rHaqMCvDjqs6wRZDcmQ6bZJX3gLFrPnEvUGlERVm1g=="

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
		d, err := f.Extract(responseMeta)
		valueOf := reflect.ValueOf(d)
		typeOf := valueOf.Type()

		for i := 0; i < valueOf.NumField(); i++ {
			field := valueOf.Field(i)
			log.Printf("Field Name: %s, Field Value: %v\n", typeOf.Field(i).Name, field.Interface())
		}
		if err != nil {
			log.Println(err)
		}
		assert.Equal(t, "103.28.116.40", d.ClientIp)
	})
}

func encrypt() string {
	keyP, _ := os.ReadFile("./sample/public_1.pub")
	pubKey := bytesToPublicKey(keyP)
	encResponseData := encryptWithPublicKey([]byte("koala"), pubKey)
	return base64.StdEncoding.EncodeToString(encResponseData)
}
