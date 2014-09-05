package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

var path string
var repo string
var p string

func init() {
	flag.StringVar(&path, "o", "version.go", "filename")
	flag.StringVar(&repo, "i", ".", "repository")
	flag.StringVar(&p, "p", "version", "package")

}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func main() {

	flag.Parse()

	cmd := exec.Command("git", "--git-dir", repo+"/.git", "rev-parse", "HEAD")
	bytes, err := cmd.CombinedOutput()
	file_contents := fmt.Sprintf(Template, p, stripchars(string(bytes), "\r\n "))

	if err = ioutil.WriteFile(path, []byte(file_contents), 0644); err != nil {
		log.Fatal(err)
	}
}

var Template = `package %s

var GIT_COMMIT_HASH = "%s"`
