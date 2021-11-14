package network

type Network string

const (
	testnet1 Network = "testnet1"
	testnet2 Network = "testnet2"
)

type Params struct {
	network Network
}

func (p Params) Network() Network {
	return p.network
}

func Testnet1() *Params {
	return &Params{network: testnet1}
}

func Testnet2() *Params {
	return &Params{network: testnet2}
}
