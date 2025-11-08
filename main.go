package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/ManasRaut/md_lex/converters"
)

//go:embed tests/sample_markdown.md
var exmapleSource string

func main() {

	mdToElementCompiler, err := NewMDLexCompiler(converters.HTMLConverter{})
	if err != nil {
		panic(err)
	}
	html, err := mdToElementCompiler.Compile(strings.NewReader(exmapleSource))
	if err != nil {
		panic(err)
	}

	fmt.Println(*html)

	output := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
	<link rel="stylesheet" href="./style.css">
</head>
<body>
<main>
%s
</main>
</body>
</html>
`, string(*html))
	os.WriteFile("./result.html", []byte(output), 0644)

}
