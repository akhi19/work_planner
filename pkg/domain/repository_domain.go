package domain

type WorkerDTO struct {
	ID     SqlID        `json:"id"`
	Name   string       `json:"name"`
	Email  string       `json:"email"`
	Phone  string       `json:"phone"`
	Status EntityStatus `json:"status"`
}

type UpdateWorkerDTO struct {
	Name  string
	Phone string
}

type WorkerShiftDTO struct {
	ID       SqlID        `json:"id"`
	WorkerID SqlID        `json:"worker_id"`
	ShiftID  SqlID        `json:"shift_id"`
	Date     int64        `json:"date"`
	Status   EntityStatus `json:"status"`
}

type UpdateWorkerShiftDTO struct {
	Date    int64
	ShiftID SqlID
}

type ShiftDTO struct {
	ID       SqlID        `json:"id"`
	FromTime int64        `json:"from_time"`
	ToTime   int64        `json:"to_time"`
	Status   EntityStatus `json:"status"`
}

type WorkerOccupiedDTO struct {
	ID       SqlID        `json:"id"`
	ShiftID  SqlID        `json:"shift_id"`
	WorkerID SqlID        `json:"worker_id"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Phone    string       `json:"phone"`
	Status   EntityStatus `json:"status"`
}
