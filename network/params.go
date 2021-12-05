package network

// Network denotes the network params.
type Network string

const (
	testnet1 Network = "testnet1"
	testnet2 Network = "testnet2"
)

// Params holds the network object.
type Params struct {
	network Network
}

// Network returns the Network type.
func (p Params) Network() Network {
	return p.network
}

// Testnet1 returns Testnet1 params.
func Testnet1() *Params {
	return &Params{network: testnet1}
}

// Testnet2 returns Testnet2 params.
func Testnet2() *Params {
	return &Params{network: testnet2}
}
