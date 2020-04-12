package errors

func FailOnError(err error) {
	if err != nil {
		panic(err)
	}
}
