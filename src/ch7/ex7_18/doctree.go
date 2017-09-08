package main

import (
	"encoding/xml"
	"io"
	"fmt"
	"bytes"
	"os"
)

// CharData or *Element
type Node interface {
	String() string
}

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	buf := bytes.Buffer{}
	// name
	buf.WriteString(e.Type.Local)
	// attr
	if len(e.Attr) > 0 {
		buf.WriteByte('<')
		for i, a := range e.Attr {
			if i > 0 {
				buf.WriteByte(' ')
				buf.WriteByte(',')
			}
			buf.WriteString(a.Name.Local)
			buf.WriteByte('=')
			buf.WriteString(a.Value)
		}
		buf.WriteByte('>')
	}
	// children
	buf.WriteByte('(')
	for _, c := range e.Children {
		buf.WriteString(c.String())
		buf.WriteByte(' ')
	}
	buf.WriteByte(')')
	return buf.String()
}

func decode(r io.Reader) (Node, error) {
	stack := []*Element{}
	dec := xml.NewDecoder(r)
	// the first token must be the doc root
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			c := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 1 {
				last := stack[len(stack)-1]
				last.Children = append(last.Children, c)
			}
			stack = append(stack, c)
		case xml.EndElement:
			// check the name must be equal to the start one
			last := stack[len(stack)-1]
			if last.Type.Local != tok.Name.Local {
				return nil, fmt.Errorf("unmatched token: %s %s", last, tok)
			}
			// ignore the tailing ends
			if len(stack) == 1 {
				return last, nil
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			last := stack[len(stack)-1]
			last.Children = append(last.Children, CharData(tok))
		}
	}
	if len(stack) != 0 {
		buf := bytes.Buffer{}
		for _, s := range stack {
			buf.WriteString(s.Type.Local)
			buf.WriteByte(' ')
		}
		return nil, fmt.Errorf("unclosed token %s\n", buf.String())
	}
	return nil, fmt.Errorf("empty xml")
}

func main() {
	doc, err := decode(os.Stdin)
	if err != nil {
		fmt.Printf("Error :%s\n", err)
		return
	}
	fmt.Printf("Doc tree : %s", doc.String())
}
