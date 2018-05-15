package main

import (
	"fmt"
	"sync"
	"time"
)

type Table struct {
	forks []*sync.RWMutex
}

type Philosopher struct {
	name  string
	left  byte
	right byte
}

func newPhilosopher(name string, left, right byte) *Philosopher {
	return &Philosopher{
		name:  name,
		left:  left,
		right: right,
	}
}

func (p *Philosopher) eat(table *Table) {
	left := table.forks[p.left]
	right := table.forks[p.right]

	left.Lock()
	defer left.Unlock()

	time.Sleep(time.Duration(time.Millisecond * 150))

	right.Lock()
	defer right.Unlock()

	fmt.Printf("%s начала есть\n", p.name)
	time.Sleep(time.Second)
	fmt.Printf("%s закончила есть\n", p.name)
}

func main() {
	table := &Table{
		forks: []*sync.RWMutex{
			{},
			{},
			{},
			{},
			{},
		},
	}

	philosophers := []*Philosopher{
		newPhilosopher("Джудит Батлер", 0, 1),
		newPhilosopher("Рая Дунаевская", 1, 2),
		newPhilosopher("Зарубина Наталья", 2, 3),
		newPhilosopher("Эмма Гольдман", 3, 4),
		newPhilosopher("Анна Шмидт", 0, 4),
	}

	wg := sync.WaitGroup{}
	wg.Add(len(philosophers))

	for _, p := range philosophers {
		go func(philosopher *Philosopher) {
			defer wg.Done()
			philosopher.eat(table)
		}(p)
	}

	wg.Wait()
}
