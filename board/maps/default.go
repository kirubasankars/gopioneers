package maps

type DefaultMap struct {
}

func (defaultMap DefaultMap) GetTileConfig() string {
	return `
-,-,sw2,s,s?3,s,-
-,s,m,p,t,sg3,-
-,s?1,f,h,p,h,s
s,f,t,d,t,m,sl4
-,s?1,t,m,f,p,s
-,s,h,f,p,s?5,-
-,-,so0,s,sb5,s,-
	`
}

func (defaultMap DefaultMap) GetChits() []int {
	return []int{10, 2, 9, 12, 6, 4, 10, 9, 11, 3, 8, 8, 3, 4, 5, 5, 6, 11}
}