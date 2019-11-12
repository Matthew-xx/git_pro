package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)  //因为函数会改变里面的值，所以取地址

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	} // n行 n 列表格

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}  //上左下右四种走动

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}  //步伐加1（表示位置
}

func (p point) at(grid [][]int) (int, bool) { //bool表示有没有值(走出去)
	if p.i < 0 || p.i >= len(grid) {  //i行越界（超过了迷宫范围
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {  //j列越界（超过了迷宫范围
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int,
	start, end point) [][]int {
	steps := make([][]int, len(maze))  //步骤
	for i := range steps {
		steps[i] = make([]int, len(maze[i])) //maze[i]有多少行，steps也有多少行
	}

	Q := []point{start} //队列，用来添加走过的点，初始值是起点

	for len(Q) > 0 { //队列为空时退出（就是没有新的点添加进来，死路
		cur := Q[0]  //当前节点
		Q = Q[1:]    //走过后就截掉

		if cur == end {
			break
		}  //到终点后，退出

		for _, dir := range dirs {
			next := cur.add(dir)  //下个节点为当前节点加上上下左右其中一个方向

			val, ok := next.at(maze)
			if !ok || val == 1 { //没有值或者值为1（撞墙)
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {  //不等于0，说明已经走过
				continue
			}

			if next == start {  //回到原点
				continue
			}
			//以上三种情况是不继续探索的，

			curSteps, _ := cur.at(steps)  //当前步骤数
			steps[next.i][next.j] =
				curSteps + 1    //下一步

			Q = append(Q, next)  //加到队列
		}
	}

	return steps
}

func main() {
	maze := readMaze("./maze.in")

	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})  //起点，终点

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)  //3d表示3位对齐，因为步骤最多是两位数，所以用3位对齐，打印出来整齐
		}
		fmt.Println()
	}

	// TODO: construct path from steps
}
