package Info

func NewResponse(name, email string) *Response {
	return &Response{
		Name:  name,
		Email: email,
	}
}

type Response struct {
	Name  string
	Email string
}
