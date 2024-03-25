package main

type iPhone struct {
}

func (ip iPhone) ara() {

}

type samsung struct {
}

func (sa samsung) ara() {

}

type telefon interface {
	ara()
}

type Object1 struct {
}

func (obj Object1) object() {

}

type Object interface {
	object()
}

func main() {

	var b int = 10
	var a *int = &b
	println(&a)

	iphone := iPhone{}
	_samsung := samsung{}

	var telefon telefon
	telefon = &iphone
	telefon = &_samsung

	telefon.ara()

	getObject(Object1{})

}

func ara(t telefon) {
	t.ara()

}

func getObject(object Object) {

}
