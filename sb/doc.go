package sb

// Package sb means Super Bytes, the name is abbreviated to facilitate import
// paths and usage as a library. This package exists because bytes are the
// most common and core type in Bitcoin development, and it is useful to have
// a consistent way to handle them. It provides a set of types and functions
// to work with byte arrays and slices, including parsing, creating, and
// zeroing them.

// The first version of this package used generics to avoid code duplication,
// but it sucks a little bit, so it was removed. The reason is to achieve
// simplicity and clarity in the codebase, and also facilitate compiling
// optimizations and performance.
