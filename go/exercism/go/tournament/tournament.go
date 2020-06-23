package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const teamNamePadLength int = 30

type team struct {
	name    string
	wins    int
	losses  int
	draws   int
	points  int
	matches int
}

// Tally accepts a series of football match records representing a
// competition and returns a summary table of statistics for
// the competition.
func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	competition := make(map[string]team)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			return fmt.Errorf("invalid line (want 'teamA;teamB;result'): %s", line)
		}

		a, b, matchResult := fields[0], fields[1], fields[2]
		teamA, teamB := competition[a], competition[b]
		teamA.name, teamB.name = a, b
		teamA.matches++
		teamB.matches++
		switch matchResult {
		case "win":
			teamA.wins++
			teamA.points += 3
			teamB.losses++
		case "loss":
			teamA.losses++
			teamB.wins++
			teamB.points += 3
		case "draw":
			teamA.draws++
			teamA.points++
			teamB.draws++
			teamB.points++
		default:
			return fmt.Errorf("invalid match result (want 'win', 'loss', or 'draw'): %s", matchResult)
		}
		competition[a] = teamA
		competition[b] = teamB
	}
	allTeamRecords := make([]team, 0, len(competition))
	for _, v := range competition {
		allTeamRecords = append(allTeamRecords, v)
	}
	sort.Slice(allTeamRecords, func(i, j int) bool {
		if allTeamRecords[i].points == allTeamRecords[j].points {
			return allTeamRecords[i].name < allTeamRecords[j].name
		}
		return allTeamRecords[i].points > allTeamRecords[j].points
	})
	fmt.Fprintf(writer,
		"%*s | %2s | %2s | %2s | %2s | %2s\n",
		-teamNamePadLength, "Team", "MP", "W", "D", "L", "P")
	for _, t := range allTeamRecords {
		fmt.Fprintf(writer,
			"%*s | %2d | %2d | %2d | %2d | %2d\n",
			-teamNamePadLength, t.name, t.matches, t.wins, t.draws, t.losses, t.points)
	}
	return nil
}
