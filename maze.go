package main
import "os"
import "os/exec"
import "fmt"
import "bytes"
import "math/rand"
//import "time"

//Do all the getCh setup and declaration.
//makes use of the os and os/exec imports.
var _ = exec.Command("/bin/stty", "-F", "/dev/tty", "-icanon", "min", "1" ).Run()
var getChBuff = make([]byte, 1)
func getCh() (byte, error){
    ct,err := os.Stdin.Read(getChBuff)
    if ct !=0  {
        return getChBuff[0], nil
    }
    return 0, err
}

type Cell struct {
    rune rune
    seen bool
}

func (cell *Cell) getRune() rune {
    if cell.seen {
        return cell.rune
    }
    return rune('#')
}

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

type Position struct {
    row, col int
}

func popRandom(list []Position) ([]Position, Position) {
    length := int32(len(list))
    idx := rand.Int31n(length)

    item := list[idx]

    list[idx] = list[length-1]
    list = list[0:length-1]
    return list, item
}

func main() {
    maze := NewMaze(10,20)
    fmt.Println(maze)
    getCh()//Wait for any key
    var posList []Position
    posList = append(posList, Position{1,1})
    var item Position
    for len(posList)!=0 {
        //pop a random item from the list 
        posList, item = popRandom(posList)
        neighbors := maze.getNeighbors(item)
        if len(neighbors) > 0 {
            _, neighbor := popRandom(neighbors)
            posList = append(posList, item, neighbor)
            maze.connect(item, neighbor)
        }
    }
    fmt.Println(maze)
}
// ☺
