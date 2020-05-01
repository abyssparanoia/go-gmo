package creditcard

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// SaveMemberRequest ... save member request
type SaveMemberRequest struct {
	MemberID   string `json:"MemberID" validate:"required,max=60"`
	MemberName string `json:"MemberName"`
}

// Validate ... validate
func (r *SaveMemberRequest) Validate() error {
	return validate.Struct(r)
}

// SaveMemberResponse ... save member response
type SaveMemberResponse struct {
	MemberID string `json:"MemberID"`
	ErrCode  string `json:"ErrCode"`
	ErrInfo  string `json:"ErrInfo"`
}

// SaveMember ... save member
func (cli *Client) SaveMember(
	req *SaveMemberRequest,
) (*SaveMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &SaveMemberResponse{}
	_, err := cli.do(saveMemberPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
