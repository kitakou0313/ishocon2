package main

import "strings"

// Vote Model
type Vote struct {
	ID          int
	UserID      int
	CandidateID int
	Keyword     string
}

func getVoteCountByCandidateID(candidateID int) (count int) {
	row := db.QueryRow("SELECT SUM(votes_num) AS count FROM votes_sum GROUP BY candidate_id HAVING candidate_id = ?", candidateID)
	row.Scan(&count)
	return
}

func getUserVotedCount(userID int) (count int) {
	row := db.QueryRow("SELECT SUM(votes_num) AS count FROM votes_sum GROUP BY user_id HAVING user_id = ?", userID)
	row.Scan(&count)
	return
}

func createVote(userID int, candidateID int, keyword string, voteNum int) {
	db.Exec("INSERT INTO votes_sum (user_id, candidate_id, keyword, votes_num) VALUES (?, ?, ?, ?)",
		userID, candidateID, keyword, voteNum)
}

func getVoiceOfSupporter(candidateIDs []int) (voices []string) {
	args := []interface{}{}
	for _, candidateID := range candidateIDs {
		args = append(args, candidateID)
	}
	rows, err := db.Query(`
    SELECT keyword
    FROM votes_sum
    WHERE candidate_id IN (`+strings.Join(strings.Split(strings.Repeat("?", len(candidateIDs)), ""), ",")+`)
    GROUP BY keyword
    ORDER BY SUM(votes_num) DESC
    LIMIT 10`, args...)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var keyword string
		err = rows.Scan(&keyword)
		if err != nil {
			panic(err.Error())
		}
		voices = append(voices, keyword)
	}
	return
}
