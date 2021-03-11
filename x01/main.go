package main

import (
	"log"

	log "github.com/sirupsen/logrus"

	"fmt"
	"math"

	"github.com/boltdb/bolt"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func writeToDB(h int, w int) error {

}

func main() {
	//====================================================
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//====================================================
	// normally I would populate here, but will save this for an interface

	w1 = 10
	h1 = 20

	err = writeToDB(w1, h1)
	if err != nil {
		log.Fatal("error happened...", err)
	}

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

}
