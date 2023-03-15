package service

// 专业

import (
	"context"
	"strconv"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/major"
)

func MajorAdd(d domain.Major) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Major.Create().SetMajorCategory(d.MajorCategory).SetIsMajorCategory(d.IsMajorCategory).SetEnrollmentType(d.EnrollmentType).SetSpecialType(d.SpecialType).SetName(d.Name).SetCode(d.Code).SetDescription(d.Description).Save(context.Background())
	})
	return err
}

func MajorDelete(id string) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		return client.Major.Delete().Where(major.ID(intID)).Exec(context.Background())
	})
	return err
}

func MajorUpdate(d domain.Major) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		a := client.Major.Update().Where(major.ID(d.ID)).SetMajorCategory(d.MajorCategory).SetIsMajorCategory(d.IsMajorCategory).SetEnrollmentType(d.EnrollmentType).SetSpecialType(d.SpecialType).SetName(d.Name).SetCode(d.Code).SetDescription(d.Description)
		if d.Edges.Department != nil {
			a.SetDepartmentID(d.Edges.Department.ID)
		}
		return a.Save(context.Background())
	})
	return err
}

func MajorQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Major.Query().WithDepartment().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
