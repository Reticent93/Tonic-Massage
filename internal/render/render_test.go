package render

import (
	"github.com/Reticent93/Tonic-Massage/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultdata(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	res := AddDefaultdata(&td, r)

	if res.Flash != "123" {
		t.Error("flash value 123 is not in session")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))

	r = r.WithContext(ctx)
	return r, err
}
