package main

import (
	"log"
)

// Data Fragmentation and Reconstruction Challenge
func main() {
	simpleHash := NewHash()
	dataReconstructor := NewDataReconstructor(simpleHash)

	fragments := Fragments{
		1: {Data: "Hello", Hash: simpleHash.SimpleHash(dataReconstructor.NormalizeData(1, "Hello"))},
		2: {Data: "World", Hash: simpleHash.SimpleHash(dataReconstructor.NormalizeData(2, "World"))},
		3: {Data: "!", Hash: simpleHash.SimpleHash(dataReconstructor.NormalizeData(3, "!"))},
	}
	out, err := dataReconstructor.ReconstructData(fragments)
	if err != nil {
		log.Fatal(err)
	}

	println(out)
}
