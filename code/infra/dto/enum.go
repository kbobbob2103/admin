package dto

type StatusEmployee int

const (
	Opening StatusEmployee = iota + 1
	Disable
)
