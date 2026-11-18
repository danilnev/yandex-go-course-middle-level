package main

import (
	"fmt"
	"io"
)

func WriteString(s string, w io.Writer) error {
	writtenBytes, err := w.Write([]byte(s))
	if err == nil {
		fmt.Printf("Записано %d байтов", writtenBytes)
		return nil
	}
	return err
}
