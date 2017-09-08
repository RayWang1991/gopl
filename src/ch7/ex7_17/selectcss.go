package main

import (
	"encoding/xml"
	"os"
	"io"
	"fmt"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xml err\n", err)
			os.Exit(1)
		}
		switch token := token.(type) {
		case xml.StartElement:
			stack = append(stack, token.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			//if isSelected(stack, os.Args[1:]) {
			//	fmt.Printf("%s :%s\n", strings.Join(stack, " "), string(token))
			//}
		}
	}
}

func isSelected(stack, want []string) bool {
	if len(stack) < len(want) {
		return false
	}

	i, j := len(stack)-1, len(want)-1
	for j > 0 {
		if stack[i] != want[j] {
			return false
		}
		i--
		j--
	}
	return true
}

type cssmap map[string]string

func (c cssmap) isElementSelected(element xml.StartElement) bool {
	if len(c) > len(element.Attr) {
		return false
	}
	for k, v := range c {
		found := false
		for _, a := range element.Attr {
			if a.Name.Local != k {
				found = false
			} else {
				if a.Value == v {
					found = true
				} else {
					found = false
				}
			}
		}
		if !found {
			return false
		}
	}
	return true
}