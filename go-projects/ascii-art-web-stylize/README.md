# Ascii-Art-Web-Stylize#

## Descritpion ##
This programme starts a server that executes a webpage with user interface to input a text string and output ascci-art.

### Usage: how to run ###
In the terminal under the correct directory run the following command to start the server:
 
 `go run .` 

In your browser type 'localhost:8080' in the url address field. Type the text you wish to artify and select the banner style of your choosing. Click submit to 'POST'.

### Implementation details: algorithm ###
The server begins by implementing certain handlers which initiate a file server and handler functions for parsing the main html page and css formatting. Upon submitting the form, a 'POST' request is sent to the server where ascii-art is generated and returned to the client in an 'http.ResponseWriter'.

#### Particles JS Background
Author : Vincent Garreau  - vincentgarreau.com
MIT license: http://opensource.org/licenses/MIT
Demo / Generator : vincentgarreau.com/particles.js
GitHub : github.com/VincentGarreau/particles.js
How to use? : Check the GitHub README
v2.0.0

#### Authors ####
written by Ville, Viktor, & Andre (during Gritlab fullstack develpment program, October 2022)
    
    
