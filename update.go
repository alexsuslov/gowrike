package gowrike

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexsuslov/gowrike/model"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

var UPDATE = "/tasks/%s"

func UpdateRaw(ctx context.Context, taskId string,
	req io.ReadCloser) (res io.ReadCloser, err error) {

	URL := os.Getenv("WRIKE_URL") + fmt.Sprintf(UPDATE, taskId)

	body, _, err := Request(ctx, "PUT", URL, req, nil)
	if err != nil {
		return
	}

	return body, err
}

func Update(ctx context.Context, taskId string,
	req UpdateTicket) (ticket TasksResponse, err error) {

	values, err := req.Values()
	if err != nil {
		return
	}

	data := values.Encode()
	r := ioutil.NopCloser(strings.NewReader(data))

	body, err := UpdateRaw(ctx, taskId, r)
	if err != nil {
		return
	}
	defer body.Close()

	ticket = TasksResponse{}
	return ticket, json.NewDecoder(body).Decode(&ticket)
}

type UpdateTicket struct {
	Title              *string
	Description        *string
	Status             *string
	Importance         *string
	Dates              *model.TaskDate
	AddParents         []string
	RemoveParents      []string
	AddShareds         []string
	RemoveShareds      []string
	AddResponsibles    []string
	RemoveResponsibles []string
	AddSuperTasks      []string
	RemoveSuperTasks   []string
	AddFollowers       []string
	Follow             *bool
	PriorityBefore     *string
	PriorityAfter      *string
	Metadata           model.Metadatas
	CustomFields       model.Fields
	CustomStatus       *string
	Restore            *bool
	EffortAllocation   *model.EffortAllocation
	BillingType        *string
	Fields             []string
}

func (req UpdateTicket) Values() (values url.Values, err error) {

	//title
	if req.Title != nil {
		values.Set("title", *req.Title)
	}

	//description
	if req.Description != nil {
		values.Set("description", *req.Description)
	}

	//status
	if req.Status != nil {
		values.Set("status", *req.Status)
	}

	// importance
	if req.Importance != nil {
		values.Set("importance", *req.Importance)
	}

	// dates
	if req.Dates != nil {
		data, err := json.Marshal(*req.Dates)
		if err != nil {
			return values, err
		}
		values.Set("dates", string(data))
	}

	//AddParents
	if req.AddParents != nil {
		data, err := json.Marshal(req.AddParents)
		if err != nil {
			return values, err
		}
		values.Set("addResponsibles", string(data))
	}

	// removeResponsibles
	if req.RemoveResponsibles != nil {
		data, err := json.Marshal(req.RemoveResponsibles)
		if err != nil {
			return values, err
		}
		values.Set("removeResponsibles", string(data))
	}

	//addFollowers
	if req.AddFollowers != nil {
		data, err := json.Marshal(req.AddFollowers)
		if err != nil {
			return values, err
		}
		values.Set("addFollowers", string(data))
	}

	//follow
	if req.Follow != nil {
		Follow := "false"
		if *req.Follow {
			Follow = "true"
		}
		values.Set("follow", Follow)
	}

	//priorityBefore
	if req.PriorityBefore != nil {
		data, err := json.Marshal(req.PriorityBefore)
		if err != nil {
			return values, err
		}
		values.Set("priorityBefore", string(data))
	}

	//priorityAfter
	if req.PriorityAfter != nil {
		data, err := json.Marshal(req.PriorityAfter)
		if err != nil {
			return values, err
		}
		values.Set("priorityAfter", string(data))
	}

	//addSuperTasks
	if req.AddSuperTasks != nil {
		data, err := json.Marshal(req.AddSuperTasks)
		if err != nil {
			return values, err
		}
		values.Set("addSuperTasks", string(data))
	}

	//removeSuperTasks
	if req.RemoveSuperTasks != nil {
		data, err := json.Marshal(req.RemoveSuperTasks)
		if err != nil {
			return values, err
		}
		values.Set("removeSuperTasks", string(data))
	}

	//metadata
	if req.Metadata != nil {
		data, err := json.Marshal(req.Metadata)
		if err != nil {
			return values, err
		}
		values.Set("metadata", string(data))
	}

	//customFields
	if req.CustomFields != nil {
		data, err := json.Marshal(req.CustomFields)
		if err != nil {
			return values, err
		}
		values.Set("customFields", string(data))
	}

	//customStatus
	if req.CustomStatus != nil {
		data, err := json.Marshal(req.CustomStatus)
		if err != nil {
			return values, err
		}
		values.Set("customStatus", string(data))
	}

	//restore
	if req.Restore != nil {
		data, err := json.Marshal(req.Restore)
		if err != nil {
			return values, err
		}
		values.Set("restore", string(data))
	}

	//effortAllocation
	if req.EffortAllocation != nil {
		data, err := json.Marshal(req.EffortAllocation)
		if err != nil {
			return values, err
		}
		values.Set("effortAllocation", string(data))
	}
	//billingType
	if req.BillingType != nil {
		values.Set("billingType", *req.BillingType)
	}

	//fields
	if req.Fields != nil {
		data, err := json.Marshal(req.Fields)
		if err != nil {
			return values, err
		}
		values.Set("fields", string(data))
	}
	return values, nil
}
