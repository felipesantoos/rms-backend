package fakeData

import (
	"fmt"
	"github.com/google/uuid"
)

type FakeDemandType struct {
	ID          uuid.UUID
	Name        string
	Description string
	Code        string
}

const (
	fakeDemandTypeNamePrefix        = "Demand type name"
	fakeDemandTypeDescriptionPrefix = "Demand type description"
	fakeDemandTypeCodePrefix        = "DT"
)

func NewFakeDemandType() FakeDemandType {
	return FakeDemandType{}
}

func (this FakeDemandType) WithID() FakeDemandType {
	this.ID = uuid.New()
	return this
}

func (this FakeDemandType) WithName(index int) FakeDemandType {
	this.Name = fmt.Sprintf("%s %d", fakeDemandTypeNamePrefix, index)
	return this
}

func (this FakeDemandType) WithDescription(index int) FakeDemandType {
	this.Description = fmt.Sprintf("%s %d", fakeDemandTypeDescriptionPrefix, index)
	return this
}

func (this FakeDemandType) WithCode(index int) FakeDemandType {
	this.Code = fmt.Sprintf("%s%v", fakeDemandTypeCodePrefix, index)
	return this
}

func GenerateFakeDemandType(index int) FakeDemandType {
	return NewFakeDemandType().WithID().WithName(index).WithDescription(index).WithDescription(index).WithCode(index)
}

func GenerateFakeDemandTypes(count int) []FakeDemandType {
	fakeDemandTypes := []FakeDemandType{}
	for index := 0; index < count; index++ {
		fakeDemandType := GenerateFakeDemandType(index)
		fakeDemandTypes = append(fakeDemandTypes, fakeDemandType)
	}
	return fakeDemandTypes
}
