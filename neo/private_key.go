package neo

type (
	// PrivateKey is a struct which holds the value of a NEO address private key, and has
	// a number of utility functions attached to the struct.
	PrivateKey struct {
		Value string
	}
)

// NewPrivateKey creates a PrivateKey struct.
func NewPrivateKey(s string) PrivateKey {
	return PrivateKey{
		Value: s,
	}
}
