package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	     
    //   // open file
	//   f, err := os.Open("./example.txt")
      
	//   if err != nil {
	// 	  panic(err)
	//   }
		
		
	//   fileInfo, err := f.Stat()
		
	//   if err != nil {
	// 	  panic(err)
	//   }
		
		
	//   // file info
	//   fmt.Println("file Name " , fileInfo.Size())
	//   fmt.Println("file Name " , fileInfo.Name())
	//   fmt.Println("file Name " , fileInfo.IsDir())
	//   fmt.Println("file Name " , fileInfo.ModTime())
	//   fmt.Println("file Name " , fileInfo.Mode())
		
	//   // read file data
	//   f, err := os.Open("./example.txt")
	//   if err != nil {
	// 	  panic(err)
	//   }
	//   defer f.Close()
		
	//   fileInfo, err := f.Stat()
		
	//   if err != nil {
	// 	  panic(err)
	//   }
		
	//   data := make([]byte, fileInfo.Size())
	//   _, err = f.Read(data)
	//   if err != nil {
	// 	  panic(err)
	//   }
		
	//   for i := 0; i < len(data); i++ {
	// 	  fmt.Println("data", string(data[i:i+1]))
	// 	  // data h
	// 	  // data e
	// 	  // data l
	// 	  // data l
	// 	  // data o
	//   }
  

	  	// read file
		  f, err := os.ReadFile("./example.txt")
     
		  if err != nil {
			  panic(err)
		  }
		   
		  fmt.Println(string(f))
		   
		  // read folder
		  dir, err := os.Open(".")
		   
		  if err != nil {
			  panic(err)
		  }
		   
		  defer dir.Close()
		   
		   
		  files, err := dir.Readdir(-1)
		   
		  if err != nil {
			  panic(err)
		  }
		   
		  for _, file := range files {
			  fmt.Println("file Name ", file.Name(), file.IsDir())
		  }
		   
		   
		  // write file
		  file, err := os.Create("./example2.txt")
		   
		  if err != nil {
			  panic(err)
		  }
		   
		  defer file.Close()
		   
		  // write file
		   
		  _, err = file.WriteString("hello sandip\n")
		  if err != nil {
			  panic(err)
		  }
		   
		  // replace hello sandip to hello uday
		  _, err = file.WriteAt([]byte("hello uday\n"), 0)
		  if err != nil {
			  panic(err)
		  }
		   
		  bytes := []byte("hello sandip\n")
		  _, err = file.Write(bytes)
		  if err != nil {
			  panic(err)
		  }
		   
		   
		  // read and write to another file (streaming fashion)
		   
		  sourceFile, err := os.Open("./example3.txt")
		  if err != nil {
			  panic(err)
		  }
		   
		  defer sourceFile.Close()
		   
		  destinationFile, err := os.Create("./example4.txt")
		  if err != nil {
			  panic(err)
		  }
		   
		  defer destinationFile.Close()
		   
		   
		  // copy
		  _, err = io.Copy(destinationFile, sourceFile)
		  if err != nil {
			  panic(err)
		  }
		   
		  // copy with buffer
		  reader := bufio.NewReader(sourceFile)
		  writer := bufio.NewWriter(destinationFile)
		   
		  for {
			  str, err := reader.ReadString('\n')
			  if err == io.EOF {
				  break
			  }
			  if err != nil {
				  panic(err)
			  }
			  writer.WriteString(str)
		  }
		   
		  writer.Flush()
		   
		  // delete file 
		  err = os.Remove("./example3.txt")
		  if err != nil {
			  panic(err)
		  }
}