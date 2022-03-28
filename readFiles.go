package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//READING FILES
	plik, err := os.ReadFile("test.txt")
	check(err)
	fmt.Print(string(plik))

	//TO DZIAŁA ALE JEST ZAKOMENTOWANE ŻEBY NIE BRUDZIĆ OUTPUTU PRZY TESTACH, do używania w przypadku wykonywania w pętli
	// fmt.Println("Druga metoda")
	// plik2, err2 := os.Open("test.txt")
	// check(err2)
	// scanner := bufio.NewScanner(plik2)
	// iteracja := 1
	// for scanner.Scan() {
	// 	fmt.Println(string(iteracja) + scanner.Text())
	// 	iteracja++
	// }

	// plik2.Close()

	//WRITING TO FILES

	plik2, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0)
	defer plik2.Close()
	zapis, err := plik2.WriteString("znowu ktos tu był\n")
	check(err)
	zapis = zapis
	plik2.Sync()

}
