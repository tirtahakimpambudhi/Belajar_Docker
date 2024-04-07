package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Item struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Qtt   int       `json:"quantity"`
	Price *Currency `json:"price"`
}

type Currency struct {
	Code   string  `json:"code"`
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
}

type Stores struct {
	DB map[string]*Item
	m sync.Mutex
}

func NewStores() *Stores {
	return &Stores{DB: make(map[string]*Item)}
}

func (s *Stores) Save(item *Item)  {
	s.m.Lock()
	defer s.m.Unlock() 
	id := uuid.NewString()
	item.ID = id
	s.DB[id] = item
}
func (s *Stores) FindByID(id string) (*Item,error) {
	s.m.Lock()
	defer s.m.Unlock()
	item , ok := s.DB[id]
	if !ok {
		return nil,fmt.Errorf("item with id %s not found",id)
	}
	return item,nil
}

func (s *Stores) FindAll() []*Item {
	var items []*Item
	s.m.Lock()
	defer s.m.Unlock()
	for _,item  := range s.DB {
		items = append(items, item)
	}
	return items
}