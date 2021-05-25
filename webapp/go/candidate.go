package main

import "errors"

// Candidate Model
type Candidate struct {
	ID             int
	Name           string
	PoliticalParty string
	Sex            string
}

// CandidateElectionResult type
type CandidateElectionResult struct {
	ID             int
	Name           string
	PoliticalParty string
	Sex            string
	VoteCount      int
}

// PartyElectionResult type
type PartyElectionResult struct {
	PoliticalParty string
	VoteCount      int
}

var allCandidatesCache []Candidate
var idToCandidatesCache map[int]Candidate
var nameToCandidatesCache map[string]Candidate

func fetchAllCandidates() {
	idToCandidatesCache = map[int]Candidate{}
	nameToCandidatesCache = map[string]Candidate{}

	rows, err := db.Query("SELECT * FROM candidates")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		c := Candidate{}
		err = rows.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex)
		if err != nil {
			panic(err.Error())
		}
		allCandidatesCache = append(allCandidatesCache, c)
		idToCandidatesCache[c.ID] = c
		nameToCandidatesCache[c.Name] = c
	}
}

func getAllCandidate() []Candidate {
	return allCandidatesCache
}

func getCandidate(candidateID int) (Candidate, error) {
	return idToCandidatesCache[candidateID], nil
}

func getCandidateByName(name string) (Candidate, error) {
	if _, ok := nameToCandidatesCache[name]; !ok {
		return Candidate{}, errors.New("There is not candidate")
	}

	return nameToCandidatesCache[name], nil
}

func getAllPartyName() (partyNames []string) {
	rows, err := db.Query("SELECT political_party FROM candidates GROUP BY political_party")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		partyNames = append(partyNames, name)
	}
	return
}

func getCandidatesByPoliticalParty(party string) (candidates []Candidate) {
	rows, err := db.Query("SELECT * FROM candidates WHERE political_party = ?", party)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		c := Candidate{}
		err = rows.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex)
		if err != nil {
			panic(err.Error())
		}
		candidates = append(candidates, c)
	}
	return
}

func getElectionResult() (result []CandidateElectionResult) {
	rows, err := db.Query(`
		SELECT c.id, c.name, c.political_party, c.sex, IFNULL(v.count, 0)
		FROM candidates AS c
		LEFT OUTER JOIN
	  	(SELECT candidate_id, SUM(votes_num) AS count
	  	FROM votes_sum
	  	GROUP BY candidate_id) AS v
		ON c.id = v.candidate_id
		ORDER BY v.count DESC`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		r := CandidateElectionResult{}
		err = rows.Scan(&r.ID, &r.Name, &r.PoliticalParty, &r.Sex, &r.VoteCount)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, r)
	}
	return
}
