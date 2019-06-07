package models

type CandidateResponse struct {
	Data  Candidate `json: "candidate"`
	Error APIError  `json: "error"`
}
