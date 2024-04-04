package exmp

type Sample struct {
	name    string `json:"name"`
	surname string `json:"surname"`
}

type Parameter interface {
}

func MyMethod(a ...Parameter) {
}

func Usage() {

	MyMethod(Sample{
		name:    "",
		surname: "",
	})

	//------------
	MyMethod(1)

	//------------
	MyMethod("as")

	//------------
	MyMethod(true)

	//go to infinitive
	MyMethod(Sample{
		name:    "",
		surname: "",
	}, 1, "as", true)
}
