package main

import (
	"io"
	"os"
	"strings"
    "fmt"
)

type rot13Reader struct {
	r io.Reader
}
func (s *rot13Reader) Read(b []byte) (n int, err error) {
    res, er := s.r.Read(b)
    for i, char := range(b) {
        if char <= 'Z' && char >='A'  {
            b[i] = (char - 'A' + 13)%26 + 'A'
        }else if char >= 'a' && char <= 'z' {
            b[i] = (char - 'a' + 13)%26 + 'a'
        }
    }
    return res, er // I don't use this result directly but interfaces need it.
}

type I interface{}

func describe(i I) {
    fmt.Printf("Value = %v, Type = %T \n", i, i)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
    /* "s" is a *Reader (pointer to Reader struct)
      A "Reader" implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
      io.ByteScanner, and io.RuneScanner interfaces by reading from a string. */
	  if i := range ch {
		  continue
	  }
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
