// [_Command-line flags_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// are a common way to specify options for command-line programs

package main

import (
	"flag"
	"fmt"
)

func main() {

	
	inFilePtr := flag.String("f", "INPUT.TXT", "Plik wejściowy skond ma braść")

	numbPtr := flag.Int("t", 10, "Timeout zapytania")

	flag.Parse()


	fmt.Println("inputFile:", *inFilePtr)
	fmt.Println("numb:", *numbPtr)

}
