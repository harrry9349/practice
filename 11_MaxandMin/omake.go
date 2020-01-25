package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 構造体の定義
type Subjects struct {
	Japanese int
	Math     int
	Science  int
	Social   int
	English  int
}

func (s Subjects) ShowSum() int {
	return s.Japanese + s.Math + s.Science + s.Social + s.English
}

func (s Subjects) ShowAve() int {
	return s.ShowSum() / 5
}

type Student struct {
	Num  int
	Name string
	Subjects
}

func Cheat(s Subjects) {
	s = Subjects{500, 500, 500, 500, 500}
	fmt.Println("Cheat関数内における値", s)
}

func CheatPoint(s *Subjects) {
	*s = Subjects{500, 500, 500, 500, 500}
	fmt.Println("CheatPoint関数内における値", *s)
}

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	res := Student{
		Num:  1,
		Name: "Taro",
		Subjects: Subjects{
			Japanese: rand.Intn(50),
			Math:     rand.Intn(50),
			Science:  rand.Intn(50),
			Social:   rand.Intn(50),
			English:  rand.Intn(50),
		},
	}

	fmt.Printf("番号%dの%s君の成績・・・ 国語：%d 数学：%d 理科：%d  社会：%d 英語：%d 合計：%d 平均：%d \n", res.Num,
		res.Name, res.Subjects.Japanese, res.Subjects.Math, res.Subjects.Science, res.Subjects.Social, res.Subjects.English, res.Subjects.ShowSum(), res.Subjects.ShowAve())

	fmt.Printf("%s君はこんな成績を親には見せられないと思い、改ざんを行うことにした。\n", res.Name)

	res2 := res
	fmt.Println("res2 := res")
	res2.Subjects = Subjects{100, 100, 100, 100, 100}
	fmt.Printf("改ざんを実行した後の%s君の成績・・・ 国語：%d 数学：%d 理科：%d  社会：%d 英語：%d 合計：%d 平均：%d \n",
		res.Name, res.Subjects.Japanese, res.Subjects.Math, res.Subjects.Science, res.Subjects.Social, res.Subjects.English, res.Subjects.ShowSum(), res.Subjects.ShowAve())

	fmt.Println("「変わっていない！　なぜ？」")
	fmt.Println("res2 := resは、Student型の値をコピーしてres2に格納しているので、res2で書き換えを行っても、それはresに反映されない。（値渡し）")
	fmt.Println("ポインタを指定するとどうなる？")

	res3 := &res
	fmt.Println("res3 := &res")
	res3.Subjects = Subjects{100, 100, 100, 100, 100}
	fmt.Printf("改ざんを実行した後の%s君の成績・・・ 国語：%d 数学：%d 理科：%d  社会：%d 英語：%d 合計：%d 平均：%d \n",
		res.Name, res.Subjects.Japanese, res.Subjects.Math, res.Subjects.Science, res.Subjects.Social, res.Subjects.English, res.Subjects.ShowSum(), res.Subjects.ShowAve())

	fmt.Println("「変わった！」")
	fmt.Println("res3 := &res は、*Student型(Studentへのポインタである *Student型)をres3に格納しているので、res3はresのアドレス(Studentへのポインタである *Student型)を持っていることになる。従って、res3で書き換えを行うと、その変更はresに反映される。（参照渡し）")

	fmt.Println("関数を使用して元の構造体の値を変化させるときも引数にポインタを指定する")
	fmt.Println("値渡し：Cheat(res.Subjects)")
	Cheat(res.Subjects)
	fmt.Println("Cheat関数実行後の値", res.Subjects)

	fmt.Println("参照渡し：CheatPoint(&res.Subjects)")
	CheatPoint(&res.Subjects)
	fmt.Println("CheatPoint関数実行後の値", res.Subjects)

	fmt.Println("「よかった、これで親に見せられる！」")
	fmt.Printf("そう安心した%s君だったが、その後改ざんがバレて全教科0点にさせられ、結局大目玉を喰らった。\n", res.Name)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
