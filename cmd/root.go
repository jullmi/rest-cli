/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var SERVER string
var PORT string
var data string
var username string
var password string

type User struct {
	ID        int    `json:"id"`
	Username  string `json: "username"`
	Password  string `json: "password"`
	LastLogin int64  `json: "lastLogin"`
	Admin     int    `json: "admin"`
	Active    int    `json: "active"`
}

const (
	empty = ""
	tab = "\t"
)

func (p *User) FromJSON(r io.Reader)error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *User) ToJSON(w io.Writer) error {
 e := json.NewEncoder(w)
 return e.Encode(p)
}

func SliceFromJSON(slice interface{}, r io.Reader ) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

func SliceToJSON(slice interface{}, w io.Writer ) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rest-cli",
	Short: "A REST API client",
	Long:  `A Client for a RESTful server.`,
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "username", "The username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "admin", "The password")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "{}", "JSON Record")

	rootCmd.PersistentFlags().StringVarP(&SERVER, "server", "s", "http://localhost", "RESTful server hostname")
	rootCmd.PersistentFlags().StringVarP(&PORT, "port", "P", ":1234", "Port of RESTful Server")
}
