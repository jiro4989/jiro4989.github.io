package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file := "index.md"
	postLinks := readStdin()

	contents, err := embedLinks(file, postLinks)
	if err != nil {
		panic(err)
	}

	if err := writeFile(file, contents); err != nil {
		panic(err)
	}
}

// readStdin は標準入力を文字列の配列として返す。
func readStdin() []string {
	ret := make([]string, 0)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		ret = append(ret, line)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return ret
}

func embedLinks(path string, postLinks []string) ([]string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	isPosts := false
	contents := make([]string, 0)
	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		// ここはトリムしてはいけない
		line := sc.Text()
		if strings.Contains(line, "START_POSTS") {
			isPosts = true
			contents = append(contents, line)
			for _, link := range postLinks {
				contents = append(contents, link)
			}
		} else if strings.Contains(line, "END_POSTS") {
			isPosts = false
		}

		if !isPosts {
			contents = append(contents, line)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	return contents, nil
}

func writeFile(path string, contents []string) error {
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	for _, c := range contents {
		fp.WriteString(c)
		fp.WriteString("\n")
	}

	return nil
}
