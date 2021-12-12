package model

type Task struct {
	ID               string    `json:"id"`
	AccountId        string    `json:"accountId"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	BriefDescription string    `json:"briefDescription"`
	ParentIds        []string  `json:"parentIds"`
	SuperParentIds   []string  `json:"superParentIds"`
	SharedIds        []string  `json:"sharedIds"`
	AuthorIds        []string  `json:"authorIds"`
	SuperTaskIds     []string  `json:"superTaskIds"`
	subTaskIds       []string  `json:"subTaskIds"`
	dependencyIds    []string  `json:"dependencyIds"`
	ResponsibleIds   string    `json:"responsibleIds"`
	FollowerIDS      []string  `json:"followerIds"`
	Status           string    `json:"status"`
	Importance       string    `json:"importance"`
	CreatedDate      string    `json:"createdDate"`
	UpdatedDate      string    `json:"updatedDate"`
	Dates            TaskDate  `json:"dates"`
	Scope            string    `json:"scope"`
	CustomStatusId   string    `json:"customStatusId"`
	HasAttachments   bool      `json:"hasAttachments"`
	AttachmentCount  int       `json:"attachmentCount"`
	Permalink        string    `json:"permalink"`
	Metadata         Metadatas `json:"metadata"`
	CustomFields     Fields    `json:"customFields"`
}
