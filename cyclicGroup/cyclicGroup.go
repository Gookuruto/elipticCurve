package cyclicGroup

import (
	"fmt"
	"math/big"
)

type Z struct {
	//prime *big.Int
	X *big.Int
}

var prime *big.Int

func NewGroup(p *big.Int) {
	prime = p
}
func New(x *big.Int) *Z {
	if prime == big.NewInt(0) || prime == nil {
		panic("No group was defined")
	}
	z := new(Z)
	z.X = new(big.Int).Mod(x, prime)
	return z
}
func (z *Z) Equal(b *Z) bool {
	return new(big.Int).Mod(z.X, prime) == new(big.Int).Mod(b.X, prime)
}

func (z *Z) Add(b *Z) *Z {
	return New(new(big.Int).Add(z.X, b.X))

}

func (z *Z) Sub(b *Z) *Z {
	return New(new(big.Int).Sub(z.X, b.X))

}

func (z *Z) Mul(b *Z) *Z {
	return New(new(big.Int).Mul(z.X, b.X))

}

func (z *Z) TrueDiv(b *Z) *Z {
	z2 := New(new(big.Int).Exp(b.X, new(big.Int).Sub(prime, big.NewInt(2)), prime))
	result := z.Mul(z2)
	return result
}

func (z *Z) PrintRepr() {
	fmt.Println("value: ", z.X, "Prime: ", prime)
}

/*
func main() {
	var x, _ = new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)
	z := New(x, x)
	fmt.Println(BigIntToStr(z.prime), BigIntToStr(z.x))
}
*/
func BigIntToHexStr(bigInt *big.Int) string {
	return fmt.Sprintf("0x%x", bigInt)
}

func BigIntToStr(bigInt *big.Int) string {
	return fmt.Sprintf("%v", bigInt)
}
