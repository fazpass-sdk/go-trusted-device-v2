package gotdv2

type TrustedDevice interface {
	Extract(meta string) (Meta, error)
}
