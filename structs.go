package main
​
import "time"
​
type Resource struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
​
type BusinessHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}
​
type BlockHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}
​
type Appointment struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}
​
type Duration struct {
	Seconds int64 `json:"seconds"`
}
