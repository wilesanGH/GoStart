package main

import (
	"os"
	"fmt"
)

func readMaze(filename string) [][]int{
	file,err := os.Open(filename)
	if err !=nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file,"%d %d",&row,&col)
	maze := make([][]int,row)
	for i:= range maze{
		maze[i] = make([]int,col)
		for j := range maze[i] {
			fmt.Fscanf(file,"%d",&maze[i][j])
		}
	}
	return maze
}

type point struct{
	i,j int
}

var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add (r point) point {
	return point{p.i+r.i,p.j+r.j}
}

func (p point) at(grid [][]int) (int,bool){
	if p.i< 0||p.i>=len(grid) {
		return 0,false
	}

	if p.j<0||p.j>= len(grid[p.i]){
		return 0,false
	}

	return grid[p.i][p.j],true
}

func walk(maze [][]int,start,end point) [][]int{
	steps := make([][]int,len(maze))
	for i := range steps{
		steps[i] = make([]int,len(maze[i]))
	}
	Q := []point{start}

	for len(Q) >0{
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs{
			next := cur.add(dir)

			val,ok := next.at(maze)
			if !ok||val ==1{
				continue
			}

			val,ok = next.at(steps)
			if !ok || val != 0{
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps+1

			Q = append(Q,next)
		}
	}

	return steps
}


func printPath(steps [][]int){
	//pathPoint := make([]point,1)
	path := make([]int,1)
	row,col := len(steps)-1,len(steps[0])-1
	fmt.Println(row,col)
	path[0] = steps[len(steps)-1][len(steps[0])-1]
	cur := path[len(path)-1]
	for i := len(steps)-1;i>=0;i--{
		for j := len(steps[i])-1;j>=0;j--{
			if steps[i][j]!=0 && cur == steps[i][j]+1{
				path = append(path, steps[i][j])
				fmt.Println(i,j)
				cur = path[len(path)-1]
				if(i+2 < len(steps)){
					i=i+2
				}
				break

			}
		}
	}

	fmt.Println(path)
	//
	// fmt.Println(pathPoint)
}

func PrintPath2(steps [][]int,end point)  {
	Q := []point{end}
	fmt.Println(end)
	for len(Q) > 0{
		cur := Q[0]
		Q = Q[1:]
		for _,dir := range dirs{
			front := cur.add(dir)
			val,ok := front.at(steps)
			if !ok||val ==1{
				continue
			}

			if steps[front.i][front.j] + 1 == steps[cur.i][cur.j]{
				fmt.Println(front)
				Q = append(Q,front)
				break
			}
		}
	}
}



func main() {
	maze:=readMaze("maze/maze.in")
	fmt.Println(maze)
	for _,row := range maze{
		for _,val := range row{
			fmt.Printf("%d",val)
		}
		fmt.Println()
	}

	steps := walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})

	for _,row := range steps {
		for _,val := range row{
			fmt.Printf("%3d",val)
		}
		fmt.Println()
	}

	printPath(steps)
	PrintPath2(steps,point{len(maze)-1,len(maze[0])-1})
}


