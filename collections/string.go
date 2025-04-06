package collections

import "strings"

type String string

func NewString(str string) String {
	return String(str)
}

func (str String) Unwrap() string {
	return string(str)
}

func (str String) Length() int {
	return len(str)
}

func (str String) IsEmpty() bool {
	return str.Length() == 0 || len(strings.TrimSpace(string(str))) == 0
}

func (str String) Trim() String {
	return String(strings.TrimSpace(string(str)))
}

func (str String) Replace(from String, to String) String {
	return String(strings.ReplaceAll(string(str), string(from), string(to)))
}

func (str String) StartsWith(arg String) bool {
	return strings.HasPrefix(string(str), string(arg))
}

func (str String) EndsWith(arg String) bool {
	return strings.HasSuffix(string(str), string(arg))
}

func (str String) EqualsIgnoreCase(arg String) bool {
	return strings.EqualFold(string(str), string(arg))
}

func (str String) ToUpperCase() String {
	return String(strings.ToUpper(string(str)))
}

func (str String) ToLowerCase() String {
	return String(strings.ToLower(string(str)))
}

func (str String) Capitalize() String {
	return String(strings.ToTitle(string(str)))
}

func (str String) CapitalizeFirst() String {
	return String(strings.ToUpper(string(str[0]))) + str[1:]
}
