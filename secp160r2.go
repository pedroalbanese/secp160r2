// Parameters for the secp160k1 Elliptic curve
package secp160r2

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once
var p160 *elliptic.CurveParams

func initP160() {
	p160 = new(elliptic.CurveParams)
	p160.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFAC73", 16)
	p160.N, _ = new(big.Int).SetString("0100000000000000000000351EE786A818F3A1A16B", 16)
	p160.B, _ = new(big.Int).SetString("B4E134D3FB59EB8BAB57274904664D5AF50388BA", 16)
	p160.Gx, _ = new(big.Int).SetString("52DCB034293A117E1F4FF11B30F7199D3144CE6D", 16)
	p160.Gy, _ = new(big.Int).SetString("FEAFFEF2E331F296E071FA0DF9982CFEA7D43F2E", 16)
	p160.BitSize = 160
}

func P160() elliptic.Curve {
	initonce.Do(initP160)
	return p160
}
