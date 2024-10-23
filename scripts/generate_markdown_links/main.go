package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

var categoriesMap = map[string]string{
	"tech":   "技術",
	"movie":  "映画",
	"game":   "ゲーム",
	"daily":  "雑記",
	"illust": "イラスト",
}

type fileAttr struct {
	title    string
	year     string
	month    string
	day      string
	hmd      string
	tz       string
	sortKey  string
	category string
	fileName string
}

type inventory struct {
	LoopKeys  []string
	FileAttrs map[string][]*fileAttr
}

func main() {
	if err := Main(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}

func Main(r io.Reader, w io.Writer) error {
	paths, err := readStdin(r)
	if err != nil {
		return err
	}

	// ファイルを読み込んで属性を取得する
	fas := make([]*fileAttr, 0)
	for _, path := range paths {
		fa, err := readFileAttr(path)
		if err != nil {
			return err
		}
		fas = append(fas, fa)
	}

	sort.Slice(fas, func(i, j int) bool {
		return fas[i].sortKey > fas[j].sortKey
	})

	// 年をキーにファイル属性のスライスを登録
	yearFa := make(map[string][]*fileAttr)
	for _, fa := range fas {
		year := fa.year
		yearFa[year] = append(yearFa[year], fa)
	}

	// map のキーは順序不定なので、年で降順ソートするためのスライスを作成
	loopKeys := make([]string, 0)
	for year := range yearFa {
		loopKeys = append(loopKeys, year)
	}
	sort.Slice(loopKeys, func(i, j int) bool {
		return loopKeys[i] > loopKeys[j]
	})

	// text/template で見出しと箇条書きを生成
	inv := inventory{
		LoopKeys:  loopKeys,
		FileAttrs: yearFa,
	}
	if err := generateLinks(inv, w); err != nil {
		return err
	}

	return nil
}

// readStdin は標準入力を文字列の配列として返す。
func readStdin(r io.Reader) ([]string, error) {
	ret := make([]string, 0)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		ret = append(ret, line)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}

// readAttrLine は Jekyll の Front Matter から attr にマッチにマッチする行を取得して、属性名を削除して返却する。
//
// Ex:
//
//	categories: movie
//
// 上記のコードからは movie を返却する。
func readAttrLine(path, attr string) (string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, attr) {
			continue
		}
		return strings.TrimSpace(line[len(attr):]), nil
	}
	if err := sc.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("%s doesn't exist: file = %s", attr, path)
}

// readFileAttr はファイルの先頭にかかれている Jekyll Front Matter の属性を読み取って返却する。
func readFileAttr(path string) (*fileAttr, error) {
	titleLine, err := readAttrLine(path, "title:")
	if err != nil {
		return nil, err
	}
	title := strings.Trim(titleLine, `"`)

	dateLine, err := readAttrLine(path, "date:")
	if err != nil {
		return nil, err
	}
	dateParts := strings.Split(dateLine, " ")

	ymd := dateParts[0]
	ymdParts := strings.Split(ymd, "-")
	year := ymdParts[0]
	month := ymdParts[1]
	day := ymdParts[2]

	hmd := dateParts[1]
	tz := dateParts[2]
	sortKey := fmt.Sprintf("%s %s %s", ymd, hmd, tz)

	cat, err := readAttrLine(path, "categories:")
	if err != nil {
		return nil, err
	}

	// ファイル名先頭の日付部分を削除
	re := regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}-`)
	base := re.ReplaceAllString(filepath.Base(path), "")

	// 拡張子を削除
	fileName := strings.TrimSuffix(base, ".md")

	fa := fileAttr{
		title:    title,
		year:     year,
		month:    month,
		day:      day,
		hmd:      hmd,
		tz:       tz,
		sortKey:  sortKey,
		category: cat,
		fileName: fileName,
	}
	return &fa, nil
}

// ToMarkdown は箇条書きのマークダウンリンク文字列を返却する。
//
// text/template 内から関数呼び出しでアクセスするためパブリック関数として定義している。
func (fa *fileAttr) ToMarkdown() string {
	return fmt.Sprintf("* %s-%s-%s %s [%s](/%s/%s/%s/%s/%s.html)",
		fa.year, fa.month, fa.day, categoriesMap[fa.category], fa.title, fa.category, fa.year, fa.month, fa.day, fa.fileName,
	)
}

// generateLinks は見出しと箇条書きを生成する。
func generateLinks(i inventory, w io.Writer) error {
	const tmpl = `
{{- $fa := .FileAttrs -}}
{{- range $year := .LoopKeys -}}
{{- $attrs := index $fa $year -}}
### {{ $year }} 年

{{ range $attr := $attrs -}}
{{ $attr.ToMarkdown }}
{{ end }}
{{ end -}}
`
	t, err := template.New("posts").Parse(tmpl)
	if err != nil {
		return err
	}
	if err := t.Execute(w, i); err != nil {
		return err
	}
	return nil
}
