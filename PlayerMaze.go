package main

type PlayerMaze struct {
    Maze
    player Position
    bump bool
}

func (maze PlayerMaze) String() string {
    return String(maze)
}

func (maze PlayerMaze) RuneAt(row, col int) rune {
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
    return maze.Maze.RuneAt(row,col)
}

func (maze *PlayerMaze) MoveTo (p Position) bool {
    if maze.data[p.row][p.col].enterable {
        maze.player = p
        maze.data[p.row+1][p.col].seen = true
        maze.data[p.row-1][p.col].seen = true
        maze.data[p.row][p.col+1].seen = true
        maze.data[p.row][p.col-1].seen = true
        maze.data[p.row+1][p.col+1].seen = true
        maze.data[p.row-1][p.col+1].seen = true
        maze.data[p.row+1][p.col-1].seen = true
        maze.data[p.row-1][p.col-1].seen = true
        return true
    } 
    return false
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
