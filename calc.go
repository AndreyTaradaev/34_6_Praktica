package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calculate (a,b,op string ) string {
	var nret int
	na,err := strconv.Atoi(a)
	if(err != nil){
		fmt.Println(err)
	 os.Exit(-10)
	}
	nb,err := strconv.Atoi(b)
	if(err != nil){
		fmt.Println(err)
	 os.Exit(-11)
	}
	switch op {
	case  "-":
		nret = na-nb
	case  "+":
		nret = na+nb
	case  "*":
		nret = na*nb
	case  "/":
		if nb ==0 {
			fmt.Println("divide by zero")
			os.Exit(-12)
		}
		nret = na/nb
	default : 	
	fmt.Println("unknown operation")
	os.Exit(-13)
	}

return  fmt.Sprintf("%d%s%d=%d", na,op,nb,nret)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите файл содержащий математические операции: ")
	inputstr, err := reader.ReadString('\n')
	if err != nil {
	    fmt.Println(err)
	 os.Exit(-1)
	}
	inputstr = strings.TrimSpace(inputstr)
	fmt.Print("Введите выходной файл: ")
	outstr, err := reader.ReadString('\n')
	if err != nil {
	    fmt.Println(err)
	 os.Exit(-2)
	}
	outstr = strings.TrimSpace(outstr)
f, err := os.OpenFile(inputstr, os.O_RDONLY, 0777)  
  if err != nil {
     fmt.Println(err)
	 os.Exit(-3)
  }    
  defer f.Close()
  out , err := os.OpenFile(outstr, os.O_WRONLY|os.O_CREATE, os.ModePerm)

  if(err!= nil){
	  fmt.Println(err)
	  os.Exit(-4)	  
  }  
 defer out.Close()

  writer := bufio.NewWriter(out)

  fileReader := bufio.NewReader(f)
  re := regexp.MustCompile(`([0-9]+)([-\+\*\/])([0-9]+)=[?]$`)
  for {
	line, _, err := fileReader.ReadLine()
	if err == io.EOF {
	   break
	}
	if re.MatchString(string(line)){		
		submatches := re.FindAllStringSubmatch(string(line), -1)	
	 str:= calculate( submatches[0][1],submatches[0][3],submatches[0][2])
	 writer.WriteString( str+"\n")
	}
  }
  writer.Flush()
}