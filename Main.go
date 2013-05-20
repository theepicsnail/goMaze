package main
import "fmt"
import "math/rand"
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
