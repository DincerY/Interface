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

func main() {

	iphone := iPhone{}
	_samsung := samsung{}

	var telefon telefon
	telefon = &iphone
	telefon = &_samsung

	telefon.ara()

}

func ara(t telefon) {
	t.ara()

}
