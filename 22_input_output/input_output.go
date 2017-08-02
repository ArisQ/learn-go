package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Print("Current command:")
	fmt.Println(strings.Join(os.Args, " "))
	var fileName string
	fmt.Print("Please input filename:")
	fmt.Scanln(&fileName)
	for fileName == "" {
		fmt.Print("File name cannot be empty, please input again:")
		fmt.Scanln(&fileName)
	}
	fmt.Println("File name is", fileName)

	//GoRun will run file in a temporary directory like
	//C:\Users\Q\AppData\Local\Temp\go-build558459930\command-line-arguments\
	//_obj\exe\input_output.exe
	//So we change the directory to current
	//fileName = "E:\\GitHub\\learn-go\\22_input_output\\" + fileName
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		fmt.Println("Error:", fileErr)
		return
	}
	defer file.Close()

	// display file content
	fmt.Println("File content:")
	reader := bufio.NewReader(file)
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		_ = str
		fmt.Print(">", str)
	}

	// write file
	fmt.Println("Writing to", fileName+".backup")
	outFile, outErr := os.OpenFile(fileName+".backup", os.O_WRONLY|os.O_CREATE, 0666)
	if outErr != nil {
		fmt.Print("Create write file error:", outErr)
		return
	}

	if _, seekErr := file.Seek(0, 0); seekErr != nil {
		fmt.Println("Seek file error")
		return
	}
	writer := bufio.NewWriter(outFile)
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		writer.WriteString(str)
	}
	writer.Flush()
	outFile.Close()
	fmt.Println("Write to", fileName+".backup", "complete.")

	// compress
	fmt.Println("Writing to compress file", fileName+".gz")
	gzFile, gzErr := os.OpenFile(fileName+".gz", os.O_WRONLY|os.O_CREATE, 0666)
	if gzErr != nil {
		fmt.Print("Create gzip file error:", outErr)
		return
	}

	if _, seekErr := file.Seek(0, 0); seekErr != nil {
		fmt.Println("Seek file error")
		return
	}
	gzWriter := gzip.NewWriter(gzFile)
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		gzWriter.Write([]byte(str))
	}
	gzWriter.Flush()
	gzWriter.Close()
	gzFile.Close()
	fmt.Println("Write to", fileName+".gz", "complete.")

	// copy
	fmt.Println("Copying to cpy file", fileName+".cpy")
	cpyFile, cpyErr := os.OpenFile(fileName+".cpy", os.O_WRONLY|os.O_CREATE, 0666)
	if cpyErr != nil {
		fmt.Print("Create cpy file error:", cpyErr)
		return
	}

	if _, seekErr := file.Seek(0, 0); seekErr != nil {
		fmt.Println("Seek file error")
		return
	}
	io.Copy(cpyFile, file)
	cpyFile.Close()
	fmt.Println("Copy to", fileName+".cpy", "complete.")
}
