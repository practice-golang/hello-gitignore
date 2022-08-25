package main // import "hello-gitignore"

import (
	"fmt"

	gi "github.com/sabhiram/go-gitignore"
)

func main() {
	fpath := "../.myignore"
	g, err := gi.CompileIgnoreFile(fpath)
	if err != nil {
		panic(err)
	}

	ipath := []string{
		"/home/sample/index.php",
		"/home/sample/index.asp",
		"/home/sample/index.html",
		"/home/practice-golang/go/src/hello-gitignore/haha.test",
		"/home/haha.so",
		"/home/dll/haha.dll",
		"/not_use/haha.exe",
		"use/not_use",
		"use/not_use/hihi.doc",
		"README.md",
		"C:/lib/TODO.md",
	}

	for _, p := range ipath {
		if g.MatchesPath(p) {
			fmt.Println(p, " -> true")
		} else {
			fmt.Println(p)
		}
	}
}
