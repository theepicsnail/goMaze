package main

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


