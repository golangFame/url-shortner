package dummy

import "github.com/goferHiro/url-shortner/entities"

type Services interface {
	Dummy(dumb1 *entities.Dummy)
}

func (s *service) Dummy(dumb1 *entities.Dummy) {
	dumb1.Dummy = "dumber"
	return
}
