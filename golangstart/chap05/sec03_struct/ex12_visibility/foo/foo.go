package foo

type T struct {
	Field1 int
	field2 int
}

func (t *T) Method1() int {
	return t.Field1
}

func (t *T) method2() int {
	return t.field2
}
