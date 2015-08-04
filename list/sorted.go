package list

import (
	"log"
)

import (
	"github.com/timtadh/data-structures/types"
	"github.com/timtadh/data-structures/errors"
)


type Sorted struct {
	list List
	allowDups bool
}

func NewSorted(initialSize int, allowDups bool) *Sorted {
	return &Sorted{
		list: List{
			list: make([]types.Hashable, 0, initialSize),
		},
		allowDups: allowDups,
	}
}

func SortedFromSlice(items []types.Hashable, allowDups bool) *Sorted {
	s := NewSorted(len(items), allowDups)
	for _, item := range items {
		err := s.Add(item)
		if err != nil {
			log.Panic(err)
		}
	}
	return s
}

func (s *Sorted) Clear() {
	s.list.Clear()
}

func (s *Sorted) Size() int {
	return s.list.Size()
}

func (s *Sorted) Has(item types.Hashable) (has bool) {
	_, has, err := s.Find(item)
	if err != nil {
		log.Println(err)
		return false
	}
	return has
}

func (s *Sorted) Extend(other types.KIterator) (err error) {
	for item, next := other(); next != nil; item, next = next() {
		err := s.Add(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Sorted) Get(i int) (item types.Hashable, err error) {
	return s.list.Get(i)
}

func (s *Sorted) Add(item types.Hashable) (err error) {
	i, has, err := s.Find(item)
	if err != nil {
		return err
	} else if s.allowDups {
		return s.list.Insert(i, item)
	} else if !has {
		return s.list.Insert(i, item)
	}
	return nil
}

func (s *Sorted) Remove(item types.Hashable) (err error) {
	i, has, err := s.Find(item)
	if err != nil {
		return err
	} else if !has {
		return errors.Errorf("item %v not in the table", item)
	}
	return s.list.Remove(i)
	return nil
}

func (s *Sorted) Equals(b types.Equatable) bool {
	return s.list.Equals(b)
}

func (s *Sorted) Less(b types.Sortable) bool {
	return s.list.Less(b)
}

func (s *Sorted) Hash() int {
	return s.list.Hash()
}

func (s *Sorted) Items() (it types.KIterator) {
	return s.list.Items()
}

func (s *Sorted) Find(item types.Hashable) (int, bool, error) {
	var l int = 0
	var r int = s.Size() - 1
	var m int
	for l <= r {
		m = ((r - l) >> 1) + l
		im, err := s.list.Get(m)
		if err != nil {
			return -1, false, err
		}
		if item.Less(im) {
			r = m - 1
		} else if item.Equals(im) {
			for j := m; j > 0; j-- {
				ij_1, err := s.list.Get(j-1)
				if err != nil {
					return -1, false, err
				}
				if !item.Equals(ij_1) {
					return j, true, nil
				}
			}
			return 0, true, nil
		} else {
			l = m + 1
		}
	}
	return l, false, nil
}
