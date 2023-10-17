package fakeData

import (
	"fmt"
	"github.com/google/uuid"
)

type FakeDemandCandidatureStatus struct {
	ID          uuid.UUID
	Name        string
	Description string
	Code        string
}

const (
	fakeDemandCandidatureStatusNamePrefix        = "DCS name"
	fakeDemandCandidatureStatusDescriptionPrefix = "DCS description"
	fakeDemandCandidatureStatusCodePrefix        = "DCS"
)

func NewFakeDemandCandidatureStatus() FakeDemandCandidatureStatus {
	return FakeDemandCandidatureStatus{}
}

func (this FakeDemandCandidatureStatus) WithID() FakeDemandCandidatureStatus {
	this.ID = uuid.New()
	return this
}

func (this FakeDemandCandidatureStatus) WithName(index int) FakeDemandCandidatureStatus {
	this.Name = fmt.Sprintf("%s %d", fakeDemandCandidatureStatusNamePrefix, index)
	return this
}

func (this FakeDemandCandidatureStatus) WithDescription(index int) FakeDemandCandidatureStatus {
	this.Description = fmt.Sprintf("%s %d", fakeDemandCandidatureStatusDescriptionPrefix, index)
	return this
}

func (this FakeDemandCandidatureStatus) WithCode(index int) FakeDemandCandidatureStatus {
	this.Code = fmt.Sprintf("%s%v", fakeDemandCandidatureStatusCodePrefix, index)
	return this
}

func GenerateFakeDemandCandidatureStatus(index int) FakeDemandCandidatureStatus {
	return NewFakeDemandCandidatureStatus().WithID().WithName(index).WithDescription(index).WithDescription(index).
		WithCode(index)
}

func GenerateFakeDemandCandidatureStatuses(count int) []FakeDemandCandidatureStatus {
	fakeDemandCandidatureStatuss := []FakeDemandCandidatureStatus{}
	for index := 0; index < count; index++ {
		fakeDemandCandidatureStatus := GenerateFakeDemandCandidatureStatus(index)
		fakeDemandCandidatureStatuss = append(fakeDemandCandidatureStatuss, fakeDemandCandidatureStatus)
	}
	return fakeDemandCandidatureStatuss
}
