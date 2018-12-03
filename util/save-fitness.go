package util

import (
	"fmt"
	"image/color"

	"github.com/andersonmarin/pso/swarm"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// SaveFitness cria o grafico e salva em um png
func SaveFitness(values [][]swarm.Bird) {
	fmt.Print("Saving png...")
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Fitness"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	for i := 0; i < len(values); i++ {
		pts := make(plotter.XYs, len(values[i]))
		for j := 0; j < len(values[i]); j++ {
			pts[j].X = float64(j)
			pts[j].Y = values[i][j].BestValue
		}

		lpLine, lpPoints, err := plotter.NewLinePoints(pts)
		if err != nil {
			panic(err)
		}
		lpPoints.Color = color.Transparent
		lpLine.Color = color.RGBA{231, 76, 60, 255}
		lpLine.Width = 1

		p.Add(lpLine, lpPoints)
	}

	// Save the plot to a PNG file.
	if err := p.Save(15*vg.Inch, 7*vg.Inch, "output_fitness.png"); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
