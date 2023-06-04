package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"

    "github.com/bwmarrin/discordgo"
)

//cmd line parameters
var (
  Token string
)

const KuteGo_APIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func init() {
  flag.StringVar(&Token, "t", "", "Bot Token")
  flag.Parse()
}

func main() {
  dg, err := discordgo.New("Bot " + Token)
  if err != nil {
    fmt.Println("error creating Discord sesssion, ", err)
    return
  }
  
  //todo
  //handler, identify, error handling, running message, close dg session
