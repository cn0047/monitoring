package AppError

// Panic performs panic with AppError.
func Panic(err error) {
	panic(New(err.Error()))
}

// Panic performs panic with formatted AppError.
func Panicf(format string, params ...interface{}) {
	panic(New(format, params...))
}
