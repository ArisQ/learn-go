package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(Append([]byte{1, 2, 3}, []byte{6, 7, 8}))

	fmt.Println(len([]byte{1, 2, 3}[1:1]))
	fmt.Println(len([]byte{1, 2, 3}[1:2]))

	fmt.Println()
	slice2 := []byte{1, 2, 3, 4, 5, 6}
	s2 := slice2[1:]
	fmt.Println("Array before append:", slice2, s2)
	slice2 = Append(slice2[:3], []byte{6, 7, 8})
	fmt.Println("Array after append:", slice2, s2)

	fmt.Println()
	slice3 := []byte{2, 3, 4, 5, 6, 7}
	s3 := slice3[1:]
	fmt.Println("Array before append:", slice3, s3)
	slice3 = append(slice3[:3], []byte{6, 7, 8}...)
	fmt.Println("Array after append:", slice3, s3)

	fmt.Println("\nWill panic:")
	fmt.Println(AppendOutOfRange([]byte{1, 2, 3}, []byte{6, 7, 8}))
}

func AppendOutOfRange(slice, data []byte) []byte {
	for i, j := len(slice), 0; j < len(data); {
		slice[i] = data[j]
		i++
		j++
	}
	return slice
}

func Append(slice, data []byte) []byte {
	// This implement is really ugly
	log.Println(slice)
	log.Println(data)
	var s []byte
	if cap(slice) < len(slice)+len(data) {
		s = make([]byte, len(slice)+len(data))
		for i := 0; i < len(slice); i++ {
			s[i] = slice[i]
		}
	} else {
		s = slice[:len(slice)+len(data)] //Append will change other slice which shares the same array in this way!!
	}
	for i, j := len(slice), 0; j < len(data); i, j = i+1, j+1 {
		s[i] = data[j]
	}
	return s
}
