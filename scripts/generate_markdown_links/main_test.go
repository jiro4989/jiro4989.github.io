package main

import (
	"bytes"
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

func TestReadAttrLine(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		attr    string
		want    string
		wantErr bool
	}{
		{
			desc:    "正常系: 最初の layout を読み取る",
			path:    "sample1.md",
			attr:    "layout:",
			want:    "default",
			wantErr: false,
		},
		{
			desc:    "異常系: 該当する attr が見つからない",
			path:    "sample1.md",
			attr:    "pohe:",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)

			path := filepath.Join("testdata", tt.path)
			got, err := readAttrLine(path, tt.attr)
			if tt.wantErr {
				a.Error(err)
				a.Empty(got)
				return
			}
			a.Equal(tt.want, got)
			a.NoError(err)
		})
	}
}

func TestReadFileAttr(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		want    *fileAttr
		wantErr bool
	}{
		{
			desc: "正常系: 属性をすべて読み取れる",
			path: "sample1.md",
			want: &fileAttr{
				title:    "色彩検定 2 級に合格した",
				year:     "2021",
				month:    "12",
				day:      "14",
				hmd:      "09:00:00",
				tz:       "+0900",
				sortKey:  "2021-12-14 09:00:00 +0900",
				category: "tech",
				fileName: "sample1",
			},
			wantErr: false,
		},
		{
			desc:    "異常系: title 属性なし",
			path:    "err_no_title.md",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "異常系: date 属性なし",
			path:    "err_no_date.md",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "異常系: categories 属性なし",
			path:    "err_no_categories.md",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)

			path := filepath.Join("testdata", tt.path)
			got, err := readFileAttr(path)
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

func TestToMarkdown(t *testing.T) {
	tests := []struct {
		desc string
		fa   *fileAttr
		want string
	}{
		{
			desc: "正常系: 属性をすべて読み取れる",
			fa: &fileAttr{
				title:    "色彩検定 2 級に合格した",
				year:     "2021",
				month:    "12",
				day:      "14",
				hmd:      "09:00:00",
				tz:       "+0900",
				sortKey:  "2021-12-14 09:00:00 +0900",
				category: "tech",
				fileName: "sample1",
			},
			want: "* 2021-12-14 技術 [色彩検定 2 級に合格した](/tech/2021/12/14/sample1.html)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)

			got := tt.fa.ToMarkdown()
			a.Equal(tt.want, got)
		})
	}
}

func TestGenerateLinks(t *testing.T) {
	tests := []struct {
		desc    string
		i       inventory
		want    string
		wantErr bool
	}{
		{
			desc: "正常系: 属性をすべて読み取れる",
			i: inventory{
				LoopKeys: []string{
					"2023",
					"2022",
					"2021",
				},
				FileAttrs: map[string][]*fileAttr{
					"2023": {
						{
							title:    "Test 1 2023",
							year:     "2023",
							month:    "06",
							day:      "01",
							hmd:      "09:00:00",
							tz:       "+0900",
							sortKey:  "2023-06-01 09:00:00 +0900",
							category: "tech",
							fileName: "test-1-2023",
						},
						{
							title:    "Test 2 2023",
							year:     "2023",
							month:    "05",
							day:      "01",
							hmd:      "09:00:00",
							tz:       "+0900",
							sortKey:  "2023-05-01 09:00:00 +0900",
							category: "tech",
							fileName: "test-2-2023",
						},
						{
							title:    "Test 3 2023",
							year:     "2023",
							month:    "04",
							day:      "01",
							hmd:      "09:00:00",
							tz:       "+0900",
							sortKey:  "2023-04-01 09:00:00 +0900",
							category: "tech",
							fileName: "test-3-2023",
						},
					},
					"2022": {
						{
							title:    "Test 1",
							year:     "2022",
							month:    "03",
							day:      "01",
							hmd:      "09:00:00",
							tz:       "+0900",
							sortKey:  "2022-03-01 09:00:00 +0900",
							category: "tech",
							fileName: "test-1",
						},
						{
							title:    "Test 2",
							year:     "2022",
							month:    "02",
							day:      "01",
							hmd:      "09:00:00",
							tz:       "+0900",
							sortKey:  "2022-02-01 09:00:00 +0900",
							category: "tech",
							fileName: "test-2",
						},
					},
					"2021": {
						{
							title:    "色彩検定 2 級に合格した",
							year:     "2021",
							month:    "12",
							day:      "14",
							hmd:      "09:00:00",
							tz:       "+0900",
							sortKey:  "2021-12-14 09:00:00 +0900",
							category: "tech",
							fileName: "sample1",
						},
					},
				},
			},
			want: `### 2023 年

* 2023-06-01 技術 [Test 1 2023](/tech/2023/06/01/test-1-2023.html)
* 2023-05-01 技術 [Test 2 2023](/tech/2023/05/01/test-2-2023.html)
* 2023-04-01 技術 [Test 3 2023](/tech/2023/04/01/test-3-2023.html)

### 2022 年

* 2022-03-01 技術 [Test 1](/tech/2022/03/01/test-1.html)
* 2022-02-01 技術 [Test 2](/tech/2022/02/01/test-2.html)

### 2021 年

* 2021-12-14 技術 [色彩検定 2 級に合格した](/tech/2021/12/14/sample1.html)

`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)

			w := &bytes.Buffer{}
			err := generateLinks(tt.i, w)
			if tt.wantErr {
				a.Error(err)
				return
			}
			got := w.String()
			a.Equal(tt.want, got)
		})
	}
}
