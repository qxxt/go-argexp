package argexp_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/qxxt/go-argexp"
)

func Test_basic(t *testing.T) {
	// this is a sample slices of arguments (string) produced by os.Args
	// It is the equivalent of the following arguments:
	// --message "this is (\") an escape string" --number 123456 --verbose --url=https://google.com -txyzabc "emacs rocks!!"
	sampleArgs := []string{"--message", "this is (\") a\n escape string\"", "--number", "123456", "--verbose", "--url=https://google.com", "-txyzabc", "emacs rocks!!"}
	a := argexp.Marshall(sampleArgs)
	fmt.Println("a: ", a)
	message := argexp.GetString(&a, "--message")
	url := argexp.GetString(&a, "--url")

	// For integer you must convert the string into integer yourself.
	// This package is dumb, thus will not outsmart you
	// by converting string into an integer.
	n := argexp.GetString(&a, "--number")
	number, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Is it even a number?")
	}

	// You can use multiple cases for boolean
	verbose := argexp.GetBool(&a, "--verbose", "-v")
	trace := argexp.GetBool(&a, "--trace", "-t")

	// Gnu style multiple boolean flag can work as well!!
	x := argexp.GetBool(&a, "-x")
	y := argexp.GetBool(&a, "-y")
	z := argexp.GetBool(&a, "-z")

	fmt.Println("parsed arguments:")
	fmt.Printf("message: %v\n", message)
	fmt.Printf("number: %v\n", number)
	fmt.Printf("url: %v\n", url)
	fmt.Printf("verbose: %v\n", verbose)
	fmt.Printf("trace: %v\n", trace)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", z)

	// This will rearrange the leftover flags back into a slices
	b := argexp.UnMarshall(&a)
	fmt.Printf("\nleftover arguments:\n")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%q\n", b[i])
	}

	// Output:
	// parsed arguments:
	// message: this is (") an escape string
	// number: 123456
	// url: https://google.com
	// verbose: true
	// trace: true
	// x: true
	// y: true
	// z: true
	//
	// leftover arguments:
	// "-a"
	// "-b"
	// "-c"
	// "emacs rocks!!"
}
