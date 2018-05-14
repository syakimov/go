package main

import (
	"sort"
	"strings"
)

func lowestManager(n int, e1 string, e2 string, r []string) string {
	employees := make(map[string]*Employee, 0)

	// create tree
	for _, v := range r {
		p := strings.Fields(v)
		ensureInMap(employees, p)
		man, sub := p[0], p[1]
		// create relation
		employees[man].Subordinates = append(employees[man].Subordinates, employees[sub])
		employees[sub].Managers = append(employees[sub].Managers, employees[man])
	}

	// get first employee managers
	man1 := make(map[string]bool)
	getManagers(employees[e1], e1, &man1)
	if _, present := man1[e2]; present {
		return e2
	}

	// get second employee managers
	man2 := make(map[string]bool)
	getManagers(employees[e2], e2, &man2)
	if _, present := man2[e1]; present {
		return e1
	}

	// common
	cm := make(map[string]bool)
	for k, _ := range man1 {
		if _, present := man2[k]; present {
			cm[k] = true
		}
	}

	// create hierarchy
	h := make(map[int][]string)
	root := employees[strings.Fields(r[0])[0]]
	createHyrarchy(root, &h, &cm, 0)

	keys := make([]int, 0)
	for k := range h {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	manager := ""
	for i := len(keys) - 1; i >= 0; i-- {
		level := keys[i]
		if len(h[level]) == 1 {
			manager = h[level][0]
			break
		}
	}

	return manager
}

func getManagers(e *Employee, r string, managers *map[string]bool) {
	if e == nil {
		return
	}

	if e.Name != r {
		(*managers)[e.Name] = true
	}

	for _, man := range e.Managers {
		getManagers(man, r, managers)
	}
}

func createHyrarchy(r *Employee, m *map[int][]string, filter *map[string]bool, lvl int) {
	if r == nil {
		return
	}

	if _, ok := (*filter)[r.Name]; ok {
		nodes, present := (*m)[lvl]
		if !present {
			(*m)[lvl] = make([]string, 0)
			nodes, _ = (*m)[lvl]
		}

		if !contains(&nodes, r.Name) {
			(*m)[lvl] = append(nodes, r.Name)
		}
	}

	for _, e := range r.Subordinates {
		createHyrarchy(e, m, filter, lvl+1)
	}
}

func ensureInMap(m map[string]*Employee, l []string) {
	for _, e := range l {
		_, present := m[e]

		if !present {
			m[e] = &Employee{e, make([]*Employee, 0), make([]*Employee, 0)}
		}
	}
}

func contains(s *[]string, val string) bool {
	for _, str := range *s {
		if str == val {
			return true
		}
	}

	return false
}

// Employee is a building block for the company tree hierarchy.
type Employee struct {
	Name         string
	Subordinates []*Employee
	Managers     []*Employee
}
