package model

import "encoding/json"

type CreateTicket struct {
	Title            string            `json:"title"`
	Description      *string           `json:"description"`
	Status           *string           `json:"status"`
	СustomStatus     *string           `json:"customStatus"`
	Importance       *string           `json:"importance"`
	Dates            *TaskDate         `json:"dates"`
	BillingType      string            `json:"billingType"`
	Shareds          []string          `json:"shareds"`
	Parents          []string          `json:"parents"`
	Responsibles     []string          `json:"responsibles"`
	Followers        []string          `json:"followers"`
	Follow           bool              `json:"follow"`
	PriorityBefore   *string           `json:"priorityBefore"`
	PriorityAfter    *string           `json:"priorityAfter"`
	SuperTasks       *string           `json:"superTasks"`
	Metadata         Metadatas         `json:"metadata"`
	CustomFields     Fields            `json:"customFields"`
	EffortAllocation *EffortAllocation `json:"effortAllocation"`
}

type TaskDate struct {
	Type     string  `json:"type"`
	Duration *int    `json:"duration"`
	Start    *string `json:"start"`
	Due      *string `json:"due"`
}

type Metadata struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}

type Metadatas []Metadata

func (Metadatas Metadatas) ToBytes() []byte {
	data, _ := json.Marshal(Metadatas)
	return data
}

type Field struct {
	ID    string  `json:"id"`
	Value *string `json:"value"`
}

type Fields []Field

func (Fields Fields) ToBytes() []byte {
	data, _ := json.Marshal(Fields)
	return data
}

type EffortAllocation struct {
	mode            string
	totalEffort     *int
	allocatedEffort *int
	billingType     *string
	fields          Fields
}
