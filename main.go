package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
)

func main() {
	quietly := flag.Bool("q", false, "turn off output")
	showList := flag.Bool("list", false, "show required list")
	binDir := flag.String("dir", filepath.Join(".", "bin"), "bin directory")
	flag.Parse()

	log.SetPrefix("depinst: ")
	log.SetFlags(0)

	tree, err := toml.LoadFile("Gopkg.toml")
	if err != nil {
		log.Fatal(err)
	}

	var goPkg struct {
		Required []string `toml:"required"`
	}
	if err := tree.Unmarshal(&goPkg); err != nil {
		log.Fatal("broken Gopkg.toml: ", err)
	}

	if *showList {
		for _, p := range goPkg.Required {
			fmt.Println(p)
		}
		return
	}

	for _, p := range goPkg.Required {
		name := path.Base(p)
		binPath := filepath.Join(*binDir, name)
		pkgPath := "." + string(filepath.Separator) + filepath.Join("vendor", filepath.FromSlash(p))

		cmd := exec.Command("go", "build", "-o", binPath, pkgPath)
		if !*quietly {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		if !*quietly {
			log.Print("running ", cmd.Args, " ...")
		}
		if err := cmd.Run(); err != nil {
			log.Fatalf("failed to run (%v): %v", strings.Join(cmd.Args, " "), err)
		}
	}
}
