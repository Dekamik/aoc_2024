package assert

import "fmt"

func Assert(predicate bool, becausef string, args ...any) {
	if !predicate {
		panic(fmt.Sprintf(becausef, args...))
	}
}
