package robot

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    kEmpty = iota     // empty position
    kAgent            // a agent is here
    kChess            // a chess is here
    kCombo            // agent + chess are here
)

type Position struct {
    ridx int          // index of row
    cidx int          // index of column
}

type Environment struct {
    agent      Position
    rows       int
    cols       int
    total      int
    chessboard [][]int
}

func NewEnvironment(n, max int) *Environment {
    rand.Seed(time.Now().Unix())
    c, m, t := make([]int, n), 0, 0
    for i := 0; i < n; i++ {
        c[i] = rand.Int() % max + 1
        if m < c[i] {
            m = c[i]
        }
        t += c[i]
    }

    cb := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        cb[i] = make([]int, n+1)
    }
    for i := 0; i < n; i++ {
        for j := 0; j <= m; j++ {
            if j < c[i] {
                cb[j][i] = kChess
            } else {
                cb[j][i] = kEmpty
            }
        }
    }
    for j := 0; j <= m; j++ {
        cb[j][n] = kEmpty
    }
    cb[0][0] = kCombo

    return &Environment{ Position{0, 0}, m+1, n+1, t, cb }
}

func (e Environment) symbol(i, j int) string {
    // inner method, without check boundary
    switch e.chessboard[i][j] {
    case kChess:
        return "*"
    case kAgent:
        return "○"
    case kCombo:
        return "⊛"
    default:
        return " "
    }
}

func (e Environment) String() string {
    s := ""
    for i := e.rows - 1; i >= 0; i-- {
        for j := 0; j < e.cols; j++ {
            s += fmt.Sprintf(" %s", e.symbol(i, j))
        }
        s += "\n"
    }
    return s
}

func (e Environment) hasChess() bool {
    // see if the agent position has chess
    return e.chessboard[e.agent.ridx][e.agent.cidx] == kCombo
}

func (e *Environment) pickUp() {
    // now the agent position should be kCombo, pick it up
    e.chessboard[e.agent.ridx][e.agent.cidx] = kAgent
}

func (e *Environment) putDown() {
    // now the agent position should be kAgent, put a chess down
    e.chessboard[e.agent.ridx][e.agent.cidx] = kCombo
}

func (e *Environment) move(r Direction) bool {
    // now the agent position value should be kAgent or kCombo
    // kAgent - 1 = kEmpty, kCombo - 1 = kChess
    // kEmpty + 1 = kAgent, kChess + 1 = kCombo
    switch r {
    case Left:
        if e.agent.cidx == 0 {
            return false
        }
        e.chessboard[e.agent.ridx][e.agent.cidx] -= 1
        e.agent.cidx--
        e.chessboard[e.agent.ridx][e.agent.cidx] += 1
    case Right:
        e.chessboard[e.agent.ridx][e.agent.cidx] -= 1
        e.agent.cidx++
        e.chessboard[e.agent.ridx][e.agent.cidx] += 1
    case Up:
        e.chessboard[e.agent.ridx][e.agent.cidx] -= 1
        e.agent.ridx++
        e.chessboard[e.agent.ridx][e.agent.cidx] += 1
    case Down:
        if e.agent.ridx == 0 {
            return false
        }
        e.chessboard[e.agent.ridx][e.agent.cidx] -= 1
        e.agent.ridx--
        e.chessboard[e.agent.ridx][e.agent.cidx] += 1
    }
    return true
}
