# Pokedex CLI

A command-line Pokemon adventure! Explore locations, discover Pokemon, catch them, and build your personal Pokedex using the PokéAPI.

## Features

- 🗺️ Explore Pokemon world locations
- 🔍 Discover Pokemon in different areas
- ⚾ Catch Pokemon with realistic probability mechanics
- 📖 Build and manage your personal Pokedex
- ⚡ Fast performance with built-in caching

## Installation

**Prerequisites**: Go 1.25.1 or later

```bash
# Clone and navigate to the repository
git clone github.com/mike-jacks/pokedexcli
cd pokedexcli

# Build the application
go build .

# Run it
./pokedexcli
```

Or run directly: `go run .`

## Quick Start

```
Pokedex > help              # See all commands
Pokedex > map               # Browse locations
Pokedex > explore pallet-town-area
Pokedex > catch pikachu     # Try to catch it!
Pokedex > pokedex           # View your collection
Pokedex > inspect pikachu   # See detailed stats
```

## Commands

| Command   | Usage                | Description                          |
| --------- | -------------------- | ------------------------------------ |
| `help`    | `help`               | Display all available commands       |
| `map`     | `map`                | Show next page of location areas     |
| `mapb`    | `mapb`               | Show previous page of location areas |
| `explore` | `explore <location>` | Find Pokemon in a specific location  |
| `catch`   | `catch <pokemon>`    | Attempt to catch a Pokemon           |
| `inspect` | `inspect <pokemon>`  | View stats of caught Pokemon         |
| `pokedex` | `pokedex`            | List all your caught Pokemon         |
| `exit`    | `exit`               | Quit the application                 |

### Detailed Examples

#### Exploring and Catching

```
Pokedex > explore canalave-city-area

Exploring canalave-city-area...
Found Pokemon:
tentacool
tentacruel
staryu
magikarp

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
Success chance: 78.45
magikarp was caught!
```

#### Inspecting Pokemon

```
Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  - HP: 35
  - Attack: 55
  - Defense: 40
  - Special Attack: 50
  - Special Defense: 50
  - Speed: 90
Types:
  - electric
```

## Catch Mechanics

The catch probability is based on the Pokemon's base experience:

- **Low experience** (common Pokemon): ~90% success rate
- **Medium experience**: ~50% success rate
- **High experience** (rare Pokemon): ~5% success rate

The system uses a logistic curve algorithm for smooth difficulty scaling. If a catch fails, try again!

## Project Structure

```
pokedexcli/
├── main.go              # Entry point
├── repl.go              # Interactive loop
├── command/             # Command implementations
├── config/              # Configuration
└── internal/
    ├── pokeapi/        # API client
    ├── pokecache/      # Caching system
    ├── pokedex/        # Pokemon storage
    └── types/          # Data structures
```

## Technical Notes

- **API**: Uses [PokéAPI](https://pokeapi.co/) for Pokemon data
- **Caching**: Automatic caching reduces API calls and improves performance
- **Error Handling**: Graceful handling of network issues and invalid inputs

## License

Educational project - not affiliated with Nintendo, Game Freak, or The Pokemon Company.
