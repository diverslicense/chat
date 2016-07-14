package main

import (
    "fmt"
)

func main() {
    fmt.Println("Testing database layer...")
    //OpenDBConn().CreateUser("testuser123")
    OpenDBConn().CreateRoom("cheerios!", 8)
    //OpenDBConn().JoinRoom(8, 8)
    //OpenDBConn().SendMessage("Hello, there!!! :)", 8, 8)
    OpenDBConn().GetMessages(8)
    OpenDBConn().GetUsers(13)
    OpenDBConn().GetRooms()
 //   OpenDBConn().clearAllRows()
}