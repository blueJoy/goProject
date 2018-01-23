package main

import (
	"crypto/md5"
	"fmt"
	"crypto/sha1"
	"os"
	"io"
	"math"
)

func main(){
	TestString := "Hi,pandaman!"

	//MD5加密
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\r\n",Result)

	//SHA1加密
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\r\n",Result)

	//对文件进行加密

	TestFile := "123.txt"
	infile,err := os.Open(TestFile)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

	md5h := md5.New()
	io.Copy(md5h,infile)
	fmt.Printf("%x %s\r\n",md5h.Sum([]byte("")),TestFile)

	sha1h := sha1.New()
	io.Copy(sha1h,infile)
	fmt.Printf("%x %s\r\n",sha1h.Sum([]byte("")),TestFile)
}
