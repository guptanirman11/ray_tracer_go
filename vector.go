package main

import (
	"math"
)

/***********************
 * Vector
 ************************/

// Vector defines a vector in 3D space
type Vector struct {
	x, y, z float64
}

func (v Vector) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector) squaredLength() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vector) normalize() Vector {
	len := v.length()
	return Vector{v.x / len, v.y / len, v.z / len}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y, v.z + v2.z}
}

func (v Vector) Subtract(v2 Vector) Vector {
	return Vector{v.x - v2.x, v.y - v2.y, v.z - v2.z}
}

func (v Vector) Multiply(v2 Vector) Vector {
	return Vector{v.x * v2.x, v.y * v2.y, v.z * v2.z}
}

func Dot(v1 Vector, v2 Vector) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func Cross(v1 Vector, v2 Vector) Vector {
	return Vector{
		v1.y*v2.z - v1.z*v2.y,
		v1.z*v2.x - v1.x*v2.z,
		v1.x*v2.y - v1.y*v2.x,
	}
}
