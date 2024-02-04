# Tsis 1
## Run the program
Run by ```go run cmd/main.go cmd/handlers.go``` command. Because, for some reason ```go run cmd/main.go``` doesn't work.

## About
This program uses [Wagner–Fischer](https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm#:~:text=The%20Wagner%E2%80%93Fischer%20algorithm%20computes,find%20the%20distance%20between%20the) algorithm of finding [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance) to find the matching films from the library.

```http://localhost:8080/films``` to get list of all films
```http://localhost:8080/films/Id/``` to get the film by id
```http://localhost:8080/films/Genre``` to search films of some genre
```http://localhost:8080/films/Name``` to get film name (spaces are allowed)
```http://localhost:8080/films/Search/{name}``` to search by name. This method finds films with simillar name or if name is misspelled Use quotation mark to specify Exact match of name (like in google search).