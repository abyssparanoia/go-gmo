package payment

// AccountTransferRecord ... account transfer csv record
type AccountTransferRecord struct {
	SiteID       string `csv:"サイトID" validate:"required"`
	ShopID       string `csv:"ショップID" validate:"required"`
	MemberID     string `csv:"会員ID" validate:"required"`
	Amount       int    `csv:"利用金額" validate:"required"`
	OtherAmount  int    `csv:"税送料"`
	PassbookText string `csv:"通帳記載内容" validate:"required"`
	FreeText     string `csv:"自由項目"`
}
