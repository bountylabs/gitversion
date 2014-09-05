package main

import (
	"flag"
	"log"

	"github.com/libgit2/git2go"
)

var path string
var repo string

func init() {
	flag.StringVar(&path, "o", "", "filename")
	flag.StringVar(&repo, "i", "", "repository")
}

func main() {

	flag.Parse()

	r, err := git.OpenRepository(repo)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := r.Revparse("HEAD")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(f.Flags())
	log.Println(f.From().Id())
	log.Println(f.To())
	//comment

}
