package routes

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"github.com/onNN/types"
	"encoding/json"
	"github.com/onNN/models"
)

type Env struct {
	Db models.Datastore
}

/* /add/neuron */
func (env *Env) AddNeuron(w http.ResponseWriter, r *http.Request) {

	switch {
		case r.Method == "POST":

			// Read body
			data, err := ioutil.ReadAll(r.Body)

			neuron := new(types.Neuron)

			/* Marshall JSON data into the structure */
			err = json.Unmarshal(data, &neuron)
			if err != nil {
				fmt.Println("error:", err)
			}


			env.Db.InsertNeuron(neuron)
			// close the reader when done
			defer r.Body.Close()

			/* Marshall JSON data into the structure */
			err = json.Unmarshal(data, &neuron)
			if err != nil {
				fmt.Println("error:", err)
			}

			/* create a new structure */
			s := fmt.Sprintf("[Mongo Insert] %+v", neuron)
			mutter(s)

			// send back a JSON response
			js, err := json.Marshal(neuron)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			println(r.Body)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)

		case r.Method == "GET":
			// This is some response
			fmt.Fprintf(w, "Add Neuron: please use POST!")
	}

}

/* This does a build */
func (env *Env)Build(w http.ResponseWriter, r *http.Request) {
	switch  {
	case r.Method == "GET":
		s,_ := env.Db.BuildNN("TestNeuron")
		fmt.Fprintf(w, "%s", s)
	}

}

func HandlerICon(w http.ResponseWriter, r *http.Request) {
	mutter("Endpoint Hit: favicon")
}



func TestTemplate(w http.ResponseWriter, r *http.Request) {

}

func ServeResource(w http.ResponseWriter, req *http.Request) {
	mutter("Endpoint Hit: Server resource")
	path := "/home/avrono/workspace/bogo/src/github.com/pml/static" + req.URL.Path
	mutter(path)
	http.ServeFile(w, req, path)
}

// log
func mutter(s string) {
	fmt.Print(time.Now().Format(time.RFC850))
	fmt.Print(" - ")
	fmt.Println(s)
}