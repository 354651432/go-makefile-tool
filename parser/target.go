package parser

import (
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

type Target struct {
	Name     string
	Deps     []string
	Commands []string
}

var (
	fileName string
)

func init() {
	fileName = "Makefile"
}

func SetFile(file string) {
	fileName = file
}

func GetFile() string {
	return fileName
}

func Parse() ([]*Target, error) {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileStr := string(fileContent)
	lines := strings.Split(fileStr, "\n")

	var currentTarget *Target
	result := make([]*Target, 0)
	for _, line := range lines {
		if isTargetDefinedLine(line) {
			if currentTarget != nil {
				result = append(result, currentTarget)
			}

			currentTarget = parseTargetNameAndDeps(line)
		}
		if len(line) > 0 && unicode.IsSpace(rune((line[0]))) {
			currentTarget.Commands = append(currentTarget.Commands, strings.Trim(line, " "))
		}
	}

	if currentTarget != nil {
		result = append(result, currentTarget)
	}
	return result, nil
}

func isTargetDefinedLine(str string) bool {
	if len(str) > 0 && str[0] == '$' {
		return false
	}
	return strings.Contains(str, ":")
}

func parseTargetNameAndDeps(str string) *Target {
	arr := strings.Split(str, ":")
	if len(arr) == 1 {
		return &Target{
			Name: arr[0],
		}
	}
	deps := strings.Split(strings.Trim(arr[1], " "), " ")

	return &Target{
		Name: arr[0],
		Deps: deps,
	}
}
