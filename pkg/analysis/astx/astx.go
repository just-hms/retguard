package astx

import (
	"go/ast"
	"slices"
)

type Block struct {
	*ast.BlockStmt
	state State
}

type State uint16

const (
	NOTHING  State = 0
	ASSIGNED State = 1
	SHADOWED State = 2
)

type Blocks struct {
	data []*Block
}

func (bb *Blocks) Insert(b *ast.BlockStmt) {
	bb.data = append(bb.data, &Block{b, NOTHING})
}

func (bb *Blocks) Assign(n ast.Node) {
	if bb.GetState(n) == SHADOWED {
		return
	}
	for _, block := range bb.data {
		if n.Pos() >= block.Lbrace && n.End() <= block.Rbrace {
			block.state = ASSIGNED
			return
		}
	}
}

func (bb *Blocks) Shadow(n ast.Node) {
	rev := slices.Clone(bb.data)
	slices.Reverse(rev)
	for _, block := range rev {
		if n.Pos() >= block.Lbrace && n.End() <= block.Rbrace {
			block.state = SHADOWED
			return
		}
	}
}

func (bb *Blocks) GetState(n ast.Node) State {
	state := NOTHING
	for _, block := range bb.data {
		if n.Pos() >= block.Lbrace && n.End() <= block.Rbrace {
			// todo: maybe do this better, in a more explainable way
			if block.state > state {
				state = block.state
			}
		}
	}
	return state
}
