package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var categoriesMap = map[string]string{
	"tech":  "技術",
	"movie": "映画",
	"game":  "ゲーム",
	"daily": "雑記",
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

func main() {
	paths := readStdin()

	// ファイルを読み込んで属性を取得する
	fas := make([]*fileAttr, 0)
	for _, path := range paths {
		fa, err := readFileAttr(path)
		if err != nil {
			panic(err)
		}
		fas = append(fas, fa)
	}

	sort.Slice(fas, func(i, j int) bool {
		return fas[i].sortKey > fas[j].sortKey
	})

	// 年で見出しを作る
	yearFound := make(map[string]bool)
	markdownLines := make([]string, 0)
	for _, fa := range fas {
		year := fa.year
		_, ok := yearFound[year]
		if !ok {
			yearFound[year] = true
			heading := fmt.Sprintf("### %s 年", year)
			markdownLines = append(markdownLines, "")
			markdownLines = append(markdownLines, heading)
			markdownLines = append(markdownLines, "")
		}
		markdownLines = append(markdownLines, fa.toMarkdown())
	}

	for _, l := range markdownLines {
		fmt.Println(l)
	}
}

// readStdin は標準入力を文字列の配列として返す。
func readStdin() []string {
	ret := make([]string, 0)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		ret = append(ret, line)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return ret
}

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
		panic(err)
	}
	return "", fmt.Errorf("%s doesn't exist: file = %s", attr, path)
}

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

	fa := fileAttr{
		title:    title,
		year:     year,
		month:    month,
		day:      day,
		hmd:      ymd,
		tz:       tz,
		sortKey:  sortKey,
		category: cat,
		fileName: strings.TrimSuffix(filepath.Base(path)[11:], ".md"),
	}
	return &fa, nil
}

func (fa *fileAttr) toMarkdown() string {
	return fmt.Sprintf("* %s/%s/%s %s [%s](/%s/%s/%s/%s/%s.html)",
		fa.year, fa.month, fa.day, categoriesMap[fa.category], fa.title, fa.category, fa.year, fa.month, fa.day, fa.fileName,
	)
}
