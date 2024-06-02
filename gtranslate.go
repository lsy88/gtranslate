package gtranslate

import (
	"golang.org/x/text/language"
	"strings"
	"time"
	"unicode/utf8"
)

var GoogleHost = "google.com"
var limit = 5000 // 单次翻译最大限制数

// TranslationParams is a util struct to pass as parameter to indicate how to translate
type TranslationParams struct {
	From       string
	To         string
	Tries      int
	Delay      time.Duration
	GoogleHost string
	Proxy      string
}

// translateChunk block translation
func translateChunk(text, from, to string, withVerification bool, tries int, delay time.Duration, proxy string) (string, error) {
	if text == "" {
		return "", nil
	}
	tSlice := strings.Split(text, "\n")
	tLength := len(tSlice)
	
	var builder strings.Builder
	var total int
	for i, j := 0, 0; j <= tLength; {
		if j < tLength {
			total += utf8.RuneCountInString(tSlice[j])
		}
		
		if total > limit || j == tLength {
			s, err := translate(strings.Join(tSlice[i:j], "\n"), from, to, withVerification, tries, delay, proxy)
			if err != nil {
				return "", err
			}
			i = j
			builder.WriteString(s)
			if j < tLength {
				builder.WriteString("\n")
			}
			total = 0
		}
		j++
	}
	
	return builder.String(), nil
}

// Translate a text using native tags offer by go language
func Translate(text string, from language.Tag, to language.Tag, googleHost ...string) (string, error) {
	if len(googleHost) != 0 && googleHost[0] != "" {
		GoogleHost = googleHost[0]
	}
	translated, err := translateChunk(text, from.String(), to.String(), false, 2, 0, "")
	if err != nil {
		return "", err
	}
	
	return translated, nil
}

// TranslateWithParams translate a text with simple params as string
func TranslateWithParams(text string, params TranslationParams) (string, error) {
	if params.GoogleHost == "" {
		GoogleHost = "google.com"
	} else {
		GoogleHost = params.GoogleHost
	}
	translated, err := translateChunk(text, params.From, params.To, true, params.Tries, params.Delay, params.Proxy)
	if err != nil {
		return "", err
	}
	return translated, nil
}
