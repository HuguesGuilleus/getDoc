// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"./doc"
	"fmt"
	"github.com/HuguesGuilleus/parseOpt"
	"log"
	"os"
)

var spec = &parseOpt.SpecList{
	&parseOpt.Spec{
		NameLong:  "version",
		Desc:      "Print the version",
		CBFlag: func() {
			fmt.Printf("getDoc VERSION: %.2f\n", 1.02)
			fmt.Println("  BSD 3-Clause License")
			os.Exit(0)
		},
	},
	&parseOpt.Spec{
		NameShort: "l",
		NameLong:  "lang",
		Desc:      "List all supported programming languages",
		CBFlag: func() {
			log.Println("LANGUAGE:")
			for _, lang := range []string{"bash", "c", "go", "js"} {
				fmt.Println("  -", lang)
			}
			fmt.Println("  More datail on Github")
			os.Exit(0)
		},
	},
	&parseOpt.Spec{
		NameShort: "v",
		NameLong:  "verbose",
		Desc:      "Verbose Mode",
		CBFlag: func() {
			log.SetOutput(os.Stdout)
		},
	},
	&parseOpt.Spec{
		NameShort:  "o",
		NameLong:   "output",
		Desc:       "Set html output file",
		OptionName: "file.html",
		NeedArg:    true,
	},
	&parseOpt.Spec{
		NameShort:  "j",
		NameLong:   "json",
		OptionName: "file.json",
		Desc:       "Set json output file",
		NeedArg:    true,
	},
	&parseOpt.Spec{
		NameShort:  "x",
		NameLong:   "xml",
		OptionName: "file.xml",
		Desc:       "Set xml output file",
		NeedArg:    true,
	},
}

func main() {
	opt := spec.ParseOs()
	// Read file
	var ind *doc.Index
	if len(opt.Option[""]) == 0 {
		ind = doc.Read(".")
	} else {
		ind = &doc.Index{}
		for _, file := range opt.Option[""] {
			*ind = append(*ind, *doc.Read(file)...)
		}
	}
	// Save the HTML
	if len(opt.Option["output"]) != 0 {
		for _, out := range opt.Option["output"] {
			ind.SaveHTML(out)
		}
	} else if len(opt.Option["json"]) == 0 && len(opt.Option["xml"]) == 0 {
		ind.SaveHTML("doc.html")
	}
	// Save in JSON and XML
	if len(opt.Option["json"]) != 0 || len(opt.Option["xml"]) != 0 {
		data := ind.DataIndex()
		for _, file := range opt.Option["json"] {
			data.Json(file)
		}
		for _, file := range opt.Option["xml"] {
			data.Xml(file)
		}
	}
}
