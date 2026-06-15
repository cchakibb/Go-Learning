// Écris une struct `RepeatReader` qui implémente `io.Reader`.
// Elle reçoit une string à la construction et la répète
// indéfiniment quand on la lit.

// type RepeatReader struct {
//     data []byte
//     pos  int
// }

// func NewRepeatReader(s string) *RepeatReader
// func (r *RepeatReader) Read(p []byte) (int, error)

// Dans main :
// - Crée un RepeatReader avec "hello "
// - Lis 20 bytes depuis ce reader (buf := make([]byte, 20))
// - Affiche le résultat

// Attendu : "hello hello hello he"

package main

import (
	"fmt"
)

type RepeatReader struct {
    data []byte
    pos  int
}

func NewRepeatReader(s string) *RepeatReader {
	return &RepeatReader{
		data:	[]byte(s),
		pos:	0,
	}
}

func (r *RepeatReader) Read(p []byte) (int, error) {
	fmt.Println("start with r.pos", r.pos)
	for i := 0; i < len(p); i++ {
		p[i] = r.data[r.pos]
		r.pos++
		if r.pos == len(r.data) {
			r.pos = 0
		}
	}
	fmt.Println("finish with r.pos ", r.pos)
	return len(p), nil
}

func main() {
	
	r := NewRepeatReader("hello ")


	buf := make([]byte, 20)
	n, _ := r.Read(buf)
	fmt.Printf("Lu %d bytes: %s\n", n, buf[:n])
	fmt.Println()
	buf = make([]byte, 5)
	n, _ = r.Read(buf)
	fmt.Printf("Lu %d bytes: %s\n", n, buf[:n])

}

