package creditcard

type EntryTranRequest struct {
	OrderID string `json:"OrderID" validate:"required,max=27"`
	JobCD   string `json:"JobCd" validate:"required"`
}
