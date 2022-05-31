package forms

type STKRequest struct {
	Msisdn string `json:"msisdn" binding:"required"`
	Amount string `json:"amount" binding:"required"`
}
