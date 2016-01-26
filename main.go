package main

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
	"text/template"
	"time"

	"github.com/bountylabs/log"
)

var path string
var repo string
var p string
var short bool

func init() {
	flag.StringVar(&path, "o", "version.go", "filename")
	flag.StringVar(&repo, "i", ".", "repository")
	flag.StringVar(&p, "p", "version", "package")
	flag.BoolVar(&short, "s", false, "short git sha")
}

func main() {

	flag.Parse()

	//get commit hash (short or not)
	cmd := func() *exec.Cmd {
		if short {
			return exec.Command("git", "--git-dir", repo+"/.git", "rev-parse", "--short", "HEAD")
		}

		return exec.Command("git", "--git-dir", repo+"/.git", "rev-parse", "HEAD")
	}()

	git, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("get git version err=%v", err)
		os.Exit(1)
	}

	data := struct {
		Pkg    string
		Git    string
		Build  string
		Tstamp int64
	}{
		Pkg:    p,
		Git:    string(bytes.TrimSpace(git)),
		Build:  os.Getenv("CIRCLE_BUILD_NUM"),
		Tstamp: time.Now().UnixNano(),
	}

	f, err := os.Create(path)
	if err != nil {
		log.Errorf("create file=%s err=%v", path, err)
		os.Exit(1)
	}

	if err = tmpl.Execute(f, &data); err != nil {
		log.Errorf("execute template err=%v", err)
		os.Exit(1)
	}

	if err = f.Close(); err != nil {
		log.Errorf("close file=%s err=%v", f.Name(), err)
		os.Exit(1)
	}
}

var tmpl = template.Must(template.New("gitversion").Parse(`package {{.Pkg}}

var GIT_COMMIT_HASH = "{{.Git}}"
var CIRCLE_BUILD_NUM = "{{.Build}}"
var GENERATED = {{.Tstamp}}
`))
