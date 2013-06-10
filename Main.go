package main
import "fmt"
import "time"

func main() {
    var rows = 15
    var cols = 30
    fmt.Println("Generate.")

    //maze := Generate(rows,cols, new(SpanningTreeGenerator))
    //maze := Generate(rows,cols, new(DepthFirstGenerator))
    maze := Generate(rows,cols, new(RDivGenerator))

    fmt.Println(ToString(&maze))

    pMaze := PlayerMaze{maze, Position{3,3}, false}
    pMaze.MoveTo(Position{1,1})

    for true {
        fmt.Println(pMaze)
        c,_ := getCh()
        fmt.Println()
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
            fmt.Println()
            fmt.Println(pMaze)
            time.Sleep(1e8)
            pMaze.bump = false
        }
    }
}
// ☺
