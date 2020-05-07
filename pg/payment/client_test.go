package payment

import (
	"net/url"
	"testing"

	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
)

func TestClient(t *testing.T) {

	type Req struct {
		CardSeq string `schema:"CardSeq"`
	}

	r := &Req{
		CardSeq: "cardSeq|cardSeq2",
	}

	form := url.Values{}
	err := parser.Encoder.Encode(r, form)
	if err != nil {
		t.Logf(err.Error())
	}

	formStr := form.Encode()

	var res Req

	q, _ := url.ParseQuery(formStr)

	err = parser.Decoder.Decode(&res, q)

	multiRes := parser.ParseToMultiObject(res)

	var result []*Req

	for _, res := range multiRes {
		var req Req
		err = parser.MapToStruct(res, &req)
		if err != nil {
			continue
		}
		result = append(result, &req)
	}

}
