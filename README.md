# Pokedex

Go CLI Pokedex guided project with boot.dev.

**Before Building:** Please configure the application in the AppConfig file.

## Commands Map:

- `help`: Displays a help message
- `exit`: Exit the Pokedex
- `map`: Displays next n locations
- `mapb`: Displays previous n locations
- `explore`: Explores chosen location
- `catch`: Attempt to catch chosen Pokemon
- `inspect`: Print extended Pokemon details
- `pokedex`: Print names of caught Pokemons

## Default Configs:

```json
{
    "PageSize": "10",
    "CacheExpIntervalSeconds": "10s",
    "LocationAreaURL": "https://pokeapi.co/api/v2/location-area/?offset=0&limit=",
    "ExploreLocationURL": "https://pokeapi.co/api/v2/location-area/",
    "PokemonDetailsURL": "https://pokeapi.co/api/v2/pokemon/"
}
