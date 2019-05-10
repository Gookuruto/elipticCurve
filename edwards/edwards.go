package Edwards

import (
	"math/big"

	"github.com/Gookuruto/Eliptic_Curve/cyclicGroup"
)

type Point struct {
	x, y *cyclicGroup.Z
}

type EdwardCurves struct {
	Group, One, Zero, d *cyclicGroup.Z
	p                   *big.Int
}

func NewCurve(p, d *big.Int) *EdwardCurves {
	cyclicGroup.NewGroup(p)
	edwards := new(EdwardCurves)
	edwards.p = p
	edwards.d = cyclicGroup.New(d)
	edwards.One = cyclicGroup.New(big.NewInt(1))
	edwards.Zero = cyclicGroup.New(big.NewInt(0))

	return new(EdwardCurves)
}

func (curve *EdwardCurves) AddPoints(p1, p2 *Point) *Point {
	x1, y1 := p1.x, p1.y
	x2, y2 := p2.x, p2.y
	xFirstHalf := (x1.Mul(y2)).Add(y1.Mul(x2))
	yFirstHalf := (x1.Mul(y2)).Sub(y1.Mul(x2))

	xSecondHalf := curve.One.Add(curve.d.Mul(x1.Mul(y1.Mul(x2.Mul(y2)))))
	ySecondHalf := curve.One.Sub(curve.d.Mul(x1.Mul(y1.Mul(x2.Mul(y2)))))

	x3 := xFirstHalf.TrueDiv(xSecondHalf)
	y3 := yFirstHalf.TrueDiv(ySecondHalf)
	point := Point{x3, y3}
	return &point

}

func (curve *EdwardCurves) ScalarMul(p *Point, scalar *big.Int) *Point {

	if scalar.Cmp(big.NewInt(0)) == 0 {
		pt := Point{cyclicGroup.New(big.NewInt(0)), cyclicGroup.New(big.NewInt(1))}
		return &pt
	}
	if scalar.Cmp(big.NewInt(1)) == 0 {
		return p
	}
	Q := curve.ScalarMul(p, scalar.Div(scalar, big.NewInt(2)))
	Q = curve.AddPoints(Q, Q)
	if scalar.Mod(scalar, big.NewInt(2)).Cmp(big.NewInt(0)) != 0 {
		Q = curve.AddPoints(Q, p)
	}
	return Q

}
