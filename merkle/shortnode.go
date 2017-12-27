package merkle

import (
	"encoding/hex"
	"github.com/spacemeshos/go-spacemesh/merkle/pb"
)

// shortNode is an immutable leaf or an extension node
type shortNode interface {
	IsLeaf() bool
	GetValue() []byte
	GetPath() string // hex encoded string
	GetParity() bool
	Marshal() pb.Node
}

func newShortNode(data *pb.Node) shortNode {

	n := &shortNodeImpl{
		nodeType: data.NodeType,
		parity:   data.Parity,
		path:     data.Path,
		value:    data.Value,
	}
	return n
}

type shortNodeImpl struct {
	nodeType pb.NodeType // extension node when false
	parity   bool        // path parity - when odd, truncate first nibble prefix to return path
	path     []byte
	value    []byte
}

func (s *shortNodeImpl) GetValue() []byte { return s.value }
func (s *shortNodeImpl) GetParity() bool  { return s.parity }
func (s *shortNodeImpl) IsLeaf() bool     { return s.nodeType == pb.NodeType_leaf }

func (s *shortNodeImpl) GetPath() string {
	// todo: consider parity
	return hex.EncodeToString(s.path)
}

func (s *shortNodeImpl) Marshal() pb.Node {

	res := pb.Node{
		NodeType: s.nodeType,
		Value:    s.value,
		Parity:   s.parity,
		Path:     s.path,
	}

	return res
}