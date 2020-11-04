package server

import (
	"../model"
	"encoding/json"
	"fmt"
	"net/http"
)

var clientRPC *model.ClientRPC

func Login(w http.ResponseWriter, r *http.Request) {

	var user model.Worker

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if  user.Login(cookieHandler, w) {
		fmt.Fprintf(w, "success")
	} else {
		fmt.Fprintf(w, "error")
	}

}

func AddAccount(w http.ResponseWriter, r *http.Request) {

	//Decodificar el cuerpo de la solicitud con formato JSON
	var data model.OperationData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientRPC.AddAccount(data.Document, data.Balance, data.Name)

}

//Controlador para condignar a una cuenta
func Withdrawals(w http.ResponseWriter, r *http.Request) {

	//Decodificar el cuerpo de la solicitud con formato JSON
	var data model.OperationData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientRPC.Withdrawals(data.Document, data.MountToRemove)

}

//Controlador para condignar a una cuenta
func AddMoney(w http.ResponseWriter, r *http.Request) {

	//Decodificar el cuerpo de la solicitud con formato JSON
	var data model.OperationData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientRPC.AddMoney(data.Document, data.MountToAdd)

}

//Controlador para modificar una cuenta
func ModifyAccount(w http.ResponseWriter, r *http.Request) {

	//Decodificar el cuerpo de la solicitud con formato JSON
	var data model.OperationData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientRPC.ModifyAccount(data.Document, data.NewDocument)

}

//Controlador para usar método para eliminar cuentas
func DeleteAccount (w http.ResponseWriter, r *http.Request) {

	//Decodificar el cuerpo de la solicitud con formato JSON
	var data model.OperationData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientRPC.DeleteAccount(data.Document)

}

//Controlador para retornal el saldo de una cuenta
func GetBalance(w http.ResponseWriter, r *http.Request) {

	//Decodificar el cuerpo de la solicitud con formato JSON
	var data model.OperationData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Establecer la estructura para la respuesta
	responseData := map[string]int{
		"balance": clientRPC.GetBalance(data.Document),
	}
	responseBody, _ := json.Marshal(responseData)

	//Responser la solicitud
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(responseBody))

}

//Método para iniciar el cliente RPC
func startRPCConnection(host string, token string) {

	clientRPC = model.NewClientRPC(host, token)
	clientRPC.StartClient()

}