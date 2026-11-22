package main

import (
	"bytes"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {

	buf := make([]byte, 1024)
	var data []byte
	tailSize := len(seq) - 1

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return false, err
		}

		if len(seq) == 0 {
			return true, nil
		}

		if n > 0 {
			data = append(data, buf[:n]...)

			if bytes.Contains(data, seq) {
				return true, nil
			}

			if len(data) > tailSize {
				copy(data, data[len(data)-tailSize:])
				data = data[:tailSize]
			}
		}

		if err == io.EOF {
			return false, nil
		}
	}
}
