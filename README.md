# vcolor [![Build Status](https://travis-ci.org/hendriklammers/vcolor.svg?branch=master)](https://travis-ci.org/hendriklammers/vcolor)

A simple command-line tool that transforms RGB hex colors into vectors with
normalized values that can easily be used in GLSL shaders.


## Installation

Make sure you have Go 1.8 or greater installed and run:
```
go get -u github.com/hendriklammers/vcolor
```


## Usage

Simply run vcolor with a valid RGB or RGBA hex color:
```
$ vcolor FFE360
vec3(1.0, 0.89, 0.376)

$ vcolor FFE360FF
vec4(1.0, 0.89, 0.376, 1.0)

$ vcolor "#FFE360"
vec3(1.0, 0.89, 0.376)
```

It's also possible to get a color palette by ID from
[COLOURlovers](http://www.colourlovers.com/)
```
$ vcolor -p=113451
vec3(0.169, 0.176, 0.259)
vec3(0.478, 0.49, 0.498)
vec3(0.694, 0.733, 0.812)
vec3(0.431, 0.043, 0.129)
vec3(0.608, 0.302, 0.451)
```

Get a random color palette
```
$ vcolor -rp
```

Get a random color
```
$ vcolor -r
```
