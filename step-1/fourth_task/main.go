package main

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	buf := make([]byte, n)

	for n > 0 {
		if uint(len(buf)) > n {
			buf = buf[:int(n)]
		}

		count, readErr := r.Read(buf)

		if count > 0 {
			written, writeErr := w.Write(buf[:count])

			if writeErr != nil {
				return writeErr
			}

			if written != count {
				return io.ErrShortWrite
			}

			n -= uint(written)
		}

		if readErr == io.EOF {
			return nil
		}

		if readErr != nil {
			return readErr
		}

		if count == 0 {
			return io.ErrNoProgress
		}
	}

	return nil
}
