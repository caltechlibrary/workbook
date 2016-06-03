package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	// Caltech Library packages
	"github.com/caltechlibrary/ead2002"
	"github.com/caltechlibrary/workbook"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help information")
	flag.BoolVar(&showVersion, "v", false, "display version information")
	flag.BoolVar(&showLicense, "l", false, "display license information")
}

func main() {
	appname := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	if showHelp == true {
		fmt.Printf(`
 USAGE: %s [OPTIONS] XLSX_FILENAME

 Demonstrate reading an Excel Workbook in xlsx format and rendering
 an EAD XML document.

 OPTIONS
`, appname)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("    -%s  (defaults to %s) %s\n", f.Name, f.DefValue, f.Usage)
		})
		fmt.Printf("\n\n Version: %s\n", ead2002.Version)
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Printf(" Version: %s\n", ead2002.Version)
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Printf(`
`, appname)
		os.Exit(0)
	}

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Missing XLSX_FILENAME, try: %s -h\n", appname)
		os.Exit(1)
	}

	fname := args[0]
	wb, _ := workbook.NewFromExcelFilename(fname)
	sNames := wb.GetSheetNames()
	for i, sName := range sNames {
		sh, _ := wb.GetSheet(i)
		c := sh.RowCount()
		fmt.Printf("%d - %s, row cnt: %d\n", i, sName, c)
	}

	fmt.Printf("DEBUG process %s\n", fname)
}
