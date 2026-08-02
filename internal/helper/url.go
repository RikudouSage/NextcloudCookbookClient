package helper

import "strings"

func JoinURLPath(basePath, path string) string {
	parts := make([]string, 0, 2)
	if base := strings.Trim(basePath, "/"); base != "" {
		parts = append(parts, base)
	}
	if suffix := strings.Trim(path, "/"); suffix != "" {
		parts = append(parts, suffix)
	}
	if len(parts) == 0 {
		return "/"
	}

	return "/" + strings.Join(parts, "/")
}
