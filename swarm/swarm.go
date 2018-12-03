package swarm

import (
	"math/rand"
)

// Swarm representa um enxame de partículas
type Swarm struct {
	Inertia                float64
	IndividualAcceleration float64
	GlobalAcceleration     float64
	GlobalBestPosition     []float64
	GlobalBestValue        float64
	Birds                  []Bird
	Fitness                Fitness
}

// Next realiza o calculo da nova posição e velocidade da partícula
func (s *Swarm) Next() {
	for i := 0; i < len(s.Birds); i++ {
		s.Birds[i].Next(s.Inertia, s.GlobalAcceleration, s.GlobalAcceleration, s.GlobalBestPosition)
	}
}

// Test verifica se a posição atual é a melhor
func (s *Swarm) Test() {
	for i := 0; i < len(s.Birds); i++ {
		s.Birds[i].Test(s.Fitness)
		if s.Birds[i].BestValue > s.GlobalBestValue {
			s.GlobalBestValue = s.Birds[i].BestValue
			s.GlobalBestPosition = s.Birds[i].BestPosition
		}
	}
}

// Initialize cria um novo swarm
func (s *Swarm) Initialize(count, dimension int, mul float64) {
	s.GlobalBestPosition = make([]float64, dimension)
	birds := make([]Bird, count)
	for i := 0; i < count; i++ {
		birds[i].Position = make([]float64, dimension)
		birds[i].BestPosition = make([]float64, dimension)
		for j := 0; j < dimension; j++ {
			birds[i].Position[j] = rand.Float64() * mul
			birds[i].BestPosition[j] = birds[i].Position[j]
		}
		birds[i].Speed = s.GlobalAcceleration * rand.Float64()
	}
	s.Birds = birds
	s.Test()
}

// Run executa o metodo test e next repetidamente
func (s *Swarm) Run(count int, c chan []Bird) {
	out := make([]Bird, len(s.Birds))
	copy(out, s.Birds)
	c <- out

	for i := 0; i < count; i++ {
		s.Next()
		s.Test()
		out := make([]Bird, len(s.Birds))
		copy(out, s.Birds)
		c <- out
	}
	close(c)
}
