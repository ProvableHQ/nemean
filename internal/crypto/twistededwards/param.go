package twistededwards

import (
	"math/big"
)

// Param defines a Twisted Edwards curve (TEC).
type Param struct {
	Name string // Name of curve

	P big.Int // Prime defining the underlying field
	Q big.Int // Order of the prime-order base point
	R int     // Cofactor: Q*R is the total size of the curve

	A, D big.Int // Edwards curve equation parameters

	FBX, FBY big.Int // Standard base point for full group
	PBX, PBY big.Int // Standard base point for prime-order subgroup

	Elligator1s big.Int // Optional s parameter for Elligator 1
	Elligator2u big.Int // Optional u parameter for Elligator 2
}

// Return the name of this curve.
func (p *Param) String() string {
	return p.Name
}

func ParamBLS12377() *Param {
	var p Param
	//var mi mod.Int

	p.Name = "BLS12-377"

	p.D.SetInt64(3021)
	p.A.SetInt64(-1)
	p.P.SetString("8444461749428370424248824938781546531375899335154063827935233455917409239041", 10)
	//p.P.SetBit(zero, 251, 1).Sub(&p.P, big.NewInt(9))
	//p.Q.SetString("45330879683285730139092453152713398835", 10)
	//p.Q.Sub(&p.P, &p.Q).Div(&p.Q, big.NewInt(4))
	p.R = 4

	//
	//// Full-group generator is (4/V,3/5)
	//mi.InitString("4", "19225777642111670230408712442205514783403012708409058383774613284963344096", 10, &p.P)
	//p.FBX.Set(&mi.V)
	//mi.InitString("3", "5", 10, &p.P)
	//p.FBY.Set(&mi.V)
	//
	//// Elligator1 parameter s for Curve1174 (Elligator paper section 4.1)
	//p.Elligator1s.SetString("1806494121122717992522804053500797229648438766985538871240722010849934886421", 10)

	return &p
}
