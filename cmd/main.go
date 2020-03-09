package main

import (
	"flag"
	"fmt"
	"github.com/initialed85/quotanizer/pkg/quota"
	"log"
	"strconv"
)

type flagArrayString []string

func (f *flagArrayString) String() string {
	return fmt.Sprintf("%d", f)
}

func (f *flagArrayString) Set(value string) error {
	*f = append(*f, value)

	return nil
}

type flagArrayInt64 []int64

func (f *flagArrayInt64) String() string {
	return fmt.Sprintf("%d", f)
}

func (f *flagArrayInt64) Set(value string) error {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}

	*f = append(*f, i)

	return nil
}

var paths flagArrayString
var quotaSizes flagArrayInt64
var suffixes flagArrayString

func main() {
	flag.Var(&paths, "path", "a path for applying a quota")
	flag.Var(&quotaSizes, "quota", "a quota to apply (in gigabytes)")
	flag.Var(&suffixes, "suffix", "a suffix to filter by")

	flag.Parse()

	if len(paths) == 0 {
		log.Fatal("no -path flags specified")
	}

	if len(quotaSizes) == 0 {
		log.Fatal("no -quota flags specified")
	}

	if len(suffixes) == 0 {
		log.Print("warning: no suffixes specified; will remove all files that contravene the quota")
	}

	if len(paths) != len(quotaSizes) {
		log.Fatal("unbalanced mixture of -path and -quota flags")
	}

	quotas := make([]quota.Quota, 0)
	for i, q := range quotaSizes {
		quotaSizes[i] = q * 1e+9
		quotas = append(quotas, quota.New(paths[i], suffixes, q*1e+9))
	}

	for _, q := range quotas {
		err := q.Walk()
		if err != nil {
			log.Fatal(err)
		}

		files := q.Candidates()

		err = q.Delete(files)
		if err != nil {
			log.Fatal(err)
		}
	}
}
