package main

type PlayerMaze struct {
    Maze
    player Position
    bump bool
}

func (maze PlayerMaze) String() string {
    return ToString(&maze)
}

func (maze PlayerMaze) RuneAt(pos Position) rune {
    row := pos.row
    col := pos.col
    if row == maze.player.row {
        if  col == maze.player.col {
            if maze.bump {
                return PLAYER_BUMP
            }
            return PLAYER
        }
    }
    if !maze.Maze.data[row][col].seen {
        return FOG
    }
    return maze.Maze.RuneAt(pos)
}

func (maze *PlayerMaze) MoveTo (p Position) bool {
    if maze.data[p.row][p.col].enterable {
        maze.setSeen(false)
        maze.player = p
        maze.setSeen(true)
        return true
    }
    return false
}
func (maze *PlayerMaze) setSeen (v bool) {
    p := maze.player
    offsets := []int{-3, -2, -1, 0, 1, 2, 3}
    for _,dr := range(offsets){
        for _,dc := range(offsets){
            r := p.row+dr
            c := p.col+dc
            if r < 0 || c < 0 || r >= maze.rows || c>= maze.cols {
                continue
            }
        maze.data[p.row+dr][p.col+dc].seen = v
        }
    } 
}

func (maze *PlayerMaze) MoveRight () bool {
    return maze.MoveTo(maze.player.Right())
}

func (maze *PlayerMaze) MoveLeft () bool {
    return maze.MoveTo(maze.player.Left())
}

func (maze *PlayerMaze) MoveUp () bool {
    return maze.MoveTo(maze.player.Up())
}

func (maze *PlayerMaze) MoveDown () bool {
    return maze.MoveTo(maze.player.Down())
}
