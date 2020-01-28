package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// インタフェースの定義
type Subjects interface {
	GetSum() int
	GetAve() int
	GetMedian() int
}

// 構造体の定義
type ElementarySubjects struct {
	Japanese int
	Math     int
	English  int
}

func (e ElementarySubjects) GetSum() int {
	return e.Japanese + e.Math + e.English
}

func (e ElementarySubjects) GetAve() int {
	return e.GetSum() / 3
}

func (e ElementarySubjects) GetMedian() int {
	fmt.Println("小学生のテストは３教科です")
	fmt.Printf("国語：%d 数学：%d 英語：%d \n", e.Japanese, e.Math, e.English)
	array := []int{e.Japanese, e.Math, e.English}
	sort.IntSlice(array).Sort()
	return array[1]
}

type JuniorSubjects struct {
	Japanese int
	Math     int
	Science  int
	Social   int
	English  int
}

func (j JuniorSubjects) GetSum() int {
	return j.Japanese + j.Math + j.Science + j.Social + j.English
}

func (j JuniorSubjects) GetAve() int {
	return j.GetSum() / 5
}

func (j JuniorSubjects) GetMedian() int {
	fmt.Println("中学生のテストは５教科です")
	fmt.Printf("国語：%d 数学：%d 理科：%d  社会：%d 英語：%d \n", j.Japanese, j.Math, j.Science, j.Social, j.English)
	array := []int{j.Japanese, j.Math, j.Science, j.Social, j.English}
	sort.IntSlice(array).Sort()
	return array[2]
}

type HighSubjects struct {
	Japanese  int
	Math      int
	Chemistry int
	Pyshics   int
	Geography int
	History   int
	English   int
}

func (h HighSubjects) GetSum() int {
	return h.Japanese + h.Math + h.Chemistry + h.Pyshics + h.Geography + h.History + h.English
}

func (h HighSubjects) GetAve() int {
	return h.GetSum() / 7
}

func (h HighSubjects) GetMedian() int {
	fmt.Println("高校生のテストは７教科です")
	fmt.Printf("国語：%d 数学：%d 化学：%d  物理：%d 地理：%d 歴史：%d 英語：%d \n", h.Japanese, h.Math, h.Chemistry, h.Pyshics, h.Geography, h.History, h.English)
	array := []int{h.Japanese, h.Math, h.Chemistry, h.Pyshics, h.Geography, h.History, h.English}
	sort.IntSlice(array).Sort()
	return array[3]
}

func ShowTest(s Subjects, name string) {
	fmt.Printf("%sさんのテストの合計値は%d、平均値は%d、中央値は%dです。\n", name, s.GetSum(), s.GetAve(), s.GetMedian())
}

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	var res1 Subjects = ElementarySubjects{
		Japanese: rand.Intn(50) + 50,
		Math:     rand.Intn(50) + 50,
		English:  rand.Intn(50) + 50,
	}

	ShowTest(res1, "Yukari")

	var res2 Subjects = JuniorSubjects{
		Japanese: rand.Intn(50) + 50,
		Math:     rand.Intn(50) + 50,
		Science:  rand.Intn(50) + 50,
		Social:   rand.Intn(50) + 50,
		English:  rand.Intn(50) + 50,
	}

	ShowTest(res2, "Mifuyu")

	var res3 Subjects = HighSubjects{
		Japanese:  rand.Intn(150) + 50,
		Math:      rand.Intn(150) + 50,
		Chemistry: rand.Intn(50) + 50,
		Pyshics:   rand.Intn(50) + 50,
		Geography: rand.Intn(50) + 50,
		History:   rand.Intn(50) + 50,
		English:   rand.Intn(150) + 50,
	}

	ShowTest(res3, "Tamaki")

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
