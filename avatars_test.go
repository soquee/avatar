package avatar_test

import (
	"net/http/httptest"
	"testing"

	"code.soquee.net/avatar"
)

func TestAvatarContentType(t *testing.T) {
	h := avatar.Handler()

	req := httptest.NewRequest("GET", "/avatars", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	const expect = "image/png"
	ct := w.Header().Get("Content-Type")
	if ct != expect {
		t.Errorf("Unexpected Content-Type: want=%q, got=%q", expect, ct)
	}
}
