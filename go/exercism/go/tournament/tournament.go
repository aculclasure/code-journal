package tournament

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

type teamRecord struct {
	teamName  string
	numWins   int
	numLosses int
	numDraws  int
	points    int
}

type byPoints []*teamRecord

func (a byPoints) Len() int { return len(a) }
func (a byPoints) Less(i, j int) bool {
	if a[i].points < a[j].points {
		return true
	}
	if a[i].points > a[j].points {
		return false
	}
	return a[i].teamName > a[j].teamName
}
func (a byPoints) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

var competitionRecords map[string]*teamRecord

// Tally accepts a series of football match records representing a
// competition and returns a summary table of statistics for
// the competition.
func Tally(reader io.Reader, writer io.Writer) error {
	competitionRecords = map[string]*teamRecord{}
	addRecords(reader)
	sortedResults := getSortedCompetitionResults()
	if len(sortedResults) == 0 {
		return fmt.Errorf("competition should have records of matches played")
	}
	return writeResults(writer, sortedResults)
}

func addRecords(r io.Reader) {
	bufReader := bufio.NewReader(r)

	for line, err := bufReader.ReadString('\n'); err == nil; {
		if matchData, err := getMatchData(getTrimmedFields(line)); err == nil {
			t1, t2 := getMatchRecords(matchData)
			addMatchRecord(t1)
			addMatchRecord(t2)
		} else {
			log.Printf("got error processing line \"%s\": %s", line, err.Error())
		}
		line, err = bufReader.ReadString('\n')
	}
}

func getTrimmedFields(line string) []string {
	fields := strings.Split(line, ";")
	var trimmedFields []string

	for _, field := range fields {
		trimmedFields = append(trimmedFields, strings.TrimSpace(field))
	}
	return trimmedFields
}

func getMatchData(matchFields []string) (map[string]string, error) {
	if len(matchFields) != 3 {
		return nil, fmt.Errorf("a match record line should have at least 3 fields")
	}

	if matchFields[0] == matchFields[1] {
		return nil, fmt.Errorf("a match record should be for 2 different teams")
	}

	if matchFields[2] != "win" && matchFields[2] != "loss" && matchFields[2] != "draw" {
		return nil, fmt.Errorf("a match record should end with win, loss, or draw")
	}

	return map[string]string{
		"firstTeamName":  matchFields[0],
		"secondTeamName": matchFields[1],
		"matchResult":    matchFields[2],
	}, nil
}

func addMatchRecord(record *teamRecord) {
	if r, ok := competitionRecords[record.teamName]; !ok {
		competitionRecords[record.teamName] = record
	} else {
		r.numWins += record.numWins
		r.numLosses += record.numLosses
		r.numDraws += record.numDraws
		r.points = (3 * r.numWins) + r.numDraws
	}
}

func writeResults(w io.Writer, sortedResults []*teamRecord) error {
	pad := getLongestTeamNameLength() + 8

	_, err := w.Write([]byte(fmt.Sprintf("%*s|%3s |%3s |%3s |%3s |%3s\n",
		-pad, "Team", "MP", "W", "D", "L", "P")))
	if err != nil {
		return err
	}

	for _, r := range sortedResults {
		_, err = w.Write([]byte(fmt.Sprintf("%*s|%3d |%3d |%3d |%3d |%3d\n",
			-pad,
			r.teamName,
			r.numWins+r.numLosses+r.numDraws,
			r.numWins,
			r.numDraws,
			r.numLosses,
			r.points)))
		if err != nil {
			return err
		}
	}
	return nil
}

func getSortedCompetitionResults() []*teamRecord {
	allRecords := make([]*teamRecord, 0, len(competitionRecords))

	for _, v := range competitionRecords {
		allRecords = append(allRecords, v)
	}
	sort.Sort(sort.Reverse(byPoints(allRecords)))
	return allRecords
}

func getLongestTeamNameLength() int {
	max := 0

	for teamName := range competitionRecords {
		if len(teamName) > max {
			max = len(teamName)
		}
	}
	return max
}

func getMatchRecords(recordData map[string]string) (*teamRecord, *teamRecord) {
	team1 := &teamRecord{teamName: recordData["firstTeamName"]}
	team2 := &teamRecord{teamName: recordData["secondTeamName"]}

	switch recordData["matchResult"] {
	case "win":
		team1.numWins++
		team2.numLosses++
	case "loss":
		team1.numLosses++
		team2.numWins++
	default:
		team1.numDraws++
		team2.numDraws++
	}
	for _, t := range []*teamRecord{team1, team2} {
		t.points = (3 * t.numWins) + t.numDraws
	}
	return team1, team2
}
