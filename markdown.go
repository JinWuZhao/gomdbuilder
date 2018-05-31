package gomdbuilder

import (
	"fmt"
	"strings"
)

const (
	BR  = "  "
	HLE = "===" + LE
	HR  = LE + "---" + LE
)

func Doc(docs ...string) string {
	return strings.Join(docs, LE)
}

func H(level int, text ...interface{}) string {
	return fmt.Sprintf("%s %s", strings.Repeat("#", level),
		strings.Join(stringifySlice(text), ""))
}

func H1(text ...interface{}) string {
	return H(1, text...)
}

func H2(text ...interface{}) string {
	return H(2, text...)
}

func H3(text ...interface{}) string {
	return H(3, text...)
}

func H4(text ...interface{}) string {
	return H(4, text...)
}

func H5(text ...interface{}) string {
	return H(5, text...)
}

func H6(text ...interface{}) string {
	return H(6, text...)
}

func P(text ...interface{}) string {
	return strings.Join(stringifySlice(text), "") + BR
}

func I(text ...interface{}) string {
	return fmt.Sprintf(" _%s_ ", strings.Join(stringifySlice(text), ""))
}

func B(text ...interface{}) string {
	return fmt.Sprintf(" **%s** ", strings.Join(stringifySlice(text), ""))
}

func M(text ...interface{}) string {
	return fmt.Sprintf(" `%s` ", strings.Join(stringifySlice(text), ""))
}

func BLi(items ...interface{}) string {
	return strings.Join(rangeStrings(func(_ int, s string) string {
		return "- " + s
	}, stringifySlice(items)), BR+LE)
}

func NLi(items ...interface{}) string {
	return strings.Join(rangeStrings(func(i int, s string) string {
		return fmt.Sprintf("%d. ", i) + s
	}, stringifySlice(items)), BR+LE)
}

func Ln(title string, url string) string {
	return fmt.Sprintf("[%s](%s)", title, url)
}

func Img(title string, url string) string {
	return fmt.Sprintf("![%s](%s)", title, url)
}

func Q(level int, text ...interface{}) string {
	return fmt.Sprintf("%s %s", strings.Repeat(">", level), strings.Join(stringifySlice(text), "")) + LE
}

func Q1(text ...interface{}) string {
	return Q(1, text...)
}

func Q2(text ...interface{}) string {
	return Q(2, text...)
}

func Q3(text ...interface{}) string {
	return Q(3, text...)
}

func Q4(text ...interface{}) string {
	return Q(4, text...)
}

func Q5(text ...interface{}) string {
	return Q(5, text...)
}

func Q6(text ...interface{}) string {
	return Q(6, text...)
}

func Code(lang string, codes ...interface{}) string {
	return fmt.Sprintf("```%s\n%s\n```", lang, strings.Join(stringifySlice(codes), ""))
}
