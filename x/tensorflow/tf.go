package main

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"io/ioutil"
	"log"
)

func main() {
	model, err := ioutil.ReadFile("modelfile")
	if err != nil {
		log.Fatal(err)
	}
	// Construct an in-memory graph from the serialized form.
	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		log.Fatal(err)
	}
	// Create a session for inference over graph.
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
}
