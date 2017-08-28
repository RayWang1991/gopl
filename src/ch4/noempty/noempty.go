package main

// in place
func noEmpty(strings []string) []string {
	i := 0
	for _, str := range strings {
		if str != "" {
			strings[i] = str
			i++
		}
	}
	return strings[:i]
}

// not in place
func noEmpty1(strings []string) []string {
	out := []string{}
	for _, str := range strings {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

func noEmpty2(strings []string) []string {
	out := strings[:0]
	for _, str := range strings {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}
