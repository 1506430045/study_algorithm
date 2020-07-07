package skip_list

import (
	"math/rand"
	"time"
)

const P = 0.6
const MaxLevel = 8

func randomLevel() int {
	i := 1
	rand.Seed(time.Now().UnixNano())
	for i < MaxLevel {
		p := rand.Float64()
		if p < P {
			i++
		} else {
			break
		}
	}
	return i
}

type node struct {
	level int
	nexts []*node
	v     int
}

type SkipList struct {
	head *node
}

func NewSkipList() *SkipList {
	return &SkipList{
		head: &node{
			level: MaxLevel,
			nexts: make([]*node, MaxLevel),
			v:     0,
		},
	}
}

func (s *SkipList) insert(v int) {
	l := randomLevel()
	add := new(node)
	add.level = 1
	add.nexts = make([]*node, l)
	add.v = v
	i := 1

	p := s.head
	for i > 0 {
		n := p.nexts[i-1]
		for n != nil && n.v < v {
			p = n
			n = n.nexts[i-1]
		}
		p.nexts[i-1] = add
		add.nexts[i-1] = n
		i--
	}
}
