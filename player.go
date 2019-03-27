package robot

import (
    "fmt"
    "time"

    "github.com/gosuri/uilive"
)

type ChessPlayer struct {
    agent    *Agent
    step     int
    interval time.Duration
    writer   *uilive.Writer
}

func NewChessPlayer(cols, maxn, ms int) *ChessPlayer {
    e := NewEnvironment(cols, maxn)
    a := NewAgent(e)
    return &ChessPlayer{a, 0, time.Duration(ms) * time.Millisecond, uilive.New()}
}

func (p *ChessPlayer) move (r Direction) bool {
    if status := p.agent.Move(r); status {
        p.step++
        return true
    }
    return false
}

func (p *ChessPlayer) show() {
    time.Sleep(p.interval)
    fmt.Fprintf(p.writer, "%s", p.agent.env)
}

func (p *ChessPlayer) walkToRightMost() {
    for {
        // walk through the continues blank positions, if any
        if !p.agent.CurPosHasChess() {
            p.move(Right)
            p.show()
        } else {
            break
        }
    }
    for {
        // walk through the continues chess positions, then one more step
        if p.agent.CurPosHasChess() {
            p.move(Right)
            p.show()
        } else {
            break
        }
    }
    // finally, walk 1 step back
    p.move(Left)
    p.show()
}

func (p *ChessPlayer) walkOneStepUp() {
    p.move(Up)
    p.show()
}

func (p *ChessPlayer) walkOneStepDown() {
    p.move(Down)
    p.show()
}

func (p *ChessPlayer) pickTuilLeftMost() bool {
    for {
        if p.agent.PickUp() {
            p.show()
        }
        if p.move(Left) {
            p.show()
        } else {
            break
        }
    }
    return p.agent.CarryWithChess()
}

func (p *ChessPlayer) putFromRightToLeft() {
    for {
        if p.agent.PutDown() {
            p.show()
        } else {
            // step back one step
            p.move(Right)
            p.show()
            break
        }
        if p.move(Left) {
            p.show()
        }
    }
}

func (p *ChessPlayer) Play() {
    p.step = 0
    p.writer.Start()
    fmt.Fprintf(p.writer, "%s", p.agent.env)

    // first, traiverse the bottom row
    p.walkToRightMost()

    // then, to the upper row
    p.walkOneStepUp()

    for {
        // loop: pick all the chesses in this row
        if !p.pickTuilLeftMost() {
            break
        }
        // move down to the lower row
        p.walkOneStepDown()
        // go to the right most
        p.walkToRightMost()
        // move up back to this row
        p.walkOneStepUp()
        // put down all the chesses carried now
        p.putFromRightToLeft()
        // go to the right most again
        p.walkToRightMost()
        // go to the upper row for a new loop
        p.walkOneStepUp()
    }

    p.writer.Stop()
}

func (p ChessPlayer) Summary() {
    fmt.Printf("\nChessBoard Size: %d(rows) X %d(columns)\n", p.agent.env.rows-1, p.agent.env.cols-1)
    fmt.Printf("  Total Chesses: %d\n", p.agent.env.total)
    fmt.Printf("     Walk Steps: %d\n", p.step)
}
