# gtranslate

二开<https://github.com/bregydoc/gtranslate>，底层调用google翻译api

由于有道翻译，腾讯交互翻译，谷歌翻译都存在单次翻译最大5000字符限制，对于大文本翻译非常麻烦，在原功能基础上增加代理和大文本分块翻译功能
# Install

    go get github.com/lsy88/gtranslate

# Use

```go
gtranslate.Translate("I'm alive", language.English, language.Spanish)
```

```go
gtranslate.TranslateWithParams("I'm alive", gtranslate.TranslateWithParams{From: "en", To: "es"})
```

# Example

```go
package main

import (
	"fmt"

	"github.com/lsy88/gtranslate"
)

func main() {
	text := "Hello World"
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | ja: %s \n", text, translated)
	// en: Hello World | ja: こんにちは世界
}
```
