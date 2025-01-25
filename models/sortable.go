package models

import (
	"sort"
)

type Sortable interface {
	GetPosition() uint
	SetPosition(uint)
}

type sortableArray []Sortable

func NewSortableArray(items []Sortable) sortableArray {
	casted := sortableArray(items)
	casted.Sort()
	return casted
}

func (l sortableArray) Sort() {
	sort.Slice(l, func(i, j int) bool {
		return l[i].GetPosition() < l[j].GetPosition()
	})
}

func (l sortableArray) Move(from, to uint) {
	tmp := l[to]

	l[to] = l[from]
	l[to].SetPosition(to)

	l[from] = tmp
	l[from].SetPosition(from)
}

func (l sortableArray) Insert(item Sortable, position uint) {
	l = append(l, item)
	copy(l[position+1:], l[position:])
	l[position] = item

	for i := position; i < uint(len(l)); i++ {
		l[i].SetPosition(i)
	}
}

func (l sortableArray) Delete(position uint) {
	l = append(l[:position], l[position+1:]...)
}

