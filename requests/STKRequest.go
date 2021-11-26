package requests

type STKRequest struct {
	Type        string `json:"type"`
	Msisdn      string `json:"msisdn"`
	Amount      string `json:"amount"`
	ReferenceId string `json:"reference_id"`
}
