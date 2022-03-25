package jobs

type STKRequestPayload struct {
	Amount      string `json:"amount"`
	Msisdn      string `json:"msisdn"`
	Paybill     string `json:"paybill"`
	CallbackUrl string `json:"callback_url"`
}
