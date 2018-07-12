package main

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
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
	hash := sha1.New()
	io.WriteString(hash, str)

	// Create and iterate over slice and test even/odd
	nums := hash.Sum(nil)
	even := nums[:0]
	for _, v := range nums {
		if v%2 == 0 {
			even = append(even, 1)
		} else {
			even = append(even, 0)
		}
	}

	fmt.Printf("%d \n", even)

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

	// Write avatar to file
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, avatar)

	// Range Over even
	// if val = 1, set

	/* Create avatar
	avatar := image.NewRGBA(image.Rect(0, 0, 50, 50))
	avatar.Set(2, 3, color.RGBA{255, 0, 0, 255})

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, avatar)
	*/
}
