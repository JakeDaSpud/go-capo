package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {

	var argv []string = os.Args

	if len(argv) <= 1 {
		log.Println("Cannot read input image: No image provided")
	} else {
		var pixels [][]color.Color = load(argv[1])

		print_all_pixel_values(pixels, 8)
	}
}

func print_all_pixel_values(pixels [][]color.Color, max_value uint32) {
	x_len, y_len := len(pixels), len(pixels[0])

	for x := 0; x < x_len; x++ {
		for y := 0; y < y_len; y++ {
			current_pixel := pixels[x][y]
			fmt.Println(colour_to_string(current_pixel, max_value))
		}
	}
}

func colour_to_string(colour color.Color, bit_size uint32) string {
	r, g, b, a := colour.RGBA()
	var max_value int = 32

	if bit_size < 32 && bit_size > 0 {
		// Convert to different Bit sizing
		max_value = (1 << bit_size) - 1

		r = r * uint32(max_value) / 65535
		g = g * uint32(max_value) / 65535
		b = b * uint32(max_value) / 65535
		a = a * uint32(max_value) / 65535
	}

	alignment := len(fmt.Sprintf("%v", max_value))

	return fmt.Sprintf("R: %-*d G: %-*d B: %-*d A: %-*d", alignment, r, alignment, g, alignment, b, alignment, a)
}

func load(file_path string) (pixels [][]color.Color) {
	file, err := os.Open(file_path)
	if err != nil {
		log.Println("Cannot read file:", err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}

	// Create Pixels [][]color.Color Array
	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		var new_pixel []color.Color

		for y := 0; y < size.Y; y++ {
			new_pixel = append(new_pixel, img.At(x, y))
		}

		pixels = append(pixels, new_pixel)
	}

	// Return 2D Array of Pixels ([][]color.Color)
	return
}

func save(file_path string, pixels [][]color.Color) {
	// Create an image and set the pixels
	x_len, y_len := len(pixels), len(pixels[0])

	rect := image.Rect(0, 0, x_len, y_len)
	img := image.NewNRGBA(rect)

	for x := 0; x < x_len; x++ {
		for y := 0; y < y_len; y++ {
			img.Set(x, y, pixels[x][y])
		}
	}

	// Create a file and encode the image into it
	file, err := os.Create(file_path)
	if err != nil {
		log.Println("Cannot create file:", err)
	}

	defer file.Close()
	png.Encode(file, img)
}
