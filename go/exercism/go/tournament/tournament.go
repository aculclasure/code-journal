package tournament

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

type match struct {
	team1  string
	team2  string
	result string
}

type team struct {
	name   string
	wins   int
	losses int
	draws  int
}

func (t *team) getPoints() int { return (3 * t.wins) + t.draws }

type byPoints []*team

func (a byPoints) Len() int { return len(a) }

func (a byPoints) Less(i, j int) bool {
	if a[i].getPoints() < a[j].getPoints() {
		return true
	}
	if a[i].getPoints() > a[j].getPoints() {
		return false
	}
	return a[i].name > a[j].name
}

func (a byPoints) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type competition map[string]*team

func (c competition) addMatchResult(team1, team2, result string) {
	t1 := &team{name: team1}
	t2 := &team{name: team2}

	switch result {
	case "win":
		t1.wins++
		t2.losses++
	case "loss":
		t1.losses++
		t2.wins++
	default:
		t1.draws++
		t2.draws++
	}
	for _, team := range []*team{t1, t2} {
		if teamRecord, ok := c[team.name]; ok {
			teamRecord.wins += team.wins
			teamRecord.losses += team.losses
			teamRecord.draws += team.draws
		} else {
			c[team.name] = team
		}
	}
}

func (c competition) addMatchRecordsFromReader(r io.Reader) {
	bufReader := bufio.NewReader(r)

	for line, err := bufReader.ReadString('\n'); err == nil; {
		if m, err := getMatchData(getTrimmedFields(line)); err == nil {
			c.addMatchResult(m.team1, m.team2, m.result)
		} else {
			log.Printf("got error processing line \"%s\": %s", line, err.Error())
		}
		line, err = bufReader.ReadString('\n')
	}
}

func (c competition) getSortedTeamRecords() []*team {
	allRecords := make([]*team, 0, len(c))

	for _, v := range c {
		allRecords = append(allRecords, v)
	}
	sort.Sort(sort.Reverse(byPoints(allRecords)))
	return allRecords
}

func getTrimmedFields(line string) []string {
	fields := strings.Split(line, ";")
	var trimmedFields []string

	for _, field := range fields {
		trimmedFields = append(trimmedFields, strings.TrimSpace(field))
	}
	return trimmedFields
}

func getMatchData(matchFields []string) (*match, error) {
	if len(matchFields) != 3 {
		return nil, fmt.Errorf("a match record line should have at least 3 fields")
	}

	if matchFields[0] == matchFields[1] {
		return nil, fmt.Errorf("a match record should be for 2 different teams")
	}

	if matchFields[2] != "win" && matchFields[2] != "loss" && matchFields[2] != "draw" {
		return nil, fmt.Errorf("a match record should end with win, loss, or draw")
	}

	return &match{
		team1:  matchFields[0],
		team2:  matchFields[1],
		result: matchFields[2],
	}, nil
}

func writeResults(w io.Writer, sortedResults []*team) error {
	pad := getLongestTeamNameLength(sortedResults) + 8

	_, err := w.Write([]byte(fmt.Sprintf("%*s|%3s |%3s |%3s |%3s |%3s\n",
		-pad, "Team", "MP", "W", "D", "L", "P")))
	if err != nil {
		return err
	}

	for _, r := range sortedResults {
		_, err = w.Write([]byte(fmt.Sprintf("%*s|%3d |%3d |%3d |%3d |%3d\n",
			-pad,
			r.name,
			r.wins+r.losses+r.draws,
			r.wins,
			r.draws,
			r.losses,
			r.getPoints())))
		if err != nil {
			return err
		}
	}
	return nil
}

func getLongestTeamNameLength(teams []*team) int {
	max := 0

	for _, t := range teams {
		if len(t.name) > max {
			max = len(t.name)
		}
	}
	return max
}

// Tally accepts a series of football match records representing a
// competition and returns a summary table of statistics for
// the competition.
func Tally(reader io.Reader, writer io.Writer) error {
	c := competition{}
	c.addMatchRecordsFromReader(reader)
	sortedResults := c.getSortedTeamRecords()
	if len(sortedResults) == 0 {
		return fmt.Errorf("competition should have records of matches played")
	}
	return writeResults(writer, sortedResults)
}
