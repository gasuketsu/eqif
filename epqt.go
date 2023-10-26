package eqif

type Eqpt interface {
	Provision()
	Deprovision()
	IsProvisioned() bool
}
