package board

type DefaultMap struct {
}

func (defaultMap *DefaultMap) Tiles() []int {
	var nodes []int
	for _, item := range defaultMap.Connections() {
		if !defaultMap.contains(nodes, item[0]) {
			nodes = append(nodes, item[0])
		}
	}
	return nodes
}

func (defaultMap *DefaultMap) Connections() [][]int {
	return [][]int{
		{0, 1, 1},
		{0, 3, 3},
		{0, 2, 4},
		{1, 1, 2},
		{1, 2, 5},
		{1, 3, 4},
		{1, 4, 0},

		{2, 2, 6},
		{2, 3, 5},
		{2, 4, 1},

		{3, 0, 0},
		{3, 1, 4},
		{3, 2, 8},
		{3, 3, 7},

		{4, 0, 1},
		{4, 1, 5},
		{4, 2, 9},
		{4, 3, 8},
		{4, 4, 3},
		{4, 5, 0},

		{5, 0, 2},
		{5, 1, 6},
		{5, 2, 10},
		{5, 3, 9},
		{5, 4, 4},
		{5, 5, 1},

		{6, 2, 11},
		{6, 3, 10},
		{6, 4, 5},
		{6, 5, 2},

		{7, 0, 3},
		{7, 1, 8},
		{7, 2, 12},

		{8, 0, 4},
		{8, 1, 9},
		{8, 2, 13},
		{8, 3, 12},
		{8, 4, 7},
		{8, 5, 3},

		{9, 0, 5},
		{9, 1, 10},
		{9, 2, 14},
		{9, 3, 13},
		{9, 4, 8},
		{9, 5, 4},

		{10, 0, 6},
		{10, 1, 11},
		{10, 2, 15},
		{10, 3, 14},
		{10, 4, 9},
		{10, 5, 5},

		{11, 3, 15},
		{11, 4, 10},
		{11, 5, 6},

		{12, 0, 8},
		{12, 1, 13},
		{12, 2, 16},
		{12, 5, 7},

		{13, 0, 9},
		{13, 1, 14},
		{13, 2, 17},
		{13, 3, 16},
		{13, 4, 12},
		{13, 5, 8},

		{14, 0, 10},
		{14, 1, 15},
		{14, 2, 18},
		{14, 3, 17},
		{14, 4, 13},
		{14, 5, 9},

		{15, 0, 11},
		{15, 3, 18},
		{15, 4, 14},
		{15, 5, 10},

		{16, 0, 13},
		{16, 1, 17},
		{16, 5, 12},

		{17, 0, 14},
		{17, 1, 18},
		{17, 4, 16},
		{17, 5, 13},

		{18, 0, 15},
		{18, 4, 17},
		{18, 5, 14},
	}
}

func (defaultMap *DefaultMap) contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetDefaultMap() Map {
	gboard := newMap()
	gboard.build(new(DefaultMap))
	return gboard
}