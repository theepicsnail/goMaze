package main
import "bytes"

type RuneGetter interface {
    RuneAt(int, int) rune
    GetData() [][]Cell
}

type Maze struct { 
    rows, cols int
    data [][] Cell

}
func CellAt(maze RuneGetter, pos Position) Cell {
    return maze.GetData()[pos.row][pos.col] 
}
func String(maze RuneGetter) string {
    var buffer bytes.Buffer
    var data = maze.GetData()
    for row := range data {
        for col := range data[row] {
            buffer.WriteRune(maze.RuneAt(row, col))
        }
        buffer.WriteByte('\n')
    }
    return buffer.String()
}

func (maze Maze) String() string { 
    return String(maze); 
}

func (maze Maze) RuneAt(row, col int) rune {
    return maze.data[row][col].rune
}

func (maze Maze) GetData() [][]Cell { 
    return maze.data 
}

func (maze *Maze) getNeighbors(pos Position) []Position {
    var out []Position
    if pos.row-2 > 0 && maze.RuneAt(pos.row-2,pos.col) == WALL{
        out = append(out, Position{pos.row-2, pos.col})
    }
    if pos.row+2 < maze.rows && maze.RuneAt(pos.row+2,pos.col) == WALL{
        out = append(out, Position{pos.row+2, pos.col})
    }
    if pos.col-2 > 0 && maze.RuneAt(pos.row,pos.col-2) == WALL{
        out = append(out, Position{pos.row, pos.col-2})
    }
    if pos.col+2 < maze.cols && maze.RuneAt(pos.row,pos.col+2) == WALL{
        out = append(out, Position{pos.row, pos.col+2})
    }
    return out
}

func (maze *Maze) connect(p1, p2 Position) {
    cell :=  maze.data[p1.row][p1.col]
    cell.rune = OPEN
    cell.enterable = true
    maze.data[p1.row][p1.col]=cell

    cell = maze.data[(p1.row+p2.row)>>1][(p1.col+p2.col)>>1]
    cell.rune = OPEN
    cell.enterable = true
    maze.data[(p1.row+p2.row)>>1][(p1.col+p2.col)>>1] = cell

    cell = maze.data[p2.row][p2.col]
    cell.rune = OPEN
    cell.enterable = true
    maze.data[p2.row][p2.col] = cell
}

func NewMaze(rows, cols int) Maze {
    m := Maze{}
    cols = cols * 2 + 1 //Resize for adding in walls
    rows = rows * 2 + 1 //
    m.rows = rows
    m.cols = cols
    m.data = make([][]Cell, rows)
    for row:= range m.data {
        m.data[row] = make([]Cell, cols)
        for col := range m.data[row] {
            m.data[row][col].rune = WALL
        }
    }
    return m
}


