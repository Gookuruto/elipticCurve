package edwards

import (
	"fmt"
	"math/big"

	"github.com/Gookuruto/elipticCurve/cyclicGroup"
)

type Point struct {
	x, y *cyclicGroup.Z
}

type EdwardCurves struct {
	Group, One, Zero, d *cyclicGroup.Z
	p                   *big.Int
}

func (p *Point) ComparePoints(y *Point) bool {
	first := p.x.X.Cmp(y.x.X)
	second := p.y.X.Cmp(y.y.X)
	if first == 0 && second == 0 {
		return true
	} else {
		return false
	}
}

func NewCurve(p, d *big.Int) *EdwardCurves {
	cyclicGroup.NewGroup(p)
	edwards := new(EdwardCurves)
	edwards.p = p
	edwards.d = cyclicGroup.New(d)
	edwards.One = cyclicGroup.New(big.NewInt(1))
	edwards.Zero = cyclicGroup.New(big.NewInt(0))

	return edwards
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

// P=[Px,Py]
// -P = [Px,-Py] -> -Py mod p = p-Py
func (curve *EdwardCurves) Neg(p *Point) *Point {
	pt := new(Point)
	pt.x = p.x
	pt.y = cyclicGroup.New(new(big.Int).Sub(curve.p, p.y.X))

	return pt
}

func (curve *EdwardCurves) CreatePoint(x, y *big.Int) *Point {
	p := new(Point)
	p.x = cyclicGroup.New(x)
	p.y = cyclicGroup.New(y)
	return p
}

func (curve *EdwardCurves) IsOnCurve(p *Point) bool {
	x, y := p.x, p.y
	x1, y1 := x.X, y.X
	first := new(big.Int).Mod(new(big.Int).Add(x1.Exp(x1, big.NewInt(2), nil), y1.Exp(y1, big.NewInt(2), nil)), curve.p)
	second := new(big.Int).Mod(curve.d.X.Add(big.NewInt(1), curve.d.X.Mul(x1.Exp(x1, big.NewInt(2), nil), y1.Exp(y1, big.NewInt(2), nil))), curve.p)
	if first.Cmp(second) == 0 {
		return true

	} else {
		return false
	}

}

func (curve *EdwardCurves) Order(g *Point) *big.Int {
	basePoint := curve.CreatePoint(big.NewInt(0), big.NewInt(1))
	if !curve.IsOnCurve(g) {
		//return big.NewInt(-1)

	}
	if g.ComparePoints(basePoint) {
		//return big.NewInt(-1)
	}
	fmt.Println("p= ", curve.p)
	start := big.NewInt(2)
	end := curve.p
	fmt.Println(end)

	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {
		fmt.Println(i)
		if curve.ScalarMul(g, i).ComparePoints(basePoint) {
			return i
		}

	}
	return big.NewInt(1)

}

var one = big.NewInt(1)
