package ex5_16

import "bytes"

func Join(sep string, strs ...string) string {
	s := bytes.Buffer{}
	for i, str := range strs {
		if i > 0 {
			s.WriteString(sep)
		}
		s.WriteString(str)
	}
	return s.String()
}
