package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(ta *testing.T) {
	//t := time.Now() //当前时间 根据系统时区来
	//fmt.Println(t)
	//t.Unix() //时间戳
	//ts := t.Format("2006-01-02 15:04:05") //time转string
	//fmt.Println(ts)

	originDateTime := time.Unix(1662912020, 0)

	//originDateStr := originDateTime.Format("2006-01-02 00:00:00")
	//fmt.Println(originDateStr) //2022-09-12 00:00:00 丢失精度

	originDateStr := originDateTime.Format("2006-01-02 15:04:05")
	fmt.Println(originDateStr) // 2022-09-12 00:00:20

}
