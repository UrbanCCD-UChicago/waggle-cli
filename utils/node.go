package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/imroc/req"

	// Needed to connecto to SQLite database
	_ "github.com/mattn/go-sqlite3"
)

// Node wraps general information about a Waggle eanbled device
type Node struct {
	ID   string
	Name string
	Port string
	Desc string
	Loc  string
}

// String prints the node struct with line splits
func (n Node) String() string {
	return fmt.Sprintf("ID: %s\nName: %s\nPort: %s\nDesc: %s\nLoc: %s", n.ID, n.Name, n.Port, n.Desc, n.Loc)
}

// AsTableRow returns a Node struct as a string suitable to printed in tabular format
func (n *Node) AsTableRow() string {
	return fmt.Sprintf("%-16s\t%-5s\t%-5s\t%-30s\t%s", n.ID, n.Name, n.Port, n.Desc, n.Loc)
}

// AsCSVRow returns a Node struct as a string suitable to printed in CSV format
func (n *Node) AsCSVRow() string {
	return fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"", n.ID, n.Name, n.Port, n.Desc, n.Loc)
}

const dbFname = "waggle-nodes-cache.sqlite3"
const expyFname = "waggle-nodes-cache-expy.txt"

func databaseIsStale() bool {
	// get the path to the cache expiry file and read it
	fname := path.Join(os.TempDir(), expyFname)
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		os.Create(fname)
		return true
	}

	// parse the the timestamp in the file
	expy, err := time.Parse(time.RFC3339, string(content))
	if err != nil {
		return true
	}

	// get the offset between now and when the database was created and return if it's older than 5 minutes
	offset := time.Now().Sub(expy)
	return offset.Minutes() > 5
}

// LoadDatabase loads the local cache
func LoadDatabase() error {
	// get the info page that's put out on the MCS servers and store the body content
	r, err := req.Get("https://www.mcs.anl.gov/research/projects/waggle/downloads/beehive1/node-info.txt")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	body, err := r.ToString()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// get the path to the local database file; open a connection to it
	dbPath := path.Join(os.TempDir(), dbFname)
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer database.Close()

	// drop the table if it exists
	statement, _ := database.Prepare(`DROP TABLE IF EXISTS nodes;`)
	statement.Exec()

	// create a clean table
	statement, _ = database.Prepare(`CREATE TABLE nodes (id TEXT, name TEXT, port TEXT, desc TEXT, loc TEXT, alive BOOL);`)
	statement.Exec()

	// build the regexps needed to parse the body
	numRegex := regexp.MustCompile(`^[0-9]+$`)
	testingStateRegex := regexp.MustCompile(`(?i)sps|tst|testing|iuh|brt|\(t\)`)
	liveIDRegex := regexp.MustCompile(`(?i).*0001E06.*`)

	// iterator over the response body
	for _, line := range strings.Split(body, "\n") {

		// the response body is the raw output of a MySQL query. data rows start with a `|`.
		if strings.HasPrefix(line, "|") {

			// test to see if the node info meets the patterns that distinguishes between various states of operation
			alive := (!testingStateRegex.MatchString(line) && liveIDRegex.MatchString(line))

			// trim the first and last pipes; split the line on pipes
			substr := line[2 : len(line)-1]
			pieces := strings.Split(substr, "|")

			// trim the row cells' white space
			var cleanPieces []string
			for _, p := range pieces {
				cleanPieces = append(cleanPieces, strings.TrimSpace(p))
			}

			// if the row has a port number, save that row as a node record
			if numRegex.MatchString(cleanPieces[1]) {
				statement, _ = database.Prepare(`INSERT INTO nodes (name, port, id, desc, loc, alive) VALUES (?, ?, ?, ?, ?, ?);`)
				statement.Exec(cleanPieces[0], cleanPieces[1], cleanPieces[2], cleanPieces[3], cleanPieces[4], alive)
			}
		}
	}

	// timestamp the cache so we can in/validate it later
	timestamp := fmt.Sprint(time.Now().Format(time.RFC3339))
	fname := path.Join(os.TempDir(), expyFname)
	os.Truncate(fname, 0)
	_ = ioutil.WriteFile(fname, []byte(timestamp), 0666)

	return nil
}

// GetAllNodes does exactly what it sounds like: it returns a slice of Node structs
// from the local cache (database in /tmp)
func GetAllNodes() []Node {
	var nodes []Node

	if databaseIsStale() {
		LoadDatabase()
	}

	dbPath := path.Join(os.TempDir(), dbFname)
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer database.Close()

	rows, _ := database.Query(`SELECT id, name, port, desc, loc FROM nodes;`)

	for rows.Next() {
		var id string
		var name string
		var port string
		var desc string
		var loc string

		rows.Scan(&id, &name, &port, &desc, &loc)
		nodes = append(nodes, Node{ID: id, Name: name, Port: port, Desc: desc, Loc: loc})
	}

	return nodes
}

// GetLiveNodes does exactly what it sounds like: it returns a slice of Node structs
// from the local cache (database in /tmp) whose `alive` bits are true
func GetLiveNodes() []Node {
	var nodes []Node

	if databaseIsStale() {
		LoadDatabase()
	}

	dbPath := path.Join(os.TempDir(), dbFname)
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer database.Close()

	rows, _ := database.Query(`SELECT id, name, port, desc, loc FROM nodes WHERE alive = TRUE;`)

	for rows.Next() {
		var id string
		var name string
		var port string
		var desc string
		var loc string

		rows.Scan(&id, &name, &port, &desc, &loc)
		nodes = append(nodes, Node{ID: id, Name: name, Port: port, Desc: desc, Loc: loc})
	}

	return nodes
}

// GetNode does exactly what it sounds like: it returns a Node struct from the local
// cache (database in /tmp) whose `id` or `name` attrs match the param
func GetNode(idOrVsn string) (Node, error) {
	if databaseIsStale() {
		LoadDatabase()
	}

	dbPath := path.Join(os.TempDir(), dbFname)
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer database.Close()

	var row *sql.Row

	if len(idOrVsn) <= 3 {
		row = database.QueryRow(`SELECT id, name, port, desc, loc FROM nodes WHERE name = ?;`, idOrVsn)
	} else {
		idLike := fmt.Sprintf("%%%s", idOrVsn)
		row = database.QueryRow(`SELECT id, name, port, desc, loc FROM nodes WHERE (id LIKE ? OR name = ?);`, idLike, idOrVsn)
	}

	var id string
	var name string
	var port string
	var desc string
	var loc string

	switch err := row.Scan(&id, &name, &port, &desc, &loc); err {
	case sql.ErrNoRows:
		return Node{}, errors.New("Not found")
	case nil:
		return Node{ID: id, Name: name, Port: port, Desc: desc, Loc: loc}, nil
	default:
		log.Fatal(err)
		return Node{}, errors.New("Too many nodes matched")
	}
}
