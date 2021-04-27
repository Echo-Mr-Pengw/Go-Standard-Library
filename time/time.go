package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Local(), t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Unix(), t.UnixNano(),t.Format("2006-01-02 15:04:05"))

	t1, err := time.Parse("2006-01-02 15:04:05", "2020-01-01 23:01:01")
	fmt.Println(t1.Unix(), err, time.Local, time.Hour)

	t2, _ := time.Parse("2006年01月02日", "2020年03月01日")
	fmt.Println(t2)
}
