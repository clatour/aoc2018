package main

import (
	"fmt"
	"io/ioutil"
)

// Day 5

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Strip trailing newline
	f = f[:len(f)-1]

	// Part 2: preprocess the bytes, stripping types
	// ex. strip all A/a, perrform reaction, measure length

	// For all uppercase chars
	for i := 90; i >= 65; i-- {
		// Create a new slice
		s := make([]byte, len(f))
		copy(s, f)
		for j := len(s) - 1; j >= 0; j-- {
			b := int(s[j])
			if b == i || b == i+32 {
				s = append(s[:j], s[j+1:]...)
			}
		}
		result := performReaction(s)

		// Manually copy/paste shortest len
		fmt.Println("Final length with", string(i)+"/"+string(i+32), "removed:", len(result))
	}
}

// Part 1: "Reaction loop"
func performReaction(f []byte) []byte {
	prev := 0x00
	deleted := false

	for {
		// Do a pass over the slice, deleting as we go
		for i := len(f) - 1; i >= 0; i-- {
			b := f[i]
			// A + 32 == a || a - 32 == A
			if int(b) == int(prev)+32 || int(b) == int(prev)-32 {
				f = append(f[:i], f[i+2:]...)
				deleted = true
				// optimization to set the prev back one, as f[i] has changed
				if !(i+1 > len(f)) {
					prev = int(f[i])
				} else {
					prev = 0x00
				}
			} else {
				prev = int(b)
			}
		}

		// If we deleted a `unit`, reset and go again
		if deleted {
			deleted = false
			prev = 0x00
		} else {
			break
		}

	}

	return f
}
