package payment

import "github.com/abyssparanoia/go-gmo/internal/pkg/validate"

// SaveCardRequest ... save card request
type SaveCardRequest struct {
	MemberID     string `json:"MemberID" validate:"required,max=60"`
	SeqMode      string `json:"SeqMode"`
	CardSeq      int    `json:"CardSeq"`
	DefaultFlag  string `json:"DefaultFlag"`
	CardName     string `json:"CardName"`
	CardNo       string `json:"CardNo" validate:"required,len=16"`
	CardPass     string `json:"CardPass"`
	Expire       string `json:"Expire" validate:"required,len=4"`
	HolderName   string `json:"HolderName"`
	Token        string `json:"Token"`
	UpdateType   string `json:"UpdateType"`
	SecurityCode string `json:"SecurityCode"`
}

// Validate ... validate
func (r *SaveCardRequest) Validate() error {
	return validate.Struct(r)
}

// SaveCardResponse ... save card response
type SaveCardResponse struct {
	CardSeq                int    `json:"CardSeq"`
	CardNo                 string `json:"CardNo"`
	Forward                string `json:"Forward"`
	ErrCode                string `json:"ErrCode"`
	ErrInfo                string `json:"ErrInfo"`
	Brand                  string `json:"Brand"`
	DomesticFlag           string `json:"DomesticFlag"`
	IssuerCode             string `json:"IssuerCode"`
	DebitPrepaidFlag       string `json:"DebitPrepaidFlag"`
	DebitPrepaidIssuerName string `json:"DebitPrepaidIssuerName"`
	ForwardFinal           string `json:"ForwardFinal"`
}

// SaveCard ... save card
func (cli *Client) SaveCard(
	req *SaveCardRequest,
) (*SaveCardResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &SaveCardResponse{}
	_, err := cli.do(saveCardPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
