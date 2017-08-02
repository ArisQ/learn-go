package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("ab" < "ba")

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	log.Print()

	f := func() { fmt.Println("Hello world!") }
	f()
	fmt.Println("The type of f:", reflect.ValueOf(f).Kind())

	log.Print()

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(&a[0])
	func(a [5]int) {
		fmt.Println(&a[0])
	}(a)

	s := []int{1, 2}
	fmt.Println(cap(s))

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "China": "Beijing"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
	fmt.Println(unsafe.Sizeof(1))

	o := car{"BMW", 666}
	fmt.Println(o)

	fmt.Println(reflect.TypeOf(o))

	interfaceTest(1)
	interfaceTest("BMW")

	jsonData := jsonTest{"1", "2", 3, jsonSubField{"sf1", 5}}
	js, jsErr := json.Marshal(jsonData)
	if jsErr != nil {
		fmt.Println("Marshal json error:", jsErr)
	}
	fmt.Printf("%s\n", js)
	fmt.Println(js)
	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(jsonData)
	if err != nil {
		fmt.Println("Error in encoding json")
	}

	hasher := sha1.New()
	sha1Str := "ThisIsASha1TestString"
	io.WriteString(hasher, sha1Str)
	fmt.Printf("The sha1 sum of %s is %x\n", sha1Str, hasher.Sum([]byte{}))
	fmt.Printf("The sha1 sum of %s is %d\n", sha1Str, hasher.Sum([]byte{}))

	notepad := exec.Command("notepad.exe")
	notepad.Run()

	out := make(chan int)
	go func(in chan int) {
		fmt.Println(<-in)
	}(out)
	out <- 2

	go func() {
		fmt.Println("Entering lamda function")
		time.Sleep(2e9)
		fmt.Println("Lambda function finished")
		out <- 0
	}()
	fmt.Println("Wating for lambda function")
	<-out
	{
	}
}

type jsonSubField struct {
	Subfield1 string
	subfiled2 int
}
type jsonTest struct {
	Field1 string
	Field2 string
	field3 int
	Field4 jsonSubField
}

func interfaceTest(d interface{}) {
	switch d.(type) {
	case int:
		fmt.Println("interface implements int, value", d)
	case string:
		fmt.Println("interface implements string, value", d)
	default:
		fmt.Println("out of detection")
	}
}

type car struct {
	make  string //关键字make可以用作标识符
	price float32
}
