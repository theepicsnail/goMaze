package main

type Position struct {
    row, col int
}

func (p *Position) MoveLeft() {
    p.col --
}

func (p *Position) MoveRight() {
    p.col ++
}

func (p *Position) MoveUp() {
    p.row --
}

func (p *Position) MoveDown() {
    p.row ++
}

func (p Position) Left() Position {
    return Position{p.row, p.col-1}
}

func (p Position) Right() Position {
    return Position{p.row, p.col+1}
}

func (p Position) Up() Position {
    return Position{p.row-1, p.col}
}

func (p Position) Down() Position {
    return Position{p.row+1, p.col}
}


