package main

import (
	"bufio"
	"errors"
	"fmt"
	"image/png"
	"io"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func parseReader(r io.Reader) ([]byte, error) {
	reader := bufio.NewReader(r)
	l, _, err := reader.ReadLine()
	return l, err
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	info, err := os.Stdin.Stat()
	checkErr(err)

	var b []byte
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		if len(os.Args) < 2 {
			checkErr(errors.New("No data specified"))
		}
		b = []byte(os.Args[1])
	} else {
		reader := bufio.NewReader(os.Stdin)
		b, _, err = reader.ReadLine()
	}
	checkErr(err)

	// Create the barcode
	qrCode, err := qr.Encode(string(b), qr.M, qr.Auto)
	checkErr(err)

	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	checkErr(err)
	png.Encode(os.Stdout, qrCode)
}
