package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStdin(t *testing.T) {
	tests := []struct {
		desc    string
		text    string
		want    []string
		wantErr bool
	}{
		{
			desc:    "正常系: 読み取れる",
			text:    "sushi\nnuka",
			want:    []string{"sushi", "nuka"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)

			r := strings.NewReader(tt.text)
			got, err := readStdin(r)
			if tt.wantErr {
				a.Error(err)
				a.Nil(got)
				return
			}
			a.Equal(tt.want, got)
			a.NoError(err)
		})
	}
}
