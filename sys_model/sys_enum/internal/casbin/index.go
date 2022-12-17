package casbin

type casbin struct {
	Event event
}

var Casbin = casbin{
	Event: Event,
}
