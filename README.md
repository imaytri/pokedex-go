# pokedex-go
Create a Pokedex to show information about Pokemon species and item

My Solution Steps :
1. Check input argument and parse it into API URL.
2. Make Http request for pokemon. If result found then process JSON and print pokemon info.
3. If pokemon not found, make another http request for item. Then process JSON and print item info.
4. If item not found, print "Not found" info.

How to run this : 
1. open cli
2. go to directory /pokedex-go/build/
2. in windows run pokedex-go.exe with pokemon or item name as argument parameter. ex: pokedex-go.exe pikachu
3. in linux run ./pokedex-go also with pokemon or item name as argument parameter. ex: ./pokedex-go ability capsule

*) note :
if you got error (Permission Denied) when running in linux, you can try run "chmod +x pokedex-go" then run the application again.
