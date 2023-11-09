package response

type CandidateInsert struct {
	Name     string `json:"candidate_name" `
	Email    string `json:"candidate_email" `
	UserName string `json:"candidate_username" `
	Password string `json:"candidate_password"`
}

type CandidateFetch struct {
	Name     string `json:"candidate_name" `
	Email    string `json:"candidate_email" `
	UserName string `json:"candidate_username" `
}

type FilterStruct struct {
	Name string `json:"filter_string"`
}
