package types

// Person str:struct
type person struct {
	name string
	age  int8
}

// Ahmed exported composite struct
var Ahmed = person{
	name: "Ahmed",
	age:  12,
}
