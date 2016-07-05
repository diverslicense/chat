package main

import (
    "fmt"
    "log"
    "database/sql"

    _ "github.com/lib/pq"
)


type Message struct {
    senderid int
    messagebody string
}

type Room struct {
    roomid int
    roomname string
}

// create user
func createUser(db *sql.DB, username string) {
    _, err := db.Exec("INSERT INTO chatschema.users (username) VALUES ($1)", username)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Created user: ", username)    
}

// create room
func createRoom(db *sql.DB, roomname string, userid int) {
    _, err := db.Exec("INSERT INTO chatschema.rooms (roomname) VALUES ($1)", roomname)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User %d created room: %s", userid, roomname)
}

// join room
func joinRoom(db *sql.DB, roomid int, userid int) {
    _, err := db.Exec("INSERT INTO chatschema.roommembers (roomid, userid) VALUES ($1, $2)", roomid, userid)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User %d has joined room %d", userid, roomid)
}

// send a message from a user to a room 
func sendMessage(db *sql.DB, messagebody string, roomid int, senderid int) {
    _, err := db.Exec("INSERT INTO chatschema.messages (messagebody, roomid, senderid) VALUES ($1, $2, $3)", messagebody, roomid, senderid)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User %d has sent message '%s' in room %d", senderid, messagebody, roomid)
}

// return all messages in a room
func getMessages(db *sql.DB, roomid int) []Message {
    rows, err := db.Query("SELECT senderid, messagebody FROM chatschema.messages WHERE roomid = $1", roomid)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    // slice of struct Messages 
    var messages []Message
    for rows.Next() {
        var senderid int
        var messagebody string
        if err := rows.Scan(&senderid, &messagebody); err != nil {
            log.Fatal(err)
        }
        messages = append(messages, Message{senderid, messagebody})
    }
    fmt.Println("Messages: ", messages)
    if err := rows.Err(); err != nil {
            log.Fatal(err)
    }
    return messages
}

// return all users that have joined a room
func getUsers(db *sql.DB, roomid int) []int {
    rows, err := db.Query("SELECT userid FROM chatschema.roommembers WHERE roomid = $1", roomid)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    var roomUsers []int
    for rows.Next() {
        var userid int
        if err := rows.Scan(&userid); err != nil {
            log.Fatal(err)
        }      
        roomUsers = append(roomUsers, userid)
    }
    fmt.Printf("Users in room %d: %d", roomid, roomUsers)
    return roomUsers
}

// return all available rooms
func getRooms(db *sql.DB) []Room {
    rows, err := db.Query("SELECT roomid, roomname FROM chatschema.rooms")
    if err != nil {
        log.Fatal(err)
    }
    // defers execution of function until surrounding function returns
    defer rows.Close()
    var rooms []Room
    for rows.Next() {
        var roomid int
        var roomname string
        if err := rows.Scan(&roomid, &roomname); err != nil {
            log.Fatal(err)
        }
        rooms = append(rooms, Room{roomid, roomname})
    }
    fmt.Println("rooms: ", rooms)
    if err := rows.Err(); err != nil {
            log.Fatal(err)
    }
    return rooms
}

func main() {
    db, err := sql.Open("postgres", "dbname=chat user= host=127.0.0.1 port=5432 sslmode=disable")
    if err != nil {
        log.Fatal("Unable to connect to database: %s", err)
    }

    //getRooms(db)
    //createUser(db, "testfunctionuser")
    //createRoom(db, "testcreateroom", 1)
    //joinRoom(db, 1, 1)
    //getUsers(db, 1)
    //sendMessage(db, "I like cheese!", 1, 1)
    //getMessages(db, 1)
    
}