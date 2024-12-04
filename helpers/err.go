package helpers

func HandleErr(e error) {
	if e != nil {
		panic(e)
	}
}
