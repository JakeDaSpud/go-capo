package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type CapoImage [][]color.Color
type FlatCapoImage []color.Color

type CapoParameter string

var all_capo_parameters [16]CapoParameter = [16]CapoParameter{
	// Set the output location
	"-o",
	"--output",

	// Output a Palette (.PNG)
	"-p",
	"--palette",

	// Output a Palette (.JPG)
	"-j",
	"--jpg",

	// Output a Palette (.CSV)
	"-c",
	"--csv",

	// Output a Palette with no Header Row (.CSV)
	"-n",
	"--csvnoheader",

	// Output a Palette
	"-s",
	"--swatch",

	// Output a Copy of
	"-x",
	"--extendedswatch",
}

func main() {

	var argv []string = os.Args

	if len(argv) <= 1 {
		log.Println("Cannot read input image: No image provided")
	} else {
		var pixels CapoImage = load(argv[1])

		print_all_pixel_values(pixels, 8)
	}
}

func parse_parameters(input byte) byte {
	return 0
}

// Statistics Functions

func flatten(input CapoImage) FlatCapoImage {
	var output FlatCapoImage

	for _, capoImage := range input {
		output = append(output, capoImage...)
	}

	return output
}

func average_colours(colour1, colour2 color.Color) color.Color {
	r1, g1, b1, a1 := colour1.RGBA()
	r2, g2, b2, a2 := colour2.RGBA()

	newR, newG, newB, newA := (r1+r2)/2, (g1+g2)/2, (b1+b2)/2, (a1+a2)/2

	return color.RGBA{
		R: uint8(newR),
		G: uint8(newG),
		B: uint8(newB),
		A: uint8(newA),
	}
}

func average(input FlatCapoImage) color.Color {
	var average color.Color = input[0]

	for i := 1; i < len(input); i++ {
		average = average_colours(average, input[i])
	}

	return average
}

func rolling_average(input FlatCapoImage) (averages []color.Color) {
	averages = make([]color.Color, len(input))

	var current_average color.Color = input[0]
	averages = append(averages, current_average)

	for i := 1; i < len(input); i++ {
		current_average = average_colours(current_average, input[i])
		averages = append(averages, current_average)
	}

	return
}

func mode(input FlatCapoImage) color.Color {
	if len(input)%2 == 0 {
		var c1, c2 color.Color = input[len(input)/2], input[(len(input)/2)+1]
		return average_colours(c1, c2)
	} else {
		return input[len(input)/2]
	}
}

func variance(input FlatCapoImage) color.Color {

}

func standard_deviation(input FlatCapoImage) color.Color {

}

// Returns the (first, not a list of all) colour with the lowest combined RGBA values
func first_minimum(input FlatCapoImage) color.Color {
	var min color.Color = color.RGBA{255, 255, 255, 255}
	for _, colour := range input {
		r, g, b, a := colour.RGBA()
		min_r, min_g, min_b, min_a := colour.RGBA()

		if r+g+b+a < min_r+min_g+min_b+min_a {
			min = colour
		}
	}
	return min
}

func first_quartile(input FlatCapoImage) color.Color {

}

func median(input FlatCapoImage) color.Color {

}

func third_quartile(input FlatCapoImage) color.Color {

}

func first_maximum(input FlatCapoImage) color.Color {
	var max color.Color = color.RGBA{0, 0, 0, 0}
	for _, colour := range input {
		r, g, b, a := colour.RGBA()
		max_r, max_g, max_b, max_a := colour.RGBA()

		if r+g+b+a > max_r+max_g+max_b+max_a {
			max = colour
		}
	}
	return max
}

// Comparison Functions

// Returns 1 if colour1 is greater than colour2, -1 if lesser than, and 0 if equal
func compare_colours(colour1, colour2 color.Color) (result_r, result_g, result_b, result_a int) {
	r1, g1, b1, a1 := colour1.RGBA()
	r2, g2, b2, a2 := colour2.RGBA()

	if r1 > r2 {
		result_r = 1
	} else if r1 < r2 {
		result_r = -1
	} else {
		result_r = 0
	}

	if g1 > g2 {
		result_g = 1
	} else if g1 < g2 {
		result_g = -1
	} else {
		result_g = 0
	}

	if b1 > b2 {
		result_b = 1
	} else if b1 < b2 {
		result_b = -1
	} else {
		result_b = 0
	}

	if a1 > a2 {
		result_a = 1
	} else if a1 < a2 {
		result_a = -1
	} else {
		result_a = 0
	}

	return
}

// Formatting Functions

func print_all_pixel_values(pixels CapoImage, max_value uint32) {
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

// Save & Load Functions

func load(file_path string) (pixels CapoImage) {
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
	for y := 0; y < size.Y; y++ {
		var new_pixel []color.Color

		for x := 0; x < size.X; x++ {
			new_pixel = append(new_pixel, img.At(x, y))
		}

		pixels = append(pixels, new_pixel)
	}

	// Return 2D Array of Pixels ([][]color.Color)
	return
}

func save(file_path string, pixels CapoImage) {
	// Create an image and set the pixels
	x_len, y_len := len(pixels), len(pixels[0])

	rect := image.Rect(0, 0, x_len, y_len)
	img := image.NewNRGBA(rect) // NRGBA means Non-Alpha RGBA

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
