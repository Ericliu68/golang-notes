package main

import (
	"os"
	"log"
	"io/ioutil"
	//"bufio"
)
func main() {
	file, err := os.OpenFile("./main.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Println("open filed")
		return
	}
	defer file.Close()

	//str := "hello 沙河"
	//file.Write([]byte(str))       //写入字节切片数据
	//file.WriteString("hello 小王子\n") //直接写入字符串数据

	//var tmp = make([]byte, 128)
	//n, err := file.Read(tmp)
	//
	//
	//if err != nil {
	//	log.Println(err)
	//	if err == io.EOF {
	//		log.Println("文件读写完毕")
	//
	//	} else {
	//		return
	//	}
	//
	//}
	//log.Println(n)
	//log.Println(string(tmp[:n-1]))

	//	var content []byte
	//	var tmp = make([]byte, 128)
	//	for {
	//		n, err:= file.Read(tmp)
	//		if err == io.EOF{
	//			log.Println("文件读写完毕")
	//			break
	//		}
	//		if err != nil {
	//			log.Println(err)
	//			return
	//		}
	//		log.Println(string(tmp[:n-1]))
	//		content = append(content, tmp[:n-1]...)
	//
	//
	//	}
	//	log.Println(string(content))
	//}


	//reader := bufio.NewReader(file)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		log.Println(line)
	//		if len(line) != 0 {
	//
	//			log.Println(line)
	//		}
	//		break
	//	}
	//}
	content, err := ioutil.ReadFile("./main.log")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(content))

}