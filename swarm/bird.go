package swarm

import "math/rand"

// Bird é um pássaro
type Bird struct {
	Position     []float64
	Speed        float64
	BestPosition []float64
	BestValue    float64
}

// Fitness é a função que verifica a proximidade do objetivo
type Fitness func([]float64) float64

// Next realiza o calculo da nova posição e velocidade da partícula
func (b *Bird) Next(inertia, acceleration, globalAcceleration float64, globalBest []float64) {
	for i := 0; i < len(b.Position); i++ {
		r1 := rand.Float64()
		r2 := rand.Float64()
		b.Speed = inertia*b.Speed + acceleration*r1*(b.BestPosition[i]-b.Position[i]) + globalAcceleration*r2*(globalBest[i]-b.Position[i])
		b.Position[i] = b.Position[i] + b.Speed
	}
}

// Test verifica se a posição atual é a melhor
func (b *Bird) Test(fitness Fitness) {
	value := fitness(b.Position)
	if value > b.BestValue {
		b.BestPosition = b.Position
		b.BestValue = value
	}
}
