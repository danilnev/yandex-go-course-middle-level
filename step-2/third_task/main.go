package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = inputFile.Seek(int64(startpos), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return outputFile.Close()
}
