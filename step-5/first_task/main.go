package main

import (
    "bytes"
    "context"
    "io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
    if err := ctx.Err(); err != nil {
        return false, err
    }

    if len(seq) == 0 {
        return true, nil
    }

    buf := make([]byte, 4096)
    var tail []byte

    for {
        if err := ctx.Err(); err != nil {
            return false, err
        }

        n, err := r.Read(buf)

        if ctx.Err() != nil {
            return false, ctx.Err()
        }

        if err != nil && err != io.EOF {
            return false, err
        }

        data := append(tail, buf[:n]...)

        if bytes.Contains(data, seq) {
            return true, nil
        }

        keep := len(seq) - 1
        if keep > len(data) {
            keep = len(data)
        }

        tail = append([]byte(nil), data[len(data)-keep:]...)

        if err == io.EOF {
            return false, nil
        }
    }
}