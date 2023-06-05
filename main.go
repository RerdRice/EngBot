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
  //This creates a new discord session with the token defined above
  dg, err := discordgo.New("Bot " + Token)
  if err != nil {
    fmt.Println("error creating Discord sesssion, ", err)
    return
  }
  
   // Register the messageCreate func as a callback for MessageCreate events.
   dg.AddHandler(messageCreate)  
    
   //recieve only message events
   dg.Identify.Intents = discordgo.IntentsGuildMessages
    
    err = dg.Open()
    if err != nil {
        fmt.Println("error opening connection," err)
        return
    }
    
    //chill till terminated
    fmt.Println("Bot is now running. Use CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc
    
    //close the session
    dg.Close()
    
    type Gopher struct {
        Name string `json: "name"`
    }
  //todo
  //handler, identify, error handling, running message, close dg session
