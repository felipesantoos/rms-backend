package fakeData

import (
	"fmt"
	"github.com/google/uuid"
)

type FakeDemandStatus struct {
	ID          uuid.UUID
	Name        string
	Description string
	Code        string
}

const (
	fakeDemandStatusNamePrefix        = "Demand status name"
	fakeDemandStatusDescriptionPrefix = "Demand status description"
	fakeDemandStatusCodePrefix        = "DS"
)

func NewFakeDemandStatus() FakeDemandStatus {
	return FakeDemandStatus{}
}

func (this FakeDemandStatus) WithID() FakeDemandStatus {
	this.ID = uuid.New()
	return this
}

func (this FakeDemandStatus) WithName(index int) FakeDemandStatus {
	this.Name = fmt.Sprintf("%s %d", fakeDemandStatusNamePrefix, index)
	return this
}

func (this FakeDemandStatus) WithDescription(index int) FakeDemandStatus {
	this.Description = fmt.Sprintf("%s %d", fakeDemandStatusDescriptionPrefix, index)
	return this
}

func (this FakeDemandStatus) WithCode(index int) FakeDemandStatus {
	this.Code = fmt.Sprintf("%s%v", fakeDemandStatusCodePrefix, index)
	return this
}

func GenerateFakeDemandStatus(index int) FakeDemandStatus {
	return NewFakeDemandStatus().WithID().WithName(index).WithDescription(index).WithDescription(index).WithCode(index)
}

func GenerateFakeDemandStatuses(count int) []FakeDemandStatus {
	fakeDemandStatuss := []FakeDemandStatus{}
	for index := 0; index < count; index++ {
		fakeDemandStatus := GenerateFakeDemandStatus(index)
		fakeDemandStatuss = append(fakeDemandStatuss, fakeDemandStatus)
	}
	return fakeDemandStatuss
}
