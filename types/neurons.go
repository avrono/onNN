package types

/************************************************
*** Defines ths NN Data structures
*************************************************/

type Neuron struct {
	Type		string		`json:"type"`
	Id			int			`json:"id"`
	Name 		string		`json:"name"`
	Inputs  	[]int		`json:"inputs"`
	Outputs 	[]int		`json:"outputs"`
	Weights 	[]float32 	`json:"weights"`
	Function 	string		`json:"function"`
}
