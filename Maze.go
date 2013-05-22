package main
import "bytes"

type MazeI interface {
    RuneAt(Position) (rune)
    GetWidth() int
    GetHeight() int
}

func ToString(m MazeI) string {
    var buffer bytes.Buffer
    width := m.GetWidth()
    height:= m.GetHeight()
    var p Position
    for p.row = 0; p.row < height; p.row++ {
        for p.col = 0; p.col < width; p.col++ {
            buffer.WriteRune(m.RuneAt(p))
        }
        buffer.WriteByte('\n')
    }
    return buffer.String()
}

