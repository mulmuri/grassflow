package format

import (
	"path/filepath"
	"strings"
)



func ReplaceD(form string, replaceTo string) string {
	return strings.Replace(form, "{d}", replaceTo, 1)
}

func ReplaceF(form string, replaceTo string) string {
	return strings.Replace(form, "{f}", replaceTo, 1)
}

func ReplaceN(form string, replaceTo string) string {
	return strings.Replace(form, "{n}", strings.TrimSuffix(replaceTo, filepath.Ext(replaceTo)), 1)
}