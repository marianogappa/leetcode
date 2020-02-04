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
		data2020 = prepareData("2020", mustListDirectorysDirs("2020"))
		data2018 = prepareData("2018", mustListDirectorysDirs("2018"))
		readme   = mustCreateReadmeFile()
	)
	defer readme.Close()
	mustExecuteTemplate(readme, templateData{[]yearTemplateData{data2020, data2018}})
}

func prepareData(year string, dirs []string) yearTemplateData {
	exercises := []exercise{}
	for _, dir := range dirs {
		exercises = append(exercises, exercise{
			LeetcodeLink: fmt.Sprintf("https://leetcode.com/problems/%v", dir),
			SolutionLink: fmt.Sprintf("%v/%v", year, dir),
			Title:        snakeToWord(dir),
		})
	}
	return yearTemplateData{Exercises: exercises, Total: len(exercises), Year: year}
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
	Years []yearTemplateData
}

type yearTemplateData struct {
	Exercises []exercise
	Total     int
	Year      string
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

var templateString = `{{ range .Years}}## {{.Year}} Leetcode exercises ({{.Total}} solved)

|Title|Leetcode|My Solution|
|-----|:--------:|:---------:|
{{ range .Exercises -}}
|{{.Title}}|[[Link]]({{.LeetcodeLink}})|[[Link]]({{.SolutionLink}})|
{{ end }}

{{end}}`
