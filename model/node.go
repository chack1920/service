package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Node
type Node struct {
	ID          string   `json:"id"`
	Uid         string   `json:"uid"`
	Name        string   `json:"name"`
	Model       string   `json:"model"`
	Computed    Computed `json:"computed"`
	Device      Device   `json:"device"`
	Department  []string `json:"department"`
	Parent      []string `json:"parent"`
	Child       []string `json:"child"`
	Warning     Warning  `json:"warning"`
	Offined     bool     `json:"offlined"`
	TimeoutTime string   `json:"timeoutTime"`
}

// NodeCommand
type NodeMongo struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	Name       string               `json:"name" bson:"name"`
	Uid        string               `json:"uid" bson:"uid"`
	Model      primitive.ObjectID   `json:"model" bson:"model"`
	Parent     []primitive.ObjectID `json:"parent" bson:"parent"`
	Department []primitive.ObjectID `json:"department" bson:"department"`
	Device     DeviceMongo          `json:"device" bson:"device"`
	Custom     struct {
		IntervalTime int `json:"intervalTime"  bson:"intervalTime"`
	} `json:"custom" bson:"custom"`
	Models      []ModelMongo         `json:"models" bson:"models"`
	Child       []primitive.ObjectID `json:"child" bson:"child"`
	HasChild    bool                 `json:"hasChild" bson:"hasChild"`
	Computed    ComputedMongo        `json:"computed" bson:"computed"`
	Warning     WarningMongo         `json:"warning" bson:"warning"`
	Offined     bool                 `json:"offlined" bson:"offlined"`
	TimeoutTime interface{}          `json:"timeoutTime" bson:"timeoutTime"`
}

type Warning struct {
	RestrainWarning bool   `json:"restrainWarning"`
	Rule            []Rule `json:"rules"`
}

type Rule struct {
	ID          string                 `json:"id"`
	Level       string                 `json:"level"`
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Logic       map[string]interface{} `json:"logic"`
	Interval    float64                `json:"interval"`
	Disable     bool                   `json:"disable"`
	Handle      bool                   `json:"handle"`
	Alert       bool                   `json:"alert"`
	ExtraTags   []ExtraTagForRule      `json:"extraTags"`
	Delay       float64                `json:"delay"`
}

type WarningMongo struct {
	RestrainWarning bool        `json:"restrainWarning" bson:"restrainWarning"`
	RuleMongo       []RuleMongo `json:"rules" bson:"rules"`
}

type RuleMongo struct {
	ID          string                 `json:"id" bson:"id"`
	Level       string                 `json:"level" bson:"level"`
	Type        string                 `json:"type" bson:"type"`
	Description string                 `json:"description" bson:"description"`
	Logic       primitive.M            `json:"logic" bson:"logic"`
	Interval    interface{}            `json:"interval" bson:"interval"`
	Disable     bool                   `json:"disable"bson:"disable"`
	Handle      bool                   `json:"handle"bson:"handle"`
	Alert       bool                   `json:"alert"bson:"alert"`
	ExtraTags   []ExtraTagForRuleMongo `json:"extraTags" bson:"extraTags"`
	Delay       interface{}            `json:"delay" bson:"delay"`
}

type ComputedMongo struct {
	Auto      bool       `json:"auto" bson:"auto"`
	Tags      []TagMongo `json:"tags" bson:"tags"`
	ExtraTags []ExtraTag `json:"extraTags" bson:"extraTags"`
}

type ExtraTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ExtraTagForRule struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FieldsInWarn struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ExtraTagForRuleMongo struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type TagMongo struct {
	ID            string      `json:"id" bson:"id"`
	DataType      string      `json:"dataType" bson:"dataType"`
	Type          string      `json:"type" bson:"type"`
	Value         interface{} `json:"value" bson:"value"`
	Name          string      `json:"name" bson:"name"`
	Policy        string      `json:"policy" bson:"policy"`
	GroupName     string      `json:"groupname" bson:"groupname"`
	Unit          string      `json:"unit" bson:"unit"`
	Logic         primitive.M `json:"logic" bson:"logic"`
	StatsMethod   string      `json:"statsMethod" bson:"statsMethod"`
	StatsTag      string      `json:"statsTag" bson:"statsTag"`
	StatsInterval string      `json:"statsInterval" bson:"statsInterval"`
	Mapping       primitive.M `json:"mapping" bson:"mapping"`
	Rules         primitive.M `json:"rules" bson:"rules"`
}

type NodeIDUidForRedis struct {
	ID  string `json:"id"`
	Uid string `json:"uid"`
}