package assert

func Assert(predicate bool, because string) {
    if !predicate {
        panic(because)
    }
}
