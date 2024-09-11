package middleware

type AuthHeader struct {
	Authorization string `header:"Authorization"`
}
