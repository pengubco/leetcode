package simplify_path

import "strings"

func simplifyPath(path string) string {
	if len(path) == 1 {
		return path
	}
	var stack []string
	parts := strings.Split(path[1:], "/")
	for _, p := range parts {
		if p == "." || p == "" {
			continue
		}
		if p == ".." {
			m := len(stack)
			if m >= 1 {
				stack = stack[:m-1]
			}
			continue
		}
		stack = append(stack, p)
	}
	return "/" + strings.Join(stack, "/")
}
