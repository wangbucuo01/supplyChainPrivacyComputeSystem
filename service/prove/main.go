package main

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// CubicCircuit defines a simple circuit
// x < 100  x > 0
type CubicCircuit struct {
	X frontend.Variable `gnark:"x"`       // 输入值
	Y frontend.Variable `gnark:",public"` // 最小值
	Z frontend.Variable `gnark:",public"` // 最大值
}

// Define declares the circuit constraints
// x < 100  x > 0
func (circuit *CubicCircuit) Define(api frontend.API) error {
	api.AssertIsEqual(api.Cmp(circuit.X, circuit.Y), 1)
	api.AssertIsEqual(api.Cmp(circuit.Z, circuit.X), 1)
	return nil
}

func main() {
	// compiles our circuit into a R1CS
	var circuit CubicCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)

	// groth16 zkSNARK: Setup
	pk, vk, _ := groth16.Setup(ccs)

	// witness definition
	assignment := CubicCircuit{
		X: 120,
		Y: 0,
		Z: 100,
	}
	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()

	// groth16: Prove & Verify
	proof, _ := groth16.Prove(ccs, pk, witness)
	groth16.Verify(proof, vk, publicWitness)
	fmt.Println("验证通过！")
}
