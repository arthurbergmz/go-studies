package tests

type TestValidator func (result bool)

type TestRunner func () ([]int, []int)

func Test (validator TestValidator, testRunner TestRunner) {
	testResult := func () bool {
		results, expected := testRunner()
		if len(results) != len(expected) {
			return false
		}
		for index, value:= range results {
			if value != expected[index] {
				return false
			}
		}
		return true
	}()
	validator(testResult)
}