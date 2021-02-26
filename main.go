// getDoc
// 2019, 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"flag"
	"fmt"
	"github.com/HuguesGuilleus/getDoc/pkg"
	"io"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

var (
	printVersion = flag.Bool("version", false, "Print the version")
	verbose      = flag.Bool("v", false, "Enable verbose mode")
	output       = flag.String("o", "doc.html", "The output file (use extesion to get the output format: HTML(default), JSON or XML)")
	listLine     = flag.Bool("debug", false, "Only list the type of each lines")
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
	} else if *listLine {
		doc.ReadDebug(flag.Args())
		return
	}

	if *verbose {
		log.SetPrefix("--- ")
		log.SetFlags(0)
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}

	var ind doc.Index
	if args := flag.Args(); len(args) == 0 {
		ind = *doc.Read(".")
	} else {
		for _, a := range args {
			ind = append(ind, *doc.Read(a)...)
		}
	}

	switch {
	case strings.HasSuffix(*output, ".json"):
		ind.DataIndex().Json(*output)
	case strings.HasSuffix(*output, ".xml"):
		ind.DataIndex().Xml(*output)
	default:
		ind.SaveHTML(*output)
	}
}
