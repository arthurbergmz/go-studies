package closure

import "../../tests"

type twoIntegersResolverInterface func (a, b int) int

func twoIntegersResolver (resolver twoIntegersResolverInterface) twoIntegersResolverInterface {
	return func (a, b int) int {
		return resolver(a, b)
	}
}

func TwoIntegersTest (expected []int) tests.TestRunner {
	return func () ([]int, []int) {
		soma := twoIntegersResolver(func (a, b int) int {
			return a + b
		})
		subtrair := twoIntegersResolver(func (a, b int) int {
			return a - b
		})
		results := []int{ soma(1, 2), subtrair(10, 3) }
		return results, expected
	}
}