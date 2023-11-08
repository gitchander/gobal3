package ternary

// Unary operation
// https://en.wikipedia.org/wiki/Unary_operation

type UnaryFunc func(Tri) Tri

// Binary operation
// https://en.wikipedia.org/wiki/Binary_operation

type BinaryFunc func(Tri, Tri) Tri
