# vec-color

A simple command-line tool to convert a RGB hex color into a vector with
normalized values that can easily be used in GLSL shaders


## Installation

Make sure you have Go 1.8 or greater installed and run
`go get -u github.com/hendriklammers/vec-color`


## Usage

Simply run vec-color with a valid rgb hex color:  
`vec-color FFE360` returns `vec3(1.0, 0.89, 0.376)`  

When the color is prefixed with a hash sign `#FFE360`, the argument has to be
surrounded by quotes:  
`vec-color "#FFE360"`
