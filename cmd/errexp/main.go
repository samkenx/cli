package main

import (
	"errors"
	"fmt"

	"github.com/ActiveState/cli/internal/errs"
)

// ...
var (
	ErrSentinel     = errors.New("bingo, bango")
	ErrOtherExample = errors.New("some external happening")
	ErrOuter        = errors.New("outer wrap")
)

// ComplexError ...
type ComplexError struct {
	Cause error
	Data  string
}

func (e *ComplexError) Unwrap() error {
	return e.Cause
}

func (e *ComplexError) Error() string {
	return fmt.Sprintf("data %q: %v", e.Data, e.Cause)
}

func printfln(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Printf("\n")
}

func main() {
	var err error

	printfln("WrapErrors (first error arg is not detectable)")
	err = ErrSentinel
	err = errs.WrapErrors(err, ErrOuter)

	if errors.Is(err, ErrSentinel) {
		printfln("\tRightly detected as a sentinel error")
	} else {
		printfln("\tShould have been detected as a sentinel error")
	}

	if errors.Is(err, ErrOuter) {
		printfln("\tRightly detected as an outer error")
	} else {
		printfln("\tShould have been detected as an outer error")
	}

	//

	printfln("ComplexErrors (message content duplication in existing complex errors)")
	err = &ComplexError{
		Cause: ErrOtherExample,
		Data:  "this is my complex error...",
	}

	if errors.Is(err, ErrOtherExample) {
		printfln("\tRightly detected as an other example sentinel error")
	} else {
		printfln("\tShould have been detected as an other example sentinel error")
	}

	if cerr := new(*ComplexError); errors.As(err, cerr) {
		printfln("\tRightly inspected as a complex error")

		printfln("\tdirect error output from top of chain = %v", *cerr)
		printfln("\terrs.Join output of the entire chain  = %v", errs.Join(err, ": "))

		printfln("\t\tWhile using errs.Join, complex errors cannot be formed as expected (i.e. Error() return vals must not include wrapped error message)")
		printfln("\t\tComplex errors that are externally managed and that follow expectations will need to be removed from chains or they will double print")
		printfln("\t\tIt may be impossible to retain inspectable errors and accommodate errs.Join")
	} else {
		printfln("Should have been inspected as a complex error")
	}
}
