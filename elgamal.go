package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	//	"github.com/Gookuruto/elipticCurve/cyclicGroup"

	"./edwards"
)

type ElGamal struct {
	curve *edwards.EdwardCurves
	g     *edwards.Point
	n     *big.Int
}

func NewElGamal(curve *edwards.EdwardCurves, g *edwards.Point) *ElGamal {
	eg := new(ElGamal)
	eg.curve = curve
	eg.g = g
	eg.n = curve.Order(g)

	return eg
}

func (el *ElGamal) generatePublicKey(privKey *big.Int) *edwards.Point {
	x := el.curve.ScalarMul(el.g, privKey)
	fmt.Println("Public key is ", x.PrintPoint)
	return x

}

func (el *ElGamal) Encrypt(p *edwards.Point, pubKey *edwards.Point, rint *big.Int) (z, y *edwards.Point) {

	/*if !el.curve.IsOnCurve(p) {
		fmt.Println("message not on curve")
		p.PrintPoint()
		return nil, nil

	}
	if !el.curve.IsOnCurve(pubKey) {
		fmt.Println("key not on curve")
		pubKey.PrintPoint()
		return nil, nil
	}*/
	cipherFirstPart := el.curve.ScalarMul(el.g, rint)
	cipherSecondPart := el.curve.AddPoints(p, el.curve.ScalarMul(pubKey, rint))
	fmt.Println("Encrypting point ", p, " with public key ", pubKey)
	fmt.Println("Encrypted to ", cipherFirstPart, cipherSecondPart)
	return cipherFirstPart, cipherSecondPart

}

func (el *ElGamal) Decrypt(cipherPartOne, cipherPartTwo *edwards.Point, privKey *big.Int) *edwards.Point {

	/*if !el.curve.IsOnCurve(cipherPartOne) || !el.curve.IsOnCurve(cipherPartTwo) {
		return nil
	}*/
	decoded := el.curve.AddPoints(cipherPartTwo, el.curve.Neg(el.curve.ScalarMul(cipherPartOne, privKey)))
	fmt.Println("Decoding ", cipherPartOne, " ", cipherPartTwo, " with private key ", privKey)
	fmt.Println("Decoded message ", decoded)
	return decoded
}

func main() {
	d := big.NewInt(5)
	p := big.NewInt(17)
	ec := edwards.NewCurve(p, d)
	g := ec.CreatePoint(big.NewInt(7), big.NewInt(12))
	/*if !ec.IsOnCurve(g) {
		fmt.Println("g is not on curve")
	}*/
	message := ec.CreatePoint(big.NewInt(12), big.NewInt(7))

	res := ec.ScalarMul(g, big.NewInt(3))
	res.PrintPoint()
	/*if !ec.IsOnCurve(message) {
		fmt.Println("message is not on curve")
	}*/
	eg := NewElGamal(ec, g)
	fmt.Println(eg.n)
	fmt.Println("Generating keys")
	temppriv, _ := rand.Int(rand.Reader, eg.n.Sub(eg.n, big.NewInt(1)))
	temptrand, _ := rand.Int(rand.Reader, eg.n.Sub(eg.n, big.NewInt(1)))
	priv_key := temppriv                 // smaller than order of point because we don't want pub_key to be base_point
	randInt := temptrand                 // smaller than order of point because we don't want pub_key to be base_point // does not have to be smaller than order of g
	pub_key := ec.ScalarMul(g, priv_key) // must be on edwards curve
	fmt.Println("Encrypting?")
	encoded1, encoded2 := eg.Encrypt(message, pub_key, randInt)
	fmt.Println("Decrypting?", encoded1, encoded2)
	decrypted := eg.Decrypt(encoded1, encoded2, priv_key)
	message.PrintPoint()
	decrypted.PrintPoint()

}
