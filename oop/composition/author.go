package composition

type Author struct {
	FirstName, LastName, Bio string
}

func (a Author) sayHi() string {
	return "I am an author " + a.FirstName + " " + a.LastName
}
