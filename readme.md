# go Capo

***C***olour ***A***verages ***P***alette ***O***utput.

Image processing tool that collects statistics about input images, and outputs palettes.

---

## Index
- [Usage](#usage)
- [CSV Format](#csv-format)
- [Roadmap](#roadmap)

---

## Usage
Pass an image file (.PNG or .JPG) into the program:\
```capo.go [./input_location(.png .jpg)]```\
If no -o --output is provided, the output will be set to where this program is running from

Parameters

_Choose output file location_\
```-o --output [output_file_location]```

_Generate an output Palette as .PNG_\
```-p --palette```

_Generate an output Palette as .JPG_\
```-j --jpg```

_Generate an output Palette as .CSV_\
```-c --csv```

_Generate an output Palette as .CSV without the Header Row_\
```-n --csvnoheader```

_Generate an output Palette as .CSV without using Colours that are composed of only 0 and Max Values
```-i --csvignoreextremes```

_Generate Palette Swatch with the same filename_\
```-s --swatch```

_Generate Palette Swatch with the original image on the left, and the average colour on the right_\
```-x --extendedswatch```

---

## CSV Format

| Average | Mode | Variance | Standard Deviation | Minimum | First Quartile | Median | Third Quartile | Maximum |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |

---

## Roadmap

1. ~~Image input~~
2. ~~Image parsing~~
3. Pixels collected
4. Statistics:
    - Averages
    - Medians
    - Mins
    - Maxes
    - Mode
    - Variance
    - Standard Deviation
    - Quartiles
5. Output generation