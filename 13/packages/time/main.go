package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println(time.Now())

	t := time.Date(2015, 7, 19, 19, 14, 23, 0, time.Local)
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.YearDay())
	fmt.Println(t.Month())
	fmt.Println(t.Month().String())
	fmt.Println((int)(t.Month()))
	fmt.Println(t.Weekday())
	fmt.Println(t.Weekday().String())
	fmt.Println((int)(t.Weekday()))
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Zone())

	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	d, _ := time.ParseDuration("2h30m")
	t1 := time.Now()
	fmt.Println(t1.Add(d))

	t2 := time.Date(2020, 7, 24, 0, 0, 0, 0, time.Local)
	t3 := time.Now()
	fmt.Println(t2.Sub(t3))

	t4 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.Local)
	t5 := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(t5.Before(t4))
	fmt.Println(t5.After(t4))
	fmt.Println(t4.Before(t5))
	fmt.Println(t4.After(t5))
	t6 := time.Now()
	fmt.Println(t6.Equal(t6))
	t7 := time.Date(2015, 10, 1, 9, 0, 0, 0, time.Local) // JST
	t8 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.UTC)   // UTC
	fmt.Println(t7.Equal(t8))

	t9 := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(t9.AddDate(1, 0, 0))
	fmt.Println(t9.AddDate(0, -1, 0))
	fmt.Println(t9.AddDate(0, 0, -1))
	fmt.Println(t9.AddDate(0, 2, -12))

	t10, err := time.Parse("2006/01/02", "2015/11/27")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t10)

	t11 := time.Date(2015, 11, 23, 13, 11, 23, 12345, time.Local)
	fmt.Println(t11.Format("2006年1月2日 15時04分05秒"))
	fmt.Println(t11.Format("2006/01/02"))

	t12 := time.Date(2015, 11, 23, 13, 11, 23, 12345, time.Local)
	fmt.Println(t12)
	fmt.Println(t12.UTC())

	t13 := time.Date(2015, 11, 23, 13, 11, 23, 12345, time.UTC)
	fmt.Println(t13)
	fmt.Println(t13.Local())

	t14 := time.Date(2015, 11, 23, 13, 11, 23, 12345, time.UTC)
	fmt.Println(t14.Unix())

	time.Sleep(1 * time.Second)
}
