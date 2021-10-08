package main

type Proxy interface {
	Passthrough(s string) string
}

type pr struct{}

func (p *pr) Passthrough(s string) string {
	return s
}

type prpr struct {
	pr Proxy
}

func (p *prpr) validate(s string) string {
	return p.pr.Passthrough(s)
}
