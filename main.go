package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var (
		filename string
		outfile  string
		sep      string
		header   bool
	)

	flag.StringVar(&filename, "f", "", "path to input file")
	flag.StringVar(&outfile, "o", "", "path to output file")
	flag.StringVar(&sep, "s", ",", "separator")
	flag.BoolVar(&header, "header", true, "indicates whether first row of input is headers")
	flag.Parse()

	var in io.Reader = os.Stdin
	var out io.Writer = os.Stdout

	if filename != "" {
		f, err := os.Open(filename)
		check(err)
		defer f.Close()
		in = f
	}

	if outfile != "" {
		f, err := os.Create(outfile)
		check(err)
		defer f.Close()
		out = f
	}

	r := csv.NewReader(in)
	r.Comma = rune(sep[0])
	w := bufio.NewWriter(out)

	for i := 0; ; i++ {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		tr := mdTableRow(rec)
		_, err = w.WriteString(tr + "\n")
		check(err)

		if i == 0 && header {
			vals := []string{}
			for j := 0; j < r.FieldsPerRecord; j++ {
				vals = append(vals, " --- ")
			}
			tr := mdTableRow(vals)
			_, err = w.WriteString(tr + "\n")
			check(err)
		}
	}
	w.Flush()
}

func mdTableRow(vals []string) string {
	if len(vals) <= 0 {
		return ""
	}

	var row strings.Builder
	for _, v := range vals {
		row.WriteString("| ")
		row.WriteString(v)
	}
	row.WriteString(" |")

	return row.String()
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
