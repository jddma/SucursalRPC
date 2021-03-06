package model

import (
	"encoding/json"
	"log"
	"net/rpc"
	"os/exec"
	"strconv"
)

type ClientRPC struct {
	host string
	connection *rpc.Client
	token string
}

func (c *ClientRPC) cesar(text string) string {

	var bash *exec.Cmd

	bash = exec.Command("python3", "cesar_ext.py", "c", "5", text)
	out, _ := bash.CombinedOutput()
	return string(out)

}

func (c *ClientRPC) AddAccount(document string, balance int, name string) bool {

	var result *bool

	argsMap := map[string]string{
		"document": document,
		"balance": strconv.Itoa(balance),
		"name": name,
	}

	argsBytes, _ := json.Marshal(argsMap)
	args := c.cesar(string(argsBytes))

	err := c.connection.Call("Central.AddAccount", args, &result)
	if err != nil {
		log.Fatal("Call error: ", err)
	}

	return *result

}

//Método para retirar dinero de una cuenta bancaria
func (c *ClientRPC) Withdrawals(document string, mountToRemove string) bool {

	var result *bool

	argsMap := map[string]string{
		"document": document,
		"mountToRemove": mountToRemove,
	}

	argsBytes, _ := json.Marshal(argsMap)

	err := c.connection.Call("Central.Withdrawals", c.cesar(string(argsBytes)), &result)
	if err != nil {
		log.Fatal("Call error: ", err)
	}

	return *result

}

//Método para consignar dinero
func (c *ClientRPC) AddMoney(document string, mountToAdd string) bool {

	var result *bool

	argsMap := map[string]string{
		"document": document,
		"mountToAdd": mountToAdd,
	}

	argsBytes, _ := json.Marshal(argsMap)

	err := c.connection.Call("Central.AddMoney", c.cesar(string(argsBytes)), &result)
	if err != nil {
		log.Fatal("Call error: ", err)
	}

	return *result

}

//Método para consumir el servicio de modificar una cuenta
func (c *ClientRPC) ModifyAccount(document string, newDocument string) bool  {

	var result *bool

	argsMap := map[string]string{
		"document": document,
		"newDocument": newDocument,
	}

	argsBytes, _ := json.Marshal(argsMap)

	err := c.connection.Call("Central.ModifyAccount", c.cesar(string(argsBytes)), &result)
	if err != nil {
		log.Fatal("Call error: ", err)
	}

	return *result

}

//Método para consumir el método remoto para obtener el sald de una cuenta
func (c *ClientRPC) GetBalance(document string) int  {

	var balance *int

	argsMap := map[string]string{
		"document": document,
	}

	argsBytes, _ := json.Marshal(argsMap)

	err := c.connection.Call("Central.GetBalance", c.cesar(string(argsBytes)), &balance)
	if err != nil {
		log.Fatal("Call error: ", err)
	}

	return *balance

}

func (c *ClientRPC) DeleteAccount(document string) bool {

	var result *bool

	argsMap := map[string]string{
		"document": document,
	}

	argsBytes, _ := json.Marshal(argsMap)

	err := c.connection.Call("Central.DeleteAccount", c.cesar(string(argsBytes)), &result)
	if err != nil {
		log.Fatal("Call error: ", err)
	}

	return *result

}

//Método para iciar y validar el cliente de RPC
func (c *ClientRPC) StartClient()  {

	//Establecer la oonexión
	var err error
	c.connection, err = rpc.DialHTTP("tcp", c.host)
	if err != nil {
		log.Fatal("Dialing error: ", err)
	}

	var branchIsValid *bool

	//Solicitar la validación usando JWT
	err = c.connection.Call("Central.ValidateBranch", c.token, &branchIsValid)
	if err != nil {
		log.Fatal("Call error: ", err)
	}

	if ! *branchIsValid {
		log.Fatal("No autorizado")
	}

}

func NewClientRPC(host string, token string) *ClientRPC {

	return &ClientRPC{
		host: host,
		token: token,
	}

}
