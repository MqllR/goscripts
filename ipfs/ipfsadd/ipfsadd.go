package main

import (
	"flag"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var (
	Host  string
	Path  string
	Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
)

func init() {
	flag.StringVar(&Host, "host", os.Getenv("IPFS_API_HOST"),
		"Host with port IPFS API. Env variable IPFS_API_HOST. "+
			"Eg: ipfs.com:5001, 111.111.222.222:5001",
	)

	flag.StringVar(&Path, "path", os.Getenv("IPFS_ADD_PATH"),
		"File or directory add. Env variable IPFS_ADD_PATH",
	)

	flag.Parse()

	if Host == "" {
		fmt.Fprintf(os.Stderr, "Host is not defined")
		os.Exit(1)
	}

	if Path == "" {
		fmt.Fprintf(os.Stderr, "Path is not defined")
		os.Exit(1)
	}
}

func main() {
	sh := shell.NewShell(Host)

	file, err := os.Stat(Path)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Path does not exist")
		os.Exit(1)
	}

	var cid string
	if file.IsDir() {
		cid, err = sh.AddDir(Path)
	} else {
		f, err := os.Open(Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			os.Exit(1)
		}
		cid, err = sh.Add(f)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("IPFS object: %s\n", cid)
}
