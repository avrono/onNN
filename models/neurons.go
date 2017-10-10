package models

import (
	"github.com/onNN/types"
	"gonum.org/v1/gonum/mat"
	"fmt"
	"math/rand"
)

/* This method implements interface Datastore */
/* As long as all methods implemented we can pass Datastore around as we choose */
func (db *DB) InsertNeuron(neuron *types.Neuron) (error) {
	session := db
	d := session.DB("onNN").C("neurons")
	err := d.Insert(neuron)
	if err != nil {
		return err //TODO: handle error - no panic
	}
	return nil
}

func (db *DB) SomeOtherMethod() error {
	return nil
}

/* This constructs a matrix representing the NN */
func (db *DB) BuildNN(name string) (string, error) {

	/* (1) Work out how many Neurons there are in this structure */

	// Generate a 6×6 matrix of random values.
	data := make([]float64, 36)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	a := mat.NewDense(6, 6, data)


	// print the output
	fmt.Printf( "%f %f %f\n", a.At(0,0), a.At(1,0), a.At(2,0) )

/* (2) Initialise a matrix to represent this structure */


	return "Hello", nil




}