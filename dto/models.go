package dto

type User struct {
	Id          string `json"id,omitempty" bson"id,omitempty"`
	DisplayName string `json:"displayname,omitempty" bson:"displayname,omitempty"`
	URL         string `json:"url,omitempty" bson:"url,omitempty"`
}

type TaskStatus struct {
	Id          string `json"id,omitempty" bson"id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	StatusGroup string `json:"statusGroup,omitempty" bson:"statusGroup,omitempty"`
}

type Task struct {
	Id        string `json"id,omitempty" bson"id,omitempty"`
	Status    string `json:"status,omitempty" bson:"status,omitempty"`
	Creator   string `json:"creator,omitempty" bson:"creator,omitempty"`
	Performer string `json:"performer,omitempty" bson:"performer,omitempty"`
}

type Task2Target struct {
	Task   string `json"task,omitempty" bson"task,omitempty"`
	Target string `json"target,omitempty" bson"target,omitempty"`
}

type Target struct {
	Id             string `json"id,omitempty" bson"id,omitempty"`
	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	ExpectedResult int    `json:"expectedResult,omitempty" bson:"expectedResult,omitempty"`
	TargetGroup    string `json:"targetGroup,omitempty" bson:"targetGroup,omitempty"`
	Creator        string `json:"creator,omitempty" bson:"creator,omitempty"`
}
