package stack

type intStack struct {}

type IStack interface {

}

func Create() IStack {
	return &intStack{}
}