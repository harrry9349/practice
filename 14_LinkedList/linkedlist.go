package main

import (
	"fmt"
	"math/rand"
	"time"
)

// セル
type Cell struct {
	item int
	next *Cell
}

// セルの検索
func (cp *Cell) searchCell(n int) *Cell {
	i := -1
	for cp != nil {
		if i == n {
			return cp
		}
		i++
		cp = cp.next
	}
	return nil
}

// リスト
type List struct {
	top *Cell
}

// リストの検索
func (lst *List) searchList(n int) (int, bool) {
	cp := lst.top.searchCell(n)
	if cp == nil {
		return 0, false
	}
	return cp.item, true
}

// リストの挿入
func (lst *List) insertList(n, x int) bool {
	cp := lst.top.searchCell(n - 1)
	if cp == nil {
		return false
	}
	// n-1番目のリストが新規作成したn番目のリストを参照するようになる
	cp.next = newCell(x, cp.next)
	return true
}

// リストの削除
func (lst *List) deleteList(n int) bool {
	cp := lst.top.searchCell(n - 1)
	if cp == nil || cp.next == nil {
		return false
	}
	// n-1番目のリストがn+1番目のリストを参照するようになる
	cp.next = cp.next.next
	return true
}

// リストの空判定
func (lst *List) isEmpty() bool {
	return lst.top.next == nil
}

// リストの表示
func (lst *List) printList() {
	cp := lst.top.next
	for ; cp != nil; cp = cp.next {
		fmt.Print(cp.item, " ")
	}
	fmt.Println("")
}

// セルの生成
func newCell(x int, cp *Cell) *Cell {
	newcp := new(Cell)
	newcp.item, newcp.next = x, cp
	return newcp
}

// リストの生成
func newList() *List {
	lst := new(List)
	lst.top = new(Cell)
	return lst
}

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	a := newList()
	for i := 0; i < 4; i++ {
		fmt.Println(a.insertList(i, i))
	}
	a.printList()
	for i := 0; i < 5; i++ {
		n, ok := a.searchList(i)
		fmt.Println(n, ok)
	}
	for !a.isEmpty() {
		a.deleteList(0)
		a.printList()
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
