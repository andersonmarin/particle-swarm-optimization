package util

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot/vg/draw"

	"github.com/andersonmarin/pso/swarm"
	"github.com/ritchie46/GOPHY/img2gif"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var imagesSavePosition []string

// SavePositionFrame cria o grafico e salva em um png
func SavePositionFrame(best []float64, birds [][]swarm.Bird) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = fmt.Sprintf("Bird position - Interaction %d", len(birds[0])-1)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.X.Max = 300
	p.X.Min = -300
	p.Y.Max = 300
	p.Y.Min = -300

	// add birds
	pts := make(plotter.XYs, len(birds))
	for i := 0; i < len(birds); i++ {
		j := len(birds[i]) - 1
		pts[i].X = birds[i][j].Position[0]
		pts[i].Y = birds[i][j].Position[1]
	}
	lpLine, lpPoints, err := plotter.NewLinePoints(pts)
	if err != nil {
		panic(err)
	}
	lpPoints.Color = color.RGBA{52, 152, 219, 255}
	lpPoints.Shape = draw.CircleGlyph{}
	lpPoints.Radius = 2
	lpLine.Color = color.Transparent

	p.Add(lpLine, lpPoints)

	// add best
	pBest := make(plotter.XYs, 1)
	pBest[0].X = best[0]
	pBest[0].Y = best[1]
	pBestLine, pBestPoints, err := plotter.NewLinePoints(pBest)
	if err != nil {
		panic(err)
	}
	pBestPoints.Color = color.RGBA{231, 76, 60, 255}
	pBestPoints.Shape = draw.CircleGlyph{}
	pBestLine.Color = color.Transparent

	p.Add(pBestLine, pBestPoints)

	filename := fmt.Sprintf(".tmp/position_%d.png", len(imagesSavePosition))

	// Save the plot to a PNG file.
	if err := p.Save(7*vg.Inch, 7*vg.Inch, filename); err != nil {
		panic(err)
	}
	imagesSavePosition = append(imagesSavePosition, filename)
}

// SavePosition salva para o gif
func SavePosition() {
	fmt.Print("Saving gif...")
	img2gif.BuildGif(&imagesSavePosition, 1, "output_position.gif")
	fmt.Println("Done")
}
