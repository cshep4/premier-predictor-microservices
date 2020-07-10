package tracer

type tracer struct{}

func New() *tracer {
	return &tracer{}
}
