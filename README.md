# pokedex

go cli pokedex guided project with boot.dev

Please Prove the Application Configuration before building in AppConfig file

Commads Map :

"help": {description: "Displays a help message"}

"exit": {description: "Exit the Pokedex"}

"map": {description: "Displays next n locations"}

"mapb": {description: "Displays previous n localtions"}

"explore": {description: "explores choosen location"}

"catch": {description: "attempt to catch choosen pokemon"}

"inspect": {description: "print extended pokemon details"}

"pokedex": {description: "print names of caught pokemons"}


Default Configs :

{
    "PageSize": "10",
    "CacheExpIntervalSeconds": "10s",
    "LocationAreaURL": "https://pokeapi.co/api/v2/location-area/?offset=0&limit=",
    "ExploreLocationURL": "https://pokeapi.co/api/v2/location-area/",
    "PokemonDetailsURL": "https://pokeapi.co/api/v2/pokemon/"
}
