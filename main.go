// getDoc
// 2019, 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"flag"
	"fmt"
	"github.com/HuguesGuilleus/getDoc/pkg"
	"log"
	"os"
	"path"
	"runtime/debug"
	"strings"
)

var (
	printVersion = flag.Bool("version", false, "Print the version")
	verbose      = flag.Bool("v", false, "Enable verbose mode")
	output       = flag.String("o", "doc.html", "The output file (use extesion to get the output format: HTML(default), JSON or XML)")
	title        = flag.String("t", "", "The title of this doc")
	indent       = flag.Bool("indent", false, "Enable indent for JSON, XML and HTML")
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage of getDoc: [OPTION] [inputs]")
		fmt.Println()
		fmt.Println("Inputs can be directorys or files, by default is the")
		fmt.Println("current directory. Directory are readed recurively.")
		fmt.Println()
		log.Println("SUPORTED LANGUAGES:")
		for _, lang := range []string{"bash", "c", "go", "js"} {
			fmt.Println("  -", lang)
		}
		fmt.Println()

		fmt.Println("OPTIONS:")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *printVersion {
		if info, _ := debug.ReadBuildInfo(); info != nil {
			fmt.Println("getDoc", info.Main.Version)
			fmt.Println(info.Main.Sum)
		} else {
			fmt.Println("getDoc unkown version")
		}
		return
	}

	var d doc.Doc

	if *verbose {
		d.Log.SetPrefix("--- ")
		d.Log.SetOutput(os.Stdout)
	}

	if *title == "" {
		d.Title = path.Clean(flag.Arg(0))
	} else {
		d.Title = *title
	}

	if args := flag.Args(); len(args) == 0 {
		d.Read(".", os.DirFS("."))
	} else {
		for _, a := range args {
			d.Read(a, os.DirFS(a))
		}
	}

	out, err := os.Create(*output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create file for save output %q fail: %v\n", *output, err)
		os.Exit(1)
	}
	defer out.Close()

	failSave := func(err error) {
		if err != nil {
			fmt.Fprintln(os.Stderr, "Save the doc fail: ", err)
			os.Exit(1)
		}
	}

	switch {
	case strings.HasSuffix(*output, ".json"):
		failSave(d.SaveJSON(out, *indent))
	case strings.HasSuffix(*output, ".xml"):
		failSave(d.SaveXML(out, *indent))
	default:
		failSave(d.SaveHTML(out, *indent))
	}
}
