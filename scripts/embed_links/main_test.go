package main

import (
	"path/filepath"
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

func TestEmbedLinks(t *testing.T) {
	tests := []struct {
		desc      string
		path      string
		postLinks []string
		want      []string
		wantErr   bool
	}{
		{
			desc: "正常系: 読み取れる",
			path: "sample1.md",
			postLinks: []string{
				"* sample 1",
				"* sample 2",
			},
			want: []string{
				"# sushi",
				"",
				"* 123",
				"  * 456",
				"",
				"<!-- START_POSTS -->",
				"* sample 1",
				"* sample 2",
				"<!-- END_POSTS -->",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)

			path := filepath.Join("testdata", tt.path)
			got, err := embedLinks(path, tt.postLinks)
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
