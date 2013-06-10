package main
import "math/rand"
import "time"

func popRandom(list []Position) ([]Position, Position) {
    length := int32(len(list))
    idx := rand.Int31n(length)

    item := list[idx]

    list[idx] = list[length-1]
    list = list[0:length-1]
    return list, item
}

type MazeGenerator func(m *Maze)

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

func Generate(rows, cols int, gen MazeGenerator) Maze {
    rand.Seed(time.Now().UTC().UnixNano())
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
            m.data[row][col].seen = true
        }
    }
    gen(&m)
    return m
}

func SpanningTreeGenerator(m *Maze) {
    //Dig out the maze
    var posList = []Position{}
    var pos Position
    posList = append(posList, Position{1,1})
    for len(posList)!= 0 {
        posList, pos = popRandom(posList)

        neighbors := m.getNeighbors(pos)
        if len(neighbors) > 0 {

            _, neighbor := popRandom(neighbors)

            m.connect(pos, neighbor)
            posList = append(posList, pos, neighbor)
        }
    }
}

func DepthFirstGenerator(m *Maze) {
    var generateFromPosition func(p Position);
    generateFromPosition = func(p Position) {
        neighbors := m.getNeighbors(p)
        var pos Position
        for len(neighbors)!=0 {
            neighbors, pos = popRandom(neighbors)
            if !m.data[pos.row][pos.col].enterable {
                m.connect(p, pos)
                generateFromPosition(pos)
            }
        }
    }
    generateFromPosition(Position{1,1})
}

func RDivGenerator(m *Maze) {
    var divideMaze func(topLeft, botRight Position);
    divideMaze = func(topLeft, botRight Position) {
        //Divide the area given by topLeft to botRight (exclude points)
        //into 4 sections
        var areaRows = botRight.row - topLeft.row
        var areaCols = botRight.col - topLeft.col
        if areaRows == 2 || areaCols == 2 { //We can subdivide a hallway
            return
        }
        var splitRow = rand.Intn(areaRows/2-1)*2//Generate an even row to divide on
        var splitCol = rand.Intn(areaCols/2-1)*2
        splitRow += topLeft.row + 2
        splitCol += topLeft.col + 2
        m.data[splitRow][splitCol].rune = WALL

        for row := topLeft.row+1; row < botRight.row ; row ++ {
            m.data[row][splitCol].rune = WALL
            m.data[row][splitCol].enterable = false
        }

        for col := topLeft.col+1; col < botRight.col ; col ++ {
            m.data[splitRow][col].rune = WALL
            m.data[splitRow][col].enterable = false
        }

        var solidWall = rand.Intn(4)
        for wall:=0 ; wall < 4; wall ++ {
            if wall == solidWall {
                continue
            }
            holeRow := 0
            holeCol := 0
            switch wall {
                case 0://up from split
                    holeRow = rand.Intn((splitRow-topLeft.row)/2)*2+1 + topLeft.row
                    holeCol = splitCol
                break
                case 1://down
                    holeRow = rand.Intn((botRight.row-splitRow)/2)*2+1 + splitRow
                    holeCol = splitCol
                break
                case 2://left
                    holeRow = splitRow
                    holeCol = rand.Intn((splitCol-topLeft.col)/2)*2+1 + topLeft.col
                break
                case 3:
                    holeRow = splitRow
                    holeCol = rand.Intn((botRight.col-splitCol)/2)*2+1 + splitCol
                break

            }
            m.data[holeRow][holeCol].enterable = true
            m.data[holeRow][holeCol].rune = OPEN

        }

        divideMaze(topLeft, Position{splitRow , splitCol})
        divideMaze(Position{splitRow , splitCol}, botRight)
        divideMaze(Position{topLeft.row, splitCol}, Position{splitRow , botRight.col})
        divideMaze(Position{splitRow, topLeft.col}, Position{botRight.row , splitCol})
        //..A........
        //...........
        //....#.#....   Generate a point in the # area
        //...........
        //....#.#....
        //...........
        //....#.#....
        //...........
        //........B..
    }

    for row := m.rows-2; row >= 1; row -- {
        for col := m.cols-2; col >= 1; col -- {
            m.data[row][col].rune = OPEN
            m.data[row][col].enterable = true
        }
    }
    divideMaze(Position{0,0}, Position{m.rows-1, m.cols-1})
}

