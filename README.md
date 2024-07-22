# ascii-art-web-stylize
Ascii-art-web entails creating and running a server, in which it will be possible to use a web GUI (Graphical User Interface) version of ascii-art.

## Features
- Converts text to ASCII art
- Displays the ASCII art on a HTML template on a web browser.
- Utilizes specific graphical templates for ASCII representation

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/togondol/ascii-art-web
    ```

2. Navigate to the project directory:

    ```bash
    cd ascii-art-web
    ```

## Implementation details
 A HTTP server is set up with a handler for ASCII art generation. It parses an HTML template, handles form inputs for text and banner selection, and serves static files. The asciiArtHandler processes POST requests, validates inputs, generates ASCII art using the asciiart package, and renders the results or handles errors using template rendering functions
 
## Usage
To be able to view the ASCII art representation on the web page you first need to ensure the server is up and running. 
To start the server, you run the following command:
```go
$ go run .
```
If the server starts successfully it will return a message:
```
"Starting server at port 8080"
```
Otherwise if the server is unable to start and encounters an error it will return the message:
```
"Error starting server"
```
When the server is running you go to your browser and type the link:
```
'http://localhost:8080'
```
You should see the main page where you can input text and select a banner. After submitting, you will be able to see the generted ASCII art in the specified format.

## Testing 
To run the tests present do the following:

1. Open the inner folder using this command:

```
cd ascii/
```
2. Run the test using this command:

```
go test -v
```

## Contributing

If you have suggestions for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## Author

This project was build and maintained by:

[Thadeus Ogondola](https://learn.zone01kisumu.ke/git/togondol/)

[Rodney Ochieng](https://learn.zone01kisumu.ke/git/rodnochieng)

