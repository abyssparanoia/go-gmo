package payment

import (
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
)

func TestClient(t *testing.T) {

	type Req struct {
		CardSeq string `schema:"CardSeq"`
		ErrCode string `schema:"ErrCode"`
		ErrInfo string `schema:"ErrInfo"`
	}

	r := &Req{
		CardSeq: "cardSeq|cardSeq2",
		ErrCode: "errCode|errCode2",
		ErrInfo: "errInfo",
	}

	form := url.Values{}
	err := parser.Encoder.Encode(r, form)
	if err != nil {
		t.Logf(err.Error())
	}

	formStr := form.Encode()

	t.Logf(formStr)

	var res Req

	q, _ := url.ParseQuery(formStr)

	err = parser.Decoder.Decode(&res, q)

	multiRes := parser.ParseToMultiObject(res)

	t.Logf("%+v", multiRes)

	var result []*Req

	for _, res := range multiRes {
		var req Req
		t.Logf("%+v", res)
		err = parser.MapToStruct(res, &req)
		if err != nil {
			t.Logf(err.Error())
			continue
		}
		result = append(result, &req)
	}

	t.Logf("%+v", result[0])
	t.Logf("%+v", result[1])
}
