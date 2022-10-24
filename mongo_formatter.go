package mf

import "regexp"

// Escapes the mongo `key`.

func Escape(key string) string {
	re := regexp.MustCompile(`\\`)
	key = re.ReplaceAllString(key, "\uFF03")
	re = regexp.MustCompile(`\$`)
	key = re.ReplaceAllString(key, "\uFF04")
	re = regexp.MustCompile(`\&`)
	key = re.ReplaceAllString(key, "\uFF05")
	re = regexp.MustCompile(`\+`)
	key = re.ReplaceAllString(key, "\uFF06")
	re = regexp.MustCompile(`\*`)
	key = re.ReplaceAllString(key, "\uFF07")
	re = regexp.MustCompile(`\?`)
	key = re.ReplaceAllString(key, "\uFF08")
	re = regexp.MustCompile(`\]`)
	key = re.ReplaceAllString(key, "\uFF09")
	re = regexp.MustCompile(`\[`)
	key = re.ReplaceAllString(key, "\uFF0A")
	re = regexp.MustCompile(`\(`)
	key = re.ReplaceAllString(key, "\uFF0B")
	re = regexp.MustCompile(`\)`)
	key = re.ReplaceAllString(key, "\uFF0C")
	re = regexp.MustCompile(`\|`)
	key = re.ReplaceAllString(key, "\uFF0D")
	re = regexp.MustCompile(`\.`)
	key = re.ReplaceAllString(key, "\uFF0E")
	re = regexp.MustCompile(`\{`)
	key = re.ReplaceAllString(key, "\uFF0F")
	re = regexp.MustCompile(`\}`)
	key = re.ReplaceAllString(key, "\uFF0G")
	return key
}

// Unescapes the mongo `key`.

func Unescape(key string) string {
	re := regexp.MustCompile("\uFF03")
	key = re.ReplaceAllString(key, `\`)
	re = regexp.MustCompile("\uFF04")
	key = re.ReplaceAllString(key, `$`)
	re = regexp.MustCompile("\uFF05")
	key = re.ReplaceAllString(key, `&`)
	re = regexp.MustCompile("\uFF06")
	key = re.ReplaceAllString(key, `+`)
	re = regexp.MustCompile("\uFF07")
	key = re.ReplaceAllString(key, `*`)
	re = regexp.MustCompile("\uFF08")
	key = re.ReplaceAllString(key, `?`)
	re = regexp.MustCompile("\uFF09")
	key = re.ReplaceAllString(key, `]`)
	re = regexp.MustCompile("\uFF0A")
	key = re.ReplaceAllString(key, `[`)
	re = regexp.MustCompile("\uFF0B")
	key = re.ReplaceAllString(key, `(`)
	re = regexp.MustCompile("\uFF0C")
	key = re.ReplaceAllString(key, `)`)
	re = regexp.MustCompile("\uFF0D")
	key = re.ReplaceAllString(key, `|`)
	re = regexp.MustCompile("\uFF0E")
	key = re.ReplaceAllString(key, `.`)
	re = regexp.MustCompile("\uFF0F")
	key = re.ReplaceAllString(key, `{`)
	re = regexp.MustCompile("\uFF0G")
	key = re.ReplaceAllString(key, `}`)
	return key
}