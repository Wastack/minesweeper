package mine

import (
    "math/rand"
    "strconv"
)

type Position struct {
    X int
    Y int
}

type MineField struct {
    Pos Position
    IsMine bool
    IsRevealed bool
}

func (m *MineField) Display( sweeper MineSweeper) byte {
    if !m.IsRevealed {
        var ret byte = '_'
        return ret
    }
    if m.IsMine {
        var ret byte = 'X'
        return ret
    }
    return strconv.Itoa(int(sweeper.NeighborMineCount(m.Pos)))[0]
}

type MineMap [][]MineField

type MineSweeper interface {
    Get(p Position) (f *MineField, ok bool)
    Neighbors(p Position) []*MineField
    NeighborMineCount(p Position) uint
    AllNotMineRevealed() bool
}

func CreateMatrix(width int, height int, mine_prob float32) MineSweeper {
    m := make(MineMap, height)
    for y := range m {
        m[y] = make([]MineField, width)
        for x := range m[y] {
            m[y][x].Pos = Position{x, y}
            m[y][x].IsMine = (rand.Float32() < mine_prob)
        }
    }
    return &m
}

func (m *MineMap) AllNotMineRevealed() bool {
    // TODO cache number of mines
    for _, row := range *m {
        for _, e := range row {
            if !e.IsRevealed && !e.IsMine {
                return false
            }
        }
    }
    return true
}

/*
0 1 2
3   4
5 6 7
*/
func (m *MineMap) Neighbors(p Position) (nbs []*MineField) {
    nbs = make([]*MineField, 0, 8)
    nb_indices := []Position{
        {p.X-1, p.Y-1},
        {p.X, p.Y-1},
        {p.X +1, p.Y-1},
        {p.X -1, p.Y},
        {p.X +1, p.Y},
        {p.X -1, p.Y+1},
        {p.X, p.Y+1},
        {p.X +1, p.Y+1},
    }
    for _, v := range nb_indices {
        field, ok := m.Get(v)
        if ok {
            nbs = append(nbs, field)
        }
    }
    return
}

func (m *MineMap) Get(p Position) (f *MineField, ok bool) {
    if p.X < 0 || p.Y < 0  || p.Y >= len(*m) || p.X >= len((*m)[0]) {
        return
    }
    f = &(*m)[p.Y][p.X]
    ok = true
    return
}

func (m *MineMap) NeighborMineCount(p Position) (count uint) {
    nbs := m.Neighbors(p)
    for _, v := range nbs {
        if v.IsMine {
            count++
        }
    }
    return
}
