package main

import (
	"fmt"
	"math/big"

	"github.com/Gookuruto/Eliptic_Curve/cyclicGroup"
)

func main() {
	cyclicGroup.NewGroup(big.NewInt(5))
	x := cyclicGroup.New(big.NewInt(14))
	y := cyclicGroup.New(big.NewInt(3))
	fmt.Println(x, y)
	z, _ := x.Add(y)

	z.PrintRepr()
	x.PrintRepr()
	y.PrintRepr()

}
