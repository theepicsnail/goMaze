package main

type Cell struct {
    rune rune
    seen bool
    enterable bool
}

func (cell *Cell) GetRune() rune {
    if cell.seen {
        return cell.rune
    }
    return WALL
}
