package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header http.Header
		want   string
	}{
		"WithAPIKey": {
			header: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey 1233344343")
				return h
			}(),
			want: "1233344343",
		},
		"WithoutAPIKey": {
			header: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey ")
				return h
			}(),
			want: "",
		},

		"MalformedAPIKey": {
			header: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "1")
				return h
			}(),
			want: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := GetAPIKey(tc.header)

			if !cmp.Equal(got, tc.want) {
				t.Errorf("GetAPIKey() = %v, want %v", got, tc.want)
			}
		})
	}
}
