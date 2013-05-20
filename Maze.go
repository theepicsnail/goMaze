package main
import "bytes"

type Maze struct { 
    rows, cols int
    data [][] Cell
}

func (maze *Maze) String() string {
    var buffer bytes.Buffer

    for row := range maze.data {
        for col := range maze.data[row] {
            buffer.WriteRune(maze.runeAt(row, col))
        }
        buffer.WriteByte('\n')
    }
    return buffer.String()
}

func (maze *Maze) runeAt(row, col int) rune {
    return maze.data[row][col].rune
}

func (maze *Maze) getNeighbors(pos Position) []Position {
    var out []Position
    if pos.row-2 > 0 && maze.runeAt(pos.row-2,pos.col) == '#'{
        out = append(out, Position{pos.row-2, pos.col})
    }
    if pos.row+2 < maze.rows && maze.runeAt(pos.row+2,pos.col) == '#'{
        out = append(out, Position{pos.row+2, pos.col})
    }
    if pos.col-2 > 0 && maze.runeAt(pos.row,pos.col-2) == '#'{
        out = append(out, Position{pos.row, pos.col-2})
    }
    if pos.col+2 < maze.cols && maze.runeAt(pos.row,pos.col+2) == '#'{
        out = append(out, Position{pos.row, pos.col+2})
    }
    return out
}

func (maze *Maze) connect(p1, p2 Position) {
   maze.data[p1.row][p1.col].rune = ' '
   maze.data[(p1.row+p2.row)>>1][(p1.col+p2.col)>>1].rune = ' '
   maze.data[p2.row][p2.col].rune = ' ' 
}

func NewMaze(rows, cols int) *Maze {
    m := new(Maze)
    cols = cols * 2 + 1 //Resize for adding in walls
    rows= rows* 2 + 1 //
    m.rows = rows
    m.cols = cols
    m.data = make([][]Cell, rows)
    for row:= range m.data {
        m.data[row] = make([]Cell, cols)
        for col := range m.data[row] {
            m.data[row][col].rune = '#'
        }
    }
    return m
}


