package main

import (
    "log"
    "database/sql"
    "time"
    "fmt"

    _ "github.com/lib/pq"
)

type DBConn struct {
    db *sql.DB
}

type Message struct {
    senderid int
    messagebody string
    created string
}

type Room struct {
    roomid int
    roomname string
}

type User struct {
    userid int
    username string
}

func OpenDBConn() DBConn {
    db, err := sql.Open("postgres", "dbname=chat user=chloekaliman host=127.0.0.1 port=5432 sslmode=disable")
    if err != nil {
        log.Fatalf("Unable to connect to database: %s", err)
    }
    dbconn := DBConn {db: db}

    return dbconn
}

// create user
func (dbconn DBConn) CreateUser(username string) error {
    _, err := dbconn.db.Exec("INSERT INTO chatschema.users (username) VALUES ($1)", username)
    if err != nil {
        log.Println(err)
        return err
    }
    log.Printf("Created user: %s", username)
    return nil
}

// create room
func (dbconn DBConn) CreateRoom(roomname string, userid int) error {
    var roomid int
    err := dbconn.db.QueryRow("INSERT INTO chatschema.rooms (roomname) VALUES ($1) RETURNING roomid", roomname).Scan(&roomid)
    if err != nil {
        log.Println(err)
        return err
    }
    log.Printf("User %d created room: %s", userid, roomname)
    
    dbconn.JoinRoom(roomid, userid)

    return nil
}

// join room
func (dbconn DBConn) JoinRoom(roomid int, userid int) error {
    _, err := dbconn.db.Exec("INSERT INTO chatschema.roommembers (roomid, userid) VALUES ($1, $2)", roomid, userid)
    if err != nil {
        log.Println(err)
        return err
    }
    log.Printf("User %d has joined room %d", userid, roomid)
    return nil
}

// send a message from a user to a room
func (dbconn DBConn) SendMessage(messagebody string, roomid int, senderid int) error {
    _, err := dbconn.db.Exec("INSERT INTO chatschema.messages (messagebody, roomid, senderid) VALUES ($1, $2, $3)", messagebody, roomid, senderid)
    if err != nil {
        log.Println(err)    
        return err  
    }
    log.Printf("User %d has sent message '%s' in room %d", senderid, messagebody, roomid)
    return nil
}

// return all messages in a room
func (dbconn DBConn) GetMessages(roomid int) ([]Message, error) {
    rows, err := dbconn.db.Query("SELECT senderid, messagebody, created FROM chatschema.messages WHERE roomid = $1", roomid)
    if err != nil {
        log.Println(err)  
        return nil, err 
    }
    defer rows.Close()
    // slice of struct Messages
    var Messages []Message
    for rows.Next() {
        var senderid int
        var messagebody string
        var created time.Time
        if err := rows.Scan(&senderid, &messagebody, &created); err != nil {
            log.Println(err)
            return nil, err
        }
        Messages = append(Messages, Message{senderid, messagebody, created.String()})
    }
    log.Println("Messages: ", Messages)
    if err := rows.Err(); err != nil {
            log.Println(err)
            return nil, err
    }
    fmt.Errorf("TESTING fmt.Errorf")
    return Messages, nil
}

// return all users that have joined a room
func (dbconn DBConn) GetUsers(roomid int) ([]User, error) {
    rows, err := dbconn.db.Query("SELECT rm.userid, u.username  FROM chatschema.roommembers rm JOIN chatschema.users u ON rm.userid = u.userid WHERE rm.roomid = $1", roomid)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    defer rows.Close()
    var RoomUsers []User
    for rows.Next() {
        var userid int
        var username string
        if err := rows.Scan(&userid, &username); err != nil {
            log.Println(err)
            return nil, err
        }
        RoomUsers = append(RoomUsers, User{userid, username})
    }
    log.Printf("Users in room %d: %v", roomid, RoomUsers)
    return RoomUsers, nil
}

// return all available rooms
func (dbconn DBConn) GetRooms() ([]Room, error) {
    rows, err := dbconn.db.Query("SELECT roomid, roomname FROM chatschema.rooms")
    if err != nil {
        log.Println(err)
        return nil, err
    }
    // defers execution of function until surrounding function returns
    defer rows.Close()
    var Rooms []Room
    for rows.Next() {
        var roomid int
        var roomname string
        if err := rows.Scan(&roomid, &roomname); err != nil {
            log.Println(err)
            return nil, err
        }
        Rooms = append(Rooms, Room{roomid, roomname})
    }
    log.Println("rooms: ", Rooms)
    if err := rows.Err(); err != nil {
            log.Println(err)
            return nil, err
    }
    return Rooms, nil
}

// clear all rows in all tables
func (dbconn DBConn) clearAllRows() error {
    userRows, err := dbconn.db.Query("DELETE FROM chatschema.users")
    if err != nil {
        log.Println(err)
        return err
    }
    log.Println("Deleted all users")
    defer userRows.Close()

    roomRows, err := dbconn.db.Query("DELETE FROM chatschema.rooms")
    if err != nil {
        log.Println(err)
        return err
    }
    log.Println("Deleted all rooms")
    defer roomRows.Close()

    roommemberRows, err := dbconn.db.Query("DELETE FROM chatschema.roommembers")
    if err != nil {
        log.Println(err)
        return err
    }
    log.Println("Deleted all room members")
    defer roommemberRows.Close()

    messageRows, err := dbconn.db.Query("DELETE FROM chatschema.messages")
    if err != nil {
        log.Println(err)
        return err
    }
    log.Println("Deleted all messages")
    defer messageRows.Close()
    return nil
}