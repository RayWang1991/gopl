package ex5_11

import (
	"errors"
)

var CircleError = errors.New("Circle")

// topo sort, if detect circle report error
func TopoSort(m map[string][]string) ([]string, error) {
	foundG := map[string]bool{}
	foundL := map[string]bool{}

	order := []string{}
	var visit func(string) error
	visit = func(n string) error {
		if foundL[n] {
			return CircleError
		}
		foundL[n] = true
		if !foundG[n] {
			foundG[n] = true
			for _, i := range m[n] {
				if err := visit(i); err != nil {
					return err
				}
			}
			order = append(order, n)
		}
		delete(foundL, n)
		return nil
	}

	for k := range m {
		if err := visit(k); err != nil {
			circle := []string{}
			for k := range foundL {
				circle = append(circle, k)
			}
			return circle, err
		}
	}
	return order, nil
}
