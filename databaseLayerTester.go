package main

import (
    "fmt"
)

func main() {
    fmt.Println("Testing database layer...")
    OpenDBConn().CreateUser("testuser123")
    OpenDBConn().CreateRoom("test room 123", 1)
    OpenDBConn().JoinRoom(1, 1)
    OpenDBConn().SendMessage("Hello, there!!! :)", 1, 1)
    OpenDBConn().GetMessages(1)
    OpenDBConn().GetUsers(1)
    OpenDBConn().GetRooms()
    OpenDBConn().clearAllRows()
}