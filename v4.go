package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)
type Reader interface {

	Read(p []byte) (n int, err error)

}
type Writer interface {

	Write(p []byte) (n int, err error)

}

func main() {
	f, err := os.Create("proverb.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	proverbs := []string{

		"don't communicate by sharing memory share memory by communicating",
	}

	var writer bytes.Buffer


	for _, p := range proverbs {

		n, err := writer.Write([]byte(p))

		if err != nil {

			fmt.Println(err)

			os.Exit(1)

		}

		if n != len(p) {

			fmt.Println("failed to write data")

			os.Exit(1)

		}

	}


	fmt.Println(writer.String())

	reader := strings.NewReader("don't communicate by sharing memory share memory by communicating")

	p := make([]byte, 65)


	for {

		n, err := reader.Read(p)

		if err != nil{

			if err == io.EOF {

				fmt.Println("EOF:", n)

				break

			}

			fmt.Println(err)

			os.Exit(1)

		}

		fmt.Println(n, string(p[:n]))

	}




}



