package main

import (
	"fmt"
)

type table struct {
	row  int
	line int
	data [][]bool

	usedX map[int]bool
}

func createTable(row, line int) *table {
	t := &table{
		row:   row,
		line:  line,
		usedX: make(map[int]bool),
	}
	t.clear()

	return t
}

func (t *table) clear() {
	t.usedX = make(map[int]bool)
	t.data = make([][]bool, t.line)
	for i := range t.data {
		t.data[i] = make([]bool, t.row)
	}
}

func (t *table) set(x, y int) {
	t.usedX[x] = true
	t.data[y][x] = true
}

func (t *table) unset(x, y int) {
	t.usedX[x] = false
	t.data[y][x] = false
}

func (t *table) print() {
	for y := range t.data {
		for i := 0; i < t.row; i++ {
			if t.data[y][i] {
				fmt.Print("●")
			} else {
				fmt.Print("○")
			}
			print("\t")
		}
		fmt.Print("\n")
	}
}

func (t *table) checkOK(x, y int) bool {
	if t.data[y][x] {
		return false
	}
	
	// 斜め ＼
	startX := x
	startY := y
	for {
		if startX-1 < 0 || startY-1 < 0 {
			break
		}
		startX--
		startY--
	}
	for {
		if t.data[startY][startX] {
			return false
		}
		startX++
		startY++
		if startX >= t.row || startY >= t.line {
			break
		}
	}

	// 斜め ／
	startX = x
	startY = y
	for {
		if startX+1 >= t.row || startY-1 < 0 {
			break
		}
		startX++
		startY--
	}
	for {
		//		fmt.Printf("／　x:%d y:%d, sx:%d sy:%d\n", x, y, startX, startY)
		if t.data[startY][startX] {
			return false
		}
		startX--
		startY++
		if startX < 0 || startY >= t.line {
			break
		}
	}

	return true
}

func calculate(tbl *table, y int) bool {
	nextX := 0
	for nextX = 0; nextX < tbl.row; nextX++ {
		if !tbl.usedX[nextX] && tbl.checkOK(nextX, y) {
			//			fmt.Printf("Y=%d: X=%d\n", y, nextX)
			tbl.set(nextX, y)
			//			tbl.print()
			if y+1 >= tbl.line {
				return true
			} else {
				if calculate(tbl, y+1) {
					return true
				}
			}
			//fmt.Printf("x=%d, y=%d, false\n", nextX, y)
			tbl.unset(nextX, y)
		}
	}

	//	fmt.Printf("y=%d, false\n", y)

	return false
}

func main() {
	tbl := createTable(8, 8)
	//rand.Seed(time.Now().Unix())
	//x := rand.Intn(tbl.row)

	for i := 0; i < tbl.row; i++ {
		tbl.set(i, 0)
		ret := calculate(tbl, 1)
		println("ANSWER")
		println(ret)
		tbl.print()
		tbl.clear()
	}
}
