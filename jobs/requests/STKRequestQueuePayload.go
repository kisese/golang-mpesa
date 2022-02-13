package requests

type STKRequestPayload struct {
	Amount          string `json:"amount"`
	Msisdn          string `json:"msisdn"`
	Paybill         string `json:"paybill"`
	TrxId           string `json:"trx_id"`
	ReferenceNumber string `json:"reference_number"`
	CallbackUrl     string `json:"callback_url"`
}
