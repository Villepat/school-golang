# Ascii-Art #

## Descritpion ##
This programme takes as an input a string in quotation marks  and banner style then outputs its equivalent ascii art characters.

### Usage: how to run ###
In the terminal under the correct directory run the following command to start the server:
 
 `go run .` 

In your browser type 'localhost:8080' in the url address field. Type the text you wish to artify and select the banner style of your choosing. Click submit to 'POST'.

### Implementation details: algorithm ###
The server begins by implementing certain handlers which initiate a file server and handler functions for parsing the main html page and css formatting. Upon submitting the form, a 'POST' request is sent to the server where ascii-art is generated and returned to the client in an 'http.ResponseWriter'.

#### Authors ####
written by Ville, Viktor, & Andre (during Gritlab fullstack develpment program, October 2022)
    
    
