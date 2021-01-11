package board

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeCoordinator struct {
	index     int
	nodes     []*Node
	sides     []int
	neighbors []*NodeCoordinator
}

func (nc *NodeCoordinator) putNodes(nodes []*Node, sides []int) {
	nc.nodes = nodes
	nc.sides = sides

	f := nc.getFlag(sides)
	if f == "N" {
		for idx, v := range sides {
			nc.sides[idx] = nc.getNextSide(v)
		}
	}
}

func (nc NodeCoordinator) getNextSide(side int) int {
	if side >= 5 {
		return 0
	}
	return side + 1
}

func (nc NodeCoordinator) getPrevSide(side int) int {
	if side >= 0 {
		return 5
	}
	return side - 1
}

func (nc NodeCoordinator) getFlag(sides []int) string {
	var s []string
	for _, i := range sides {
		s = append(s, strconv.Itoa(i))
	}

	k := strings.Join(s, "#")
	if k == "0#2#4" {
		return "N"
	}
	if k == "1#3#5" {
		return "N"
	}
	if k == "2#4#0" {
		return "N"
	}
	if k == "3#5#1" {
		return "N"
	}
	if k == "4#0#2" {
		return "N"
	}
	if k == "5#1#3" {
		return "N"
	}
	if k == "0#4#2" {
		return "C"
	}
	if k == "1#5#3" {
		return "C"
	}
	if k == "2#0#4" {
		return "C"
	}
	if k == "3#1#5" {
		return "C"
	}
	if k == "4#2#0" {
		return "C"
	}
	if k == "5#3#1" {
		return "C"
	}
	return ""
}

func (nc NodeCoordinator) String() string {
	var a []string
	for idx, n := range nc.nodes {
		a = append(a, fmt.Sprintf("%d#%d", n.index, nc.sides[idx]))
	}
	return strings.Join(a, "-")
}

type CoordinatorBuilder struct {
	coordinators map[string]*NodeCoordinator
}

func (cb *CoordinatorBuilder) GetNodeCoordinator(nodes []*Node, sides []int) *NodeCoordinator {
	nc := new(NodeCoordinator)
	nc.putNodes(nodes, sides)
	name := fmt.Sprint(nc)
	if _, ok := cb.coordinators[name]; !ok {
		cb.coordinators[name] = nc
	}
	return cb.coordinators[name]
}

func (cb CoordinatorBuilder) MakeIntersectionsConnected(nodes map[int]*Node) {
	exists := func(coord *NodeCoordinator, coords []*NodeCoordinator) bool {
		for j := 0; j < len(coords); j++ {
			if coord == coords[j] {
				return true
			}
		}
		return false
	}

	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		//make connected coordinators
		for side, src := range node.coordinators {
			p := node.getPrevSide(side)
			n := node.getNextSide(side)
			coordinator0 := node.coordinators[p]
			coordinator1 := node.coordinators[n]
			if !exists(coordinator0, src.neighbors) {
				src.neighbors = append(src.neighbors, coordinator0)
			}
			if !exists(coordinator1, src.neighbors) {
				src.neighbors = append(src.neighbors, coordinator1)
			}
		}

		fmt.Println("")
	}
}

func NewCoordinatorBuilder() *CoordinatorBuilder {
	cb := new (CoordinatorBuilder)
	cb.coordinators = make(map[string]*NodeCoordinator)
	return cb
}
