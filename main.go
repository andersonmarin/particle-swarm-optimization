package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andersonmarin/pso/swarm"
	"github.com/andersonmarin/pso/util"
	"github.com/fatih/color"
)

func main() {
	os.RemoveAll(".tmp")
	os.Mkdir(".tmp", os.ModePerm)
	rand.Seed(time.Now().Unix())

	// define as constantes do enxame
	s := swarm.Swarm{
		Inertia:                0.5,
		IndividualAcceleration: 2.05,
		GlobalAcceleration:     2.05,
		Fitness: func(p []float64) float64 {
			r := p[0] - p[1]
			if r > 123 {
				return r * -1
			}
			return r
		},
	}
	// inicializa o enxame (pássaros, dimensões e constante de multiplicação)
	s.Initialize(20, 2, 100)

	// cria o canal de comunicação para receber os dados de melhor posição de cada pássaro na execução atual
	c := make(chan []swarm.Bird)

	// executa o enxame repetidamente (numero de repetições, canal de comunicação)
	go s.Run(50, c)

	// cria a matriz para receber os dados de melhor posição dos pássaros
	history := make([][]swarm.Bird, len(s.Birds))
	for v := range c {
		for i := 0; i < len(s.Birds); i++ {
			history[i] = append(history[i], v[i])
		}

		// salva os dados no gráfico (.tmp/position_%d.png)
		util.SavePositionFrame(s.GlobalBestPosition, history)
	}

	// imprime o melhor resultado de todos os pássaros e destaca o melhor
	fmt.Print("[=== RESULT ===]\n\n")
	for i, b := range s.Birds {
		msg := fmt.Sprintf("[%02d] = %f \t %v\n", i+1, b.BestValue, b.BestPosition)
		if b.BestValue == s.GlobalBestValue {
			color.Yellow(msg)
		} else {
			color.White(msg)
		}
	}

	fmt.Println()

	// salva os dados no gráfico (output_fitness.png)
	util.SaveFitness(history)
	// salva os dados no gráfico (output_position.gif)
	util.SavePosition()

	os.RemoveAll(".tmp")
}
