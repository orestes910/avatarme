package main

import (
	"crypto/rand"
	"crypto/sha256"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"math/big"
	"os"
)

func main() {
	// Capture argument to form hash
	str := os.Args[1]

	// Instansiate hash and add argument
	hash := sha256.New()
	io.WriteString(hash, str)

	// Create and iterate over slice and test even/odd
	nums := hash.Sum(nil)
	trim := nums[:25]
	even := trim[:0]
	for _, v := range trim {
		if v%2 == 0 {
			even = append(even, 1)
		} else {
			even = append(even, 0)
		}
	}

	// Generate random color
	bigR, _ := rand.Int(rand.Reader, big.NewInt(255))
	r := bigR.Uint64()
	bigG, _ := rand.Int(rand.Reader, big.NewInt(255))
	g := bigG.Uint64()
	bigB, _ := rand.Int(rand.Reader, big.NewInt(255))
	b := bigB.Uint64()
	col := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}

	// Generate avatar and color white
	avatar := image.NewRGBA(image.Rect(0, 0, 50, 50))
	draw.Draw(avatar, avatar.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// Starting point
	x := 0
	y := 0
	sp := image.Point{x, y}

	// Range over even/odd slice, add block if even
	for _, v := range even {
		if v == 1 {
			// Create 10x10 rectangle from starting point
			rect := image.Rectangle{sp, sp.Add(image.Point{10, 10})}
			draw.Draw(avatar, rect, &image.Uniform{col}, sp, draw.Src)
		}
		// Grid sequence
		if x < 50 {
			x = x + 10
		} else {
			x = 0
		}
		if x == 50 {
			x = 0
			y = y + 10
		}
		// Update starting point
		sp = image.Point{x, y}
	}

	// Write avatar to file
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, avatar)

}
