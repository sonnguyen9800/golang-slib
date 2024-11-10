package test

type AnyParamsBoolFunc func(...interface{}) bool

type GenericTestCaseBoolOutput[T any] struct {
	Expected bool
	Input    T
}
