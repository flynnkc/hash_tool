package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var comparisonString string // Hash string passed as argument to compare with
var filename string         // File provided by -f option
const hashError = "Unsupported algorithm provided, supported options include: \n\t - md5 \n\t - sha1 \n\t - sha256"

func init() {
	// Text snippets to be called by help flags
	const (
		toolName string = "hashtool"
		useFlag  string = "Hash to compare output with"
		fileFlag string = "File name or path"
		hint     string = "Note: File or plaintext string must be provided, but not both\n"
	)

	flag.StringVar(&comparisonString, "m", "", useFlag)
	flag.StringVar(&filename, "f", "", fileFlag)
	flag.StringVar(&filename, "file", "", fileFlag)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-f|--file filename] [-m match string] hashing algorithm [plaintext string]\n%s", toolName, hint)
		flag.PrintDefaults()
	}
}

func main() {
	var target string
	var hashAlgo string
	flag.Parse()

	// Check if correct number of arguments are supplied and file exists if filename is populated
	CheckArgs(&target)

	// Get supplied args for hash algorithm and file/string
	hashAlgo = strings.ToLower(flag.Arg(0))
	hashedString := make([]byte, 16)

	// Switch statement to control flow
	switch hashAlgo {
	case "md5":
		if len(filename) > 0 {
			hashedString = HashFileMd5(filename)
		} else {
			hashedString = HashStringMd5(target)
		}
	case "sha256":
		if len(filename) > 0 {
			hashedString = HashFileSha256(filename)
		} else {
			hashedString = HashStringSha256(target)
		}
	case "sha1":
		if len(filename) > 0 {
			hashedString = HashFileSha1(filename)
		} else {
			hashedString = HashStringSha1(target)
		}
	default:
		fmt.Println(hashError)
		os.Exit(1)
	}

	if len(comparisonString) > 0 {
		CompareHash(hashAlgo, comparisonString, hashedString)
	} else {
		PrintHash(hashAlgo, hashedString, target)
	}
}

// CheckArgs does basic error checking on supplied arguments and flags
func CheckArgs(t *string) {

	fileSet := false
	n := flag.NArg()

	// if filepath is set check that file exists
	if len(filename) > 0 {
		// Set more authoritative filepath
		filename = SetFilepath(filename)

		err := CheckForFile(filename)
		handleErr(err, log.Fatal)

		// Set target in main to reduce the number of checks for filename
		*t = filename
		fileSet = true
	} else {
		// Set target in main to string to be hashed if file isn't specified
		*t = flag.Arg(1)
	}

	err := CheckNumArgs(n, fileSet)
	handleErr(err, log.Fatal)
}

// CheckNumArgs checks to make sure the correct number of args are supplied
func CheckNumArgs(n int, f bool) error {

	// Check for minimum number of arguments
	if f && n < 1 {
		return errors.New("Not enough arguments:\n\tUse -h or -help for list of options")
	} else if n < 2 && !f {
		return errors.New("No hashing algorithm specified:\n\tUse -h or -help for list of options")
	}

	// Check for excess arguements
	if n > 2 {
		return errors.New("Too many arguments supplied:\n\tUse -h or -help for list of options")
	} else if f && n > 1 {
		return errors.New("Too many arguments supplied:\n\tUse -h or -help for list of options")
	}

	return nil
}

// CheckForFile inspects if the given string is a valid filepath
func CheckForFile(path string) error {

	// Check if filepath path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := fmt.Errorf("%s is not a valid filepath or file does not exist", path)
		return err
	}
	return nil
}

// SetFilepath takes the filename provided by the user and attempts to generate an authoritative filepath
func SetFilepath(fp string) string {

	// Check if absolute path was already provided
	if filepath.IsAbs(fp) {
		return fp
	}
	p, err := os.Getwd()
	handleErr(err, log.Fatal)

	return filepath.Join(p, fp)
}

// CompareHash evaluates if strings match and prints a message
func CompareHash(algo, comparison string, hash []byte) {

	h := fmt.Sprintf("%x", hash)

	if h == comparison {
		fmt.Printf("PASS: The %s checksum of input matches %s\n", algo, comparison)
	} else {
		fmt.Printf("FAIL: The %s checksum of %s is %x\n", algo, filename, hash)
	}
}

// PrintHash prints the checksum of the hash and the algorithm in standard format
func PrintHash(ha string, hs []byte, t string) {
	fmt.Printf("The %s checksum of %s: %x\n", strings.ToUpper(ha), t, hs)
}

// handleErr checks if error is not nil and logs it with appropriate function
// Error string must be formatted prior by fmt.Errorf or other method because f must use a variadic function
// I could easily make formatting possible by implementing handleErrf, but I won't
func handleErr(err error, f func(v ...interface{})) {
	if err != nil {
		f(err)
	}
}
