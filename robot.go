package robot

type Direction int

const (
    Left Direction = iota
    Up
    Right
    Down
)

type Robot interface {
    // based on the current position, move left/up/right/down if possible
    Move(r Direction) bool
    // check if current position has a chess
    CurPosHasChess() bool
    // check if the agent carry chess(es) in its pocket
    CarryWithChess() bool
    // pick up a chess at the current position
    PickUp() bool
    // put down a chess at the current position
    PutDown() bool
}

////////////////////////////////////////////////////////////////////////////////

type Agent struct {
    env     *Environment  // the game environment
    idx     int           // stack top pointer
}

func NewAgent(e *Environment) *Agent {
    // an agent must be attched with a certain environment
    // note that the agent can not count the chess it has,
    // so we should implement a virtual stack as a chess container
    return &Agent{ e, -1 }
}

func (a Agent) CarryWithChess() bool {
    return a.idx >= 0
}

func (a Agent) CurPosHasChess() bool {
    return a.env.hasChess()
}

func (a *Agent) PickUp() bool {
    if !a.env.hasChess() {
        return false
    }
    a.env.pickUp()
    a.idx++
    return true
}

func (a *Agent) PutDown() bool {
    if a.env.hasChess() {
        return false
    }
    if a.idx < 0 {
        return false
    }
    a.env.putDown()
    a.idx--
    return true
}

func (a *Agent) Move(r Direction) bool {
    return a.env.move(r)
}
