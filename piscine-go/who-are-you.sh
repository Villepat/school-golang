#! bin/bash
// in this excercise we use curl to scrape data from an url
// after this we pipe it to jq to obtain specific data
//with jq we have to yet again specify it to retrieve only the name 
curl https://01.gritlab.ax/assets/superhero/all.json | jq ' .[] | select(.id == 70) | {name} | join (" ")'

