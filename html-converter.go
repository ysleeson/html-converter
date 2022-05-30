package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	srcFileName := ""
	if len(os.Args) < 2 {
		fmt.Println("error: please input file name (example: html-converter README.md)")
		return
	}

	srcFileName = os.Args[1]
	fmt.Println("MarkDown file:", srcFileName)
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		log.Fatalf("%s file not found", srcFileName)
	}
	defer srcFile.Close()
	scanner := bufio.NewScanner(srcFile)
	pStarted := false
	html, line := "", ""
	for scanner.Scan() {
		line, pStarted = convertLine_v2(scanner.Text(), pStarted)
		html = html + line
	}
	targetFileName := strings.Replace(srcFileName, "md", "html", 1)
	targetFile, err := os.Create(targetFileName)
	if err != nil {
		log.Fatalf("%s file creation err(%v)", targetFileName, err)
	}
	defer targetFile.Close()

	_, err = targetFile.WriteString(html)
	if err != nil {
		log.Fatalf("%s file creation err(%v)", targetFileName, err)
	}

	fmt.Println(html)
	fmt.Println("HTML file:", targetFileName)
}

func convertLine_v2(content string, pStarted bool) (string, bool) {
	if len(content) == 0 {
		if pStarted {
			return "</p>\n", false
		} else {
			return "", false
		}
	}
	if content[0] == '#' {
		content = convertHeader(content)
		if pStarted {
			content = "</p>\n" + content
			pStarted = false
		}
	} else {
		if !pStarted {
			content = fmt.Sprintf("<p>%s", content)
			pStarted = true
		}
	}
	content = convertLink(content)
	if !pStarted {
		content = content + "\n"
	}
	return content, pStarted
}

func convertHeader(content string) string {
	count := 0
	for i := 0; i < len(content); i++ {
		if content[i] != '#' {
			if content[i] == ' ' {
				break
			} else {
				count = 0
				break
			}
		}
		count++
	}
	if count > 6 {
		return fmt.Sprintf("<p>%s</p>", content)
	}
	return fmt.Sprintf("<h%d>%s</h%d>", count, content[count+1:], count)
}

func convertLink(content string) string {
	//urltext start, url start, url end
	urlIndexes := [][]int{}
	for i := 0; i < len(content); i++ {
		if content[i] == byte('[') {
			urlIndex := []int{}

			for i = i + 1; i < len(content); i++ {
				if content[i] != byte('[') {
					urlIndex = append(urlIndex, i-1)
					break
				}
			}

			for ; i < len(content); i++ {
				ended := false
				if content[i] == byte(']') {
					i++
					if i < len(content) && content[i] == byte('(') {
						urlIndex = append(urlIndex, i)
						for i = i + 1; i < len(content); i++ {
							if content[i] == byte(')') {
								urlIndex = append(urlIndex, i)
								urlIndexes = append(urlIndexes, urlIndex)
								ended = true
								break
							}
						}
					}
				}
				if ended {
					break
				}
			}
		}
	}
	start := 0
	if len(urlIndexes) == 0 {
		return content
	}
	newContent := ""
	//q[0] : index of '['
	//q[1] : index of '('
	//q[2] : index of ')'
	for _, q := range urlIndexes {
		newContent = newContent + fmt.Sprintf("%s<a href=%s>%s</a>", content[start:q[0]], `"`+content[q[1]+1:q[2]]+`"`, content[q[0]+1:q[1]-1])
		start = q[2] + 1
	}
	newContent = newContent + content[start:]
	return newContent
}
func convertLine(content string) string {
	if len(content) == 0 {
		return content
	}
	if content[0] == '#' {
		content = convertHeader(content)
	} else {
		content = fmt.Sprintf("<p>%s</p>", content)
	}
	content = convertLink(content)
	return content
}
