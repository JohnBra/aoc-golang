package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{22, 147, 227, 255}
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func setBackground(img *image.RGBA, w, h int) {
	for y := range h {
		for x := range w {
			img.Set(x, y, color.RGBA{24, 24, 24, 255})
		}
	}
}

func simulateRobots(robots [][]int) int {
	// map has 103 rows and 101 columns / image is 101 px wide and 103 px high
	const seconds = 10000
	const variations = seconds / 100
	const rows, cols = 103, 101
	const w, h = cols * variations, rows * variations
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	setBackground(img, w, h)

	i := 1
	for imgRow := range variations {
		for imgCol := range variations {
			for _, robot := range robots {
				robot[1] = (robot[1] + robot[3] + rows) % rows // row
				robot[0] = (robot[0] + robot[2] + cols) % cols // col
				img.Set(robot[0]+imgCol*cols, robot[1]+imgRow*rows, color.RGBA{65, 255, 0, 255})
			}

			addLabel(img, imgCol*cols, 9+imgRow*rows, fmt.Sprintf("%d", i))
			i++
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		r, c := robot[1], robot[0]

		if r < rows/2 && c < cols/2 { //q1: r 0-50 c 0-49
			q1 += 1
		} else if r < rows/2 && c > cols/2 { //q2: r 0-50 c 51-101
			q2 += 1
		} else if r > rows/2 && c < cols/2 { //q3: r 52-103 c 0-49
			q3 += 1
		} else if r > rows/2 && c > cols/2 { //q4 r 52-103 c 51-101
			q4 += 1
		}
	}

	file, err := os.Create("tree-search.png")
	utils.Check(err)
	defer file.Close()

	err = png.Encode(file, img)
	utils.Check(err)

	return q1 * q2 * q3 * q4
}

func main() {
	// slices of [c, r, cMod, rMod]
	robots, err := utils.GetFileContentsAsIntMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := simulateRobots(robots)
	fmt.Println("Part one res", partOneRes)
	fmt.Println("Check the generated image for part 2")
}
