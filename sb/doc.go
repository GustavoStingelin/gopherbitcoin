package sb

// Package sb means Super Bytes, the name is abbreviated to facilitate import
// paths and usage as a library. This package exists because bytes are the
// most common and core type in Bitcoin development, and it is useful to have
// a consistent way to handle them. It provides a set of types and functions
// to work with byte arrays and slices, including parsing, creating, and
// zeroing them.
//
// Byte arrays are preferred over slices because they are more efficient and do
// not require garbage collection. Byte slices uses pointers to arrays, and
// pointers need garbage collection and heap allocation.
