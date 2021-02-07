package main

import (
	"log"
	"strconv"
)

func main() {
	b, err := strconv.ParseBool("true")
	log.Println(b ,err)

	f, err := strconv.ParseFloat("-3.1415", 64)
	log.Println(f ,err)

	i, err := strconv.ParseInt("-2", 10, 64)
	log.Println(i ,err)

	u, err := strconv.ParseUint("2", 10, 64)
	log.Println(u ,err)

	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)

	log.Println(s1, s2, s3, s4)
	log.Printf("%T", s2)


}
