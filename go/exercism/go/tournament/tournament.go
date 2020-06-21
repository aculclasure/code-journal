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
	matches int
}

func (t *team) getPoints() int { return (3 * t.wins) + t.draws }

// Tally accepts a series of football match records representing a
// competition and returns a summary table of statistics for
// the competition.
func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	competition := make(map[string]*team)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			return fmt.Errorf("got invalid line input: %s", line)
		}

		team1, team2, matchResult := fields[0], fields[1], fields[2]
		if _, ok := competition[team1]; !ok {
			competition[team1] = &team{name: team1}
		}
		if _, ok := competition[team2]; !ok {
			competition[team2] = &team{name: team2}
		}
		team1Record, team2Record := competition[team1], competition[team2]
		team1Record.matches++
		team2Record.matches++
		switch matchResult {
		case "win":
			team1Record.wins++
			team2Record.losses++
		case "loss":
			team1Record.losses++
			team2Record.wins++
		case "draw":
			team1Record.draws++
			team2Record.draws++
		default:
			return fmt.Errorf("got invalid match outcome: %s", matchResult)
		}
	}
	allTeamRecords := make([]*team, 0, len(competition))
	for _, v := range competition {
		allTeamRecords = append(allTeamRecords, v)
	}
	sort.Slice(allTeamRecords, func(i, j int) bool {
		if allTeamRecords[i].getPoints() == allTeamRecords[j].getPoints() {
			return allTeamRecords[i].name < allTeamRecords[j].name
		}
		return allTeamRecords[i].getPoints() > allTeamRecords[j].getPoints()
	})
	fmt.Fprintf(writer,
		"%*s | %2s | %2s | %2s | %2s | %2s\n",
		-teamNamePadLength, "Team", "MP", "W", "D", "L", "P")
	for _, t := range allTeamRecords {
		fmt.Fprintf(writer,
			"%*s | %2d | %2d | %2d | %2d | %2d\n",
			-teamNamePadLength, t.name, t.matches, t.wins, t.draws, t.losses, t.getPoints())
	}
	return nil
}
