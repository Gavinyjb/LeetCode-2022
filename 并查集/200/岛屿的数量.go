package leetcdoe

type UnionFindInterface interface {
	Init()
	Find()
	Union()
	GetCount()
	IsConnect() bool
}
type UnionFind struct {
	parent map[string]string
	rank   map[string]string
	count  int
}
