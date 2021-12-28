package main

// Tutorials:
// https://golang.org/doc/tutorial/web-service-gin
// https://levelup.gitconnected.com/consuming-a-rest-api-using-golang-b323602ba9d8

// JSON student: { "id": 0, "name": "Gustav", "enrollment": "Dropped out", "courseworkload": 10 }

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"utils"
)

// CRUD operation as a string.
const (
	DELETE = "DELETE"
	PUT    = "PUT"
	POST   = "POST"
	GET    = "GET"
)

// A Client struct which can send request to a RESTFUL server.
type Client struct {
	ipAddress *net.TCPAddr // ipAddress is the full ip address the Client sends requests to.
	client    *http.Client // client is the http.Client used to create and send requests through.
}

func main() {
	var address = flag.String("address", "localhost", "The address of the server.")
	var port = flag.Int("port", 8080, "The port of the server.")
	flag.Parse()

	client := NewClient(*address, *port)
	client.CommandShell()
}

// CommandShell starts the command shell interface through the terminal. From here the user can make CRUD requests on the server, which response will be displayed.
func (c *Client) CommandShell() {
	scanner := bufio.NewScanner(os.Stdin)
	dispatcher := utils.NewRegexDispatcher(utils.RestPattern)

	fmt.Println("Write your request: GET, POST, PUT or DELETE")
	fmt.Print("Input: ")

	for scanner.Scan() {
		matches := dispatcher.FindAllMatches(scanner.Text())

		method := strings.ToUpper(matches["method"])
		url := matches["url"]
		body := matches["body"]

		switch method {
		case GET:
			fmt.Println(c.Get(url))
		case POST:
			fmt.Println(c.Post(url, body))
		case PUT:
			fmt.Println(c.Put(url, body))
		case DELETE:
			fmt.Println(c.Delete(url, body))
		default:
			fmt.Println("Unknown request type. Try again.")
		}

		fmt.Println()
		fmt.Print("Input: ")
	}
}

// Get dispatches a GET request to the server and returns the result.
func (c *Client) Get(url string) string {
	return c.DispatchResponse(c.CreateRequest(GET, url, ""))
}

// Post dispatches a POST request to the server and returns the result.
func (c *Client) Post(url string, body string) string {
	return c.DispatchResponse(c.CreateRequest(POST, url, body))
}

// Put dispatches a PUT request to the server and returns the result.
func (c *Client) Put(url string, body string) string {
	return c.DispatchResponse(c.CreateRequest(PUT, url, body))
}

// Delete dispatches a DELETE request to the server and returns the result.
func (c *Client) Delete(url string, body string) string {
	return c.DispatchResponse(c.CreateRequest(DELETE, url, body))
}

// DispatchResponse checks the http.Response's status code and prints the result to the terminal.
// It also unmarshal the body of the response to show the content associated with the response.
func (c *Client) DispatchResponse(response *http.Response) string {
	switch response.StatusCode {
	case http.StatusOK:
		fmt.Println("200 -> OK!")
	case http.StatusCreated:
		fmt.Println("201 -> Created!")
	case http.StatusNoContent:
		fmt.Println("204 -> No Content!")
	case http.StatusNotFound:
		fmt.Println("404 -> Not Found!")
	case http.StatusBadRequest:
		fmt.Println("400 -> Bad Request!")
	default:
		fmt.Printf("%v -> Unknown error!", response.StatusCode)
	}

	return c.UnmarshalBody(response.Body)
}

// CreateRequest creates a request and sends it to the server.
// method is the CRUD method used.
// content is the url to the resource.
// body is the body content of the request.
func (c *Client) CreateRequest(method string, content string, body string) *http.Response {
	var bodyReader io.Reader
	if len(body) > 0 {
		bodyReader = strings.NewReader(body)
	}

	req, err := http.NewRequest(method, "http://"+c.ipAddress.String()+"/"+content, bodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Printf("Error making request %v with %v. :: %v", method, content, err)
		return nil
	}

	res, err := c.client.Do(req)
	if err != nil {
		log.Printf("Error sending request %v with %v. :: %v", method, content, err)
		return nil
	}

	return res
}

// UnmarshalBody unmarshal the body to a string of some http.Response.
func (c *Client) UnmarshalBody(body io.ReadCloser) string {
	bodyBytes, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		log.Printf("Error reading body. :: %v", err)
		return ""
	}

	return string(bodyBytes)
}

// NewClient creates and returns a new Client which will send requests to the ip address made up of the specified address and port.
func NewClient(address string, port int) *Client {
	ipAddress, err := net.ResolveTCPAddr("tcp", address+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Could not resolve ip address %v:%v :: %v", address, port, err)
	}

	return &Client{
		client:    &http.Client{},
		ipAddress: ipAddress,
	}
}
