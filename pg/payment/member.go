package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// SaveMemberRequest ... save member request
type SaveMemberRequest struct {
	MemberID   string `schema:"MemberID" validate:"required,max=60"`
	MemberName string `schema:"MemberName,omitempty"`
}

// Validate ... validate
func (r *SaveMemberRequest) Validate() error {
	return validate.Struct(r)
}

// SaveMemberResponse ... save member response
type SaveMemberResponse struct {
	MemberID string `schema:"MemberID,omitempty"`
	ErrCode  string `schema:"ErrCode,omitempty"`
	ErrInfo  string `schema:"ErrInfo,omitempty"`
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

// UpdateMemberRequest ... update member request
type UpdateMemberRequest struct {
	MemberID   string `schema:"MemberID" validate:"required,max=60"`
	MemberName string `schema:"MemberName,omitempty"`
}

// Validate ... validate
func (r *UpdateMemberRequest) Validate() error {
	return validate.Struct(r)
}

// UpdateMemberResponse ... update member response
type UpdateMemberResponse struct {
	MemberID string `schema:"MemberID,omitempty"`
	ErrCode  string `schema:"ErrCode,omitempty"`
	ErrInfo  string `schema:"ErrInfo,omitempty"`
}

// UpdateMember ... update member
func (cli *Client) UpdateMember(
	req *UpdateMemberRequest,
) (*UpdateMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &UpdateMemberResponse{}
	_, err := cli.do(updateMemberPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteMemberRequest ... delete member request
type DeleteMemberRequest struct {
	MemberID string `schema:"MemberID" validate:"required,max=60"`
}

// Validate ... validate
func (r *DeleteMemberRequest) Validate() error {
	return validate.Struct(r)
}

// DeleteMemberResponse ... delete member response
type DeleteMemberResponse struct {
	MemberID string `schema:"MemberID,omitempty"`
	ErrCode  string `schema:"ErrCode,omitempty"`
	ErrInfo  string `schema:"ErrInfo,omitempty"`
}

// DeleteMember ... delete member
func (cli *Client) DeleteMember(
	req *DeleteMemberRequest,
) (*DeleteMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &DeleteMemberResponse{}
	_, err := cli.do(deleteMemberPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SearchMemberRequest ... search member request
type SearchMemberRequest struct {
	MemberID string `schema:"MemberID" validate:"required,max=60"`
}

// Validate ... validate
func (r *SearchMemberRequest) Validate() error {
	return validate.Struct(r)
}

// SearchMemberResponse ... search member response
type SearchMemberResponse struct {
	MemberID   string `schema:"MemberID,omitempty"`
	MemberName string `schema:"MemberName,omitempty"`
	DeleteFlag string `schema:"DeleteFlag,omitempty"`
	ErrCode    string `schema:"ErrCode,omitempty"`
	ErrInfo    string `schema:"ErrInfo,omitempty"`
}

// SearchMember ... search member
func (cli *Client) SearchMember(
	req *SearchMemberRequest,
) (*SearchMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &SearchMemberResponse{}
	_, err := cli.do(searchMemberPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
