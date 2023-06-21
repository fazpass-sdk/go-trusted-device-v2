package dto

type CheckRequest struct {
	PicId         string `json:"pic_id"`
	Meta          string `json:"meta"`
	MerchantAppId string `json:"merchant_app_id"`
}

type EnrollRequest struct {
	PicId         string `json:"pic_id"`
	Meta          string `json:"meta"`
	MerchantAppId string `json:"merchant_app_id"`
}

type ValidateRequest struct {
	FazpassId     string `json:"fazpass_id"`
	Meta          string `json:"meta"`
	MerchantAppId string `json:"merchant_app_id"`
}

type RemoveRequest struct {
	FazpassId     string `json:"fazpass_id"`
	Meta          string `json:"meta"`
	MerchantAppId string `json:"merchant_app_id"`
}
