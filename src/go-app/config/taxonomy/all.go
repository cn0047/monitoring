package taxonomy

const (
	// MethodHead const for head ping service method.
	MethodHead = "head"

	// MethodGet const for get ping service method.
	MethodGet = "get"

	// MethodPost const for post ping service method.
	MethodPost = "post"
)

var (
	// Methods map which contains all possible ping service methods.
	Methods = map[string]bool{
		MethodHead: true,
		MethodGet:  true,
		MethodPost: true,
	}
)
