# vcolor

A simple command-line tool that transforms RGB hex colors into vectors with
normalized values that can easily be used in GLSL shaders


## Installation

Make sure you have Go 1.8 or greater installed and run
`go get -u github.com/hendriklammers/vcolor`


## Usage

Simply run vcolor with a valid RGB or RGBA hex color:
```
$ vcolor FFE360
vec3(1.0, 0.89, 0.376)

$ vcolor FFE360FF
vec4(1.0, 0.89, 0.376, 1.0)
```

When the color is prefixed with a hash sign, the argument has to be
surrounded by quotes:
```
$ vcolor "#FFE360"
vec3(1.0, 0.89, 0.376)
```
