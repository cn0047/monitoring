package BLError

// Panic performs panic with BLError.
func Panic(err error) {
	panic(New(err.Error()))
}

// Panic performs panic with formatted BLError.
func Panicf(format string, params ...interface{}) {
	panic(New(format, params...))
}
