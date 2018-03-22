// +build ignore

// This program generates html.go/css.go. It can be invoked by running
// go generate
package main

import (
	"path/filepath"
	"strings"
	"io/ioutil"
	"os"
	"github.com/kib357/less-go"
	"github.com/brokenbydefault/Nanollet/Util"
)

func main() {
	generateSciter()
	generateLESS()
	generateHTML()
	generateCSS()
}

//@TODO Use template/text instead and rewrite the code
//@TODO Support Linux/Darwin

func generateSciter() {
	f, _ := os.Create("GUI/Front/sciter_windows.go")

	dll, err := ioutil.ReadFile("sciter.dll")
	if err != nil {
		panic(err)
	}
	hex := Util.UnsafeHexEncode(dll)

	f.WriteString("// Code generated by go generate; DO NOT EDIT. \npackage Front \n\n var Sciter = []byte{")

	var sciter []byte
	for i := 0; i < len(hex); i += 2 {
		// 0xHH,
		sciter = append(sciter, 0x30, 0x78, byte(hex[i]), byte(hex[i+1]), 0x2C)
	}

	f.Write(sciter[:len(sciter)-1])
	f.WriteString("}")
}

func generateLESS() {
	err := less.RenderFile("GUI/Front/less/style.less", "GUI/Front/css/style.css", map[string]interface{}{"compress": true})
	if err != nil {
		panic(err)
	}
}

func generateCSS() {
	f, _ := os.Create("GUI/Front/css.go")

	_, err := f.WriteString("// Code generated by go generate; DO NOT EDIT. \npackage Front \n\n");
	if err != nil {
		panic(err)
	}

	css, err := ioutil.ReadFile("GUI/Front/css/style.css")
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("const CSSStyle = `" + string(css) + "` \n");
	if err != nil {
		panic(err)
	}
}

func generateHTML() {
	f, _ := os.Create("GUI/Front/html.go")

	_, err := f.WriteString("// Code generated by go generate; DO NOT EDIT. \npackage Front \n\ntype HTMLPAGE string\n\n");
	if err != nil {
		panic(err)
	}

	gb, err := filepath.Glob("GUI/Front/html/*")
	if err != nil {
		panic(err)
	}

	for _, path := range gb {
		html, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString("const HTML" + strings.Title(strings.Replace(filepath.Base(path), ".html", "", 1)) + " HTMLPAGE = `" + string(html) + "` \n");
		if err != nil {
			panic(err)
		}
	}
}