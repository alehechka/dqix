package router

type app struct {
}

type RouterOption interface {
	apply(*app)
}
