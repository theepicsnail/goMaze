package main
import "fmt"
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

func main() {
    rand.Seed( time.Now().UTC().UnixNano())
    maze := NewMaze(10,20)
    fmt.Println(maze)
    var posList []Position
    posList = append(posList, Position{1,1})
    var item Position
    fmt.Println("Generating maze")
    for len(posList)!=0 {
        fmt.Print(".")
        //pop a random item from the list 
        posList, item = popRandom(posList)
        neighbors := maze.getNeighbors(item)
        if len(neighbors) > 0 {
            _, neighbor := popRandom(neighbors)
            posList = append(posList, item, neighbor)
            maze.connect(item, neighbor)
        }
    }
    fmt.Println()
    fmt.Println(maze)

    pMaze := PlayerMaze{maze, Position{3,3}, false}
    pMaze.MoveTo(Position{1,1})

    for true {
        c,_ := getCh()
        valid := false
        switch c {
            case 'w', 65:
                valid = pMaze.MoveUp()
            case 'a', 68:
                valid = pMaze.MoveLeft()
            case 'd', 67:
                valid = pMaze.MoveRight()
            case 's', 66:
                valid = pMaze.MoveDown()
            default:
                valid = true
        }
        if !valid {
            pMaze.bump = true
            fmt.Println(pMaze)
            time.Sleep(1e8)
            pMaze.bump = false 
        }
        fmt.Println(pMaze)
    }
}
// ☺
