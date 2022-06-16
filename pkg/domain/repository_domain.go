package domain

type WorkerDTO struct {
	ID     SqlID
	Name   string
	Email  string
	Phone  string
	Status EntityStatus
}

type UpdateWorkerDTO struct {
	Name  string
	Phone string
}

type WorkerShiftDTO struct {
	ID       SqlID
	WorkerID SqlID
	ShiftID  SqlID
	Date     int64
	Status   EntityStatus
}

type UpdateWorkerShiftDTO struct {
	Date    int64
	ShiftID SqlID
}

type ShiftDTO struct {
	ID            SqlID
	FromTimestamp int64
	ToTimestamp   int64
	Status        EntityStatus
}
