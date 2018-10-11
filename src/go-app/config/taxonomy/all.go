package taxonomy

const (
	MethodHead = "head"
	MethodGet  = "get"
	MethodPost = "post"
)

var (
	Methods = map[string]bool{
		MethodHead: true,
		MethodGet:  true,
		MethodPost: true,
	}
)
