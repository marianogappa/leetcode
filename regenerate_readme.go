package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	var (
		data   = prepareData(mustListDirectorysDirs("2020"))
		readme = mustCreateReadmeFile()
	)
	defer readme.Close()
	mustExecuteTemplate(readme, data)
}

func prepareData(dirs []string) templateData {
	exercises := []exercise{}
	for _, dir := range dirs {
		exercises = append(exercises, exercise{
			LeetcodeLink: fmt.Sprintf("https://leetcode.com/problems/%v", dir),
			SolutionLink: fmt.Sprintf("2020/%v", dir),
			Title:        snakeToWord(dir),
		})
	}
	return templateData{Exercises: exercises, Total: len(exercises)}
}

func snakeToWord(s string) string {
	return fmt.Sprintf("%v%v", strings.ToUpper(string(s[0])), strings.Replace(s[1:], "-", " ", -1))
}

func mustListDirectorysDirs(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	list := []string{}
	for _, f := range files {
		if f.IsDir() && f.Name()[0] != '.' {
			list = append(list, f.Name())
		}
	}

	sort.Strings(list)
	return list
}

func mustCreateReadmeFile() io.WriteCloser {
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

type exercise struct {
	LeetcodeLink string
	SolutionLink string
	Title        string
}

type templateData struct {
	Exercises []exercise
	Total     int
}

func mustExecuteTemplate(fd io.Writer, data templateData) {
	t, err := template.New("readme").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(fd, data); err != nil {
		log.Fatal(err)
	}
}

var templateString = `## 2020 Leetcode exercises ({{.Total}} solved)

|Title|Leetcode|My Solution|
|-----|:--------:|:---------:|
{{ range .Exercises -}}
|{{.Title}}|[[Link]]({{.LeetcodeLink}})|[[Link]]({{.SolutionLink}})|
{{ end }}
`
