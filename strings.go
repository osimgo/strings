// Package strings is a package that implements utility routines 
// for manipulating strings not included in the standard library
// [strings]. 
// The package's main purpose is to demonstrate how libraries work. 
package strings

// Reverse returnes a string s reversed from left to right.
func Reverse(s string) string {
	rs := ""
	for _, v := range s {
		rs = string(v) + rs
	}
	return rs
}