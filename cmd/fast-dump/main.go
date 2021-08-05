package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "os"

    "github.com/rtsf-ratings/itsf-fast/parser"
)

func main() {
    var verbose bool
    var jsonOutput bool

    flag.BoolVar(&verbose, "verbose", true, "Enable verbose")
    flag.BoolVar(&jsonOutput, "json", false, "JSON output")
    flag.Parse()

    filename := flag.Arg(0)
    if filename == "" {
        log.Fatal("missed file path")
    }

    fast, err := parser.ParseFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    if jsonOutput {
        buffer, _ := json.MarshalIndent(fast, "", "  ")
        os.Stdout.Write(buffer)
        return
    }

    for _, tornament := range fast.Tournaments {
        fmt.Printf(" -- %s\n", tornament.Name)

        for _, competition := range tornament.Competitions {
            fmt.Printf(" ---- %s\n", competition.Name)

            fmt.Printf(" ------ Games\n")

            for _, match := range competition.Matches {
                fmt.Printf(" -------- %s vs %s (%s) (", match.Team1, match.Team2, match.Time)

                for _, game := range match.Games {
                    fmt.Printf("%d:%d", game.ScoreTeam1, game.ScoreTeam2)

                    if game != match.Games[len(match.Games)-1] {
                        fmt.Printf("; ")
                    }
                }

                fmt.Printf(")\n")
            }

            fmt.Printf(" ------ Ranking\n")

            for _, rank := range competition.Ranking {
                fmt.Printf(" -------- %d. %s\n", rank.Position, rank.Team)
            }
        }
    }
}
