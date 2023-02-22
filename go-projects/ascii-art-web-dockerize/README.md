#  Ascii-Art-Web-Dockerize
!! Warning: Bash script removes all dangling containers and images before building new image and subsequently stops and deletes ALL containers and images upon exiting ('exit' - command) the shell prompt withing the created Docker container.

## Descritpion ##
This programme sets up a server through which a website, upon 'GET' request, that can take, as an input, a string and banner style then outputs (through 'POST' request) its equivalent ascii art characters.
Supplied is a bash script to build a Docker image and container for the server, by means of the 'Dockerfile'.

### How To Run ###
In the terminal under the correct directory run either command:
 
`go run .` ––To run without Docker

`sudo bash docker.sh` ––To run with Docker

In your browser type 'localhost:8080' in the url address field. Type the text you wish to artify and select the banner style of your choosing. Click submit to 'POST'.

### Implementation details: Algorithm ###
The server begins by implementing certain handlers which initiate a file server and handler functions for parsing the main html page and css formatting. Upon submitting the form, a 'POST' request is sent to the server where ascii-art is generated and returned to the client in an 'http.ResponseWriter'.

#### Particles JS Background
Author : Vincent Garreau  - vincentgarreau.com
MIT license: http://opensource.org/licenses/MIT
Demo / Generator : vincentgarreau.com/particles.js
GitHub : github.com/VincentGarreau/particles.js
How to use? : Check the GitHub README
v2.0.0

#### Authors ####
written by Ville & Andre (during Gritlab fullstack develpment program, October 2022)

