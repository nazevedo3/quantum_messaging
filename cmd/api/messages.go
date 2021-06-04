package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
	"github.com/nazevedo3/quantum_messaging/internal/data"
)

// createMessageHandler stores the received message in the database
func (app *application) createMessageHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//Let the user know that there must be a body provided in the request
	if len(b) == 0 {
		fmt.Fprintf(w, "request body is empty")
		return
	}
	message := &data.Message{
		Message: string(b),
	}
	// Insert the message into the database
	err = app.models.Messages.Insert(message)
	if err != nil {
		//check to see if the error is a pq.Error
		if pgErr, ok := err.(*pq.Error); ok {
			// 23505 is the error code for "unique_violation"
			// https://github.com/lib/pq/blob/9e747ca50601fcb6c958dd89f4cb8aea3e067767/error.go#L178
			if pgErr.Code == "23505" {
				app.duplicateKeyResponse(w, r)
			}
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// use the writeJSON helper to send the response back to the user
	err = app.writeJSON(w, http.StatusCreated, clientResponse{"hash": message.Hash}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// showMessageHandler retrieves message with the given hash
func (app *application) showMessageHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve the has from the URI
	params := httprouter.ParamsFromContext(r.Context())
	hash := params.ByName("hash")

	// Get the message from the database with the provided hash
	message, err := app.models.Messages.Get(hash)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)

		}
		return
	}
	// Use the writeJSON helper to send the response back to the user
	err = app.writeJSON(w, http.StatusOK, clientResponse{"message": strings.Trim(message.Message, "\n")}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
