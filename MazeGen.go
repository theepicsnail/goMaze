package main

type Maze struct {
    rows, cols int
    data [][]Cell
}

//Implement the MazeI interface
func (m *Maze) RuneAt(p Position) rune {
    
    return m.data[p.row][p.col].GetRune()
}

func (m *Maze) GetWidth() int {
    return m.cols
}

func (m *Maze) GetHeight() int {
    return m.rows
}
//End of MazeI interface.


func (m Maze) getNeighbors(pos Position) []Position {
    //Find possible neighbors to connect to.
    var out []Position

    possible := [4]Position{}
    possible[0] = Position{pos.row - 2, pos.col}
    possible[1] = Position{pos.row + 2, pos.col}
    possible[2] = Position{pos.row, pos.col - 2}
    possible[3] = Position{pos.row, pos.col + 2}

    
    for _,p := range(possible) {
        
        if p.row < 0 || p.col < 0 || p.row >= m.rows || p.col >= m.cols {
            continue
        }
        //Cell has not been linked to the maze yet.
        if !m.data[p.row][p.col].enterable {
            out = append(out, p)
        }
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

func Generate(rows, cols int) Maze {
    //Make our maze data structure
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

    //Dig out the maze
    var posList = []Position{}
    posList = append(posList, Position{1,1})
    var pos Position
    for len(posList)!= 0 {
        posList, pos = popRandom(posList)

        neighbors := m.getNeighbors(pos)
        if len(neighbors) > 0 {

            _, neighbor := popRandom(neighbors)
            m.connect(pos, neighbor)
            
            posList = append(posList, pos, neighbor)
        }
    }

    return m
}


