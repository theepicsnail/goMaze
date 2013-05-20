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


type Maze struct { 
    rows, cols int
    data [][] byte
}

func (maze *Maze) String() string {
    var buffer bytes.Buffer

    for row := range maze.data {
        for col := range maze.data[row] {
            buffer.WriteByte(maze.data[row][col])
        }
        buffer.WriteByte('\n')
    }
    return buffer.String()
}

func (maze *Maze) getNeighbors(pos Position) []Position {
    var out []Position
    if pos.row-2 > 0 && maze.data[pos.row-2][pos.col] == '#'{
        out = append(out, Position{pos.row-2, pos.col})
    }
    if pos.row+2 < maze.rows && maze.data[pos.row+2][pos.col] == '#'{
        out = append(out, Position{pos.row+2, pos.col})
    }
    if pos.col-2 > 0 && maze.data[pos.row][pos.col-2] == '#'{
        out = append(out, Position{pos.row, pos.col-2})
    }
    if pos.col+2 < maze.cols && maze.data[pos.row][pos.col+2] == '#'{
        out = append(out, Position{pos.row, pos.col+2})
    }
    return out
}

func (maze *Maze) connect(p1, p2 Position) {
   maze.data[p1.row][p1.col] = ' ' 
   maze.data[(p1.row+p2.row)>>1][(p1.col+p2.col)>>1] = ' '
   maze.data[p2.row][p2.col] = ' ' 
}

func NewMaze(rows, cols int) *Maze {
    m := new(Maze)
    cols = cols * 2 + 1 //Resize for adding in walls
    rows= rows* 2 + 1 //
    m.rows = rows
    m.cols = cols
    m.data = make([][]byte, rows)
    for row:= range m.data {
        m.data[row] = make([]byte, cols)
        for col := range m.data[row] {
            m.data[row][col] = '#'
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
