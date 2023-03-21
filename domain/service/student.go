package service

import (
	"context"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/certificate"
	"github.com/devcui/ncepu-cs-project/domain/class"
	"github.com/devcui/ncepu-cs-project/domain/familyinfo"
	"github.com/devcui/ncepu-cs-project/domain/practicalexperience"
	"github.com/devcui/ncepu-cs-project/domain/predicate"
	"github.com/devcui/ncepu-cs-project/domain/student"
)

func StudentAdd(s domain.Student) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Student.Create().
			SetUserID(s.Edges.User.ID).
			SetAge(s.Age).
			SetAvatar(s.Avatar).
			SetCode(s.Code).
			SetName(s.Name).
			Save(context.Background())
	})
	return err
}

func StudentBatchAdd(us []domain.User) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		ss := []*domain.StudentCreate{}
		for _, user := range us {
			ss = append(ss,
				client.Student.Create().
					SetAge(0).
					SetCode("").
					SetSex("").
					SetUserID(user.ID).
					SetName(user.Username).
					SetAvatar([]byte{}).
					SetNillableClassID(nil).
					SetNillableClassLeaderID(nil).
					SetNillableDepartmentID(nil).
					SetNillableMajorID(nil).
					SetNillableTutorID(nil))
		}
		return client.Student.CreateBulk(ss...).Save(context.Background())
	})
	return err
}

func StudentDelete(id int) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		ctx := context.Background()
		tx, _ := client.Tx(ctx)
		stu, e := tx.Student.Query().Where(student.ID(id)).WithCertificate().WithFamilyInfo().WithPracticalExperience().First(context.Background())
		if e != nil {
			tx.Rollback()
			return nil, e
		}
		//
		familyDeletes := []predicate.FamilyInfo{}
		for _, fi := range stu.Edges.FamilyInfo {
			familyDeletes = append(familyDeletes, familyinfo.ID(fi.ID))
		}
		tx.FamilyInfo.Delete().Where(familyDeletes...).Exec(ctx)
		//
		certificates := []predicate.Certificate{}
		for _, fi := range stu.Edges.Certificate {
			certificates = append(certificates, certificate.ID(fi.ID))
		}
		tx.Certificate.Delete().Where(certificates...).Exec(ctx)
		//
		practicalExperience := []predicate.PracticalExperience{}
		for _, fi := range stu.Edges.PracticalExperience {
			practicalExperience = append(practicalExperience, practicalexperience.ID(fi.ID))
		}
		tx.PracticalExperience.Delete().Where(practicalExperience...).Exec(ctx)
		//
		_, e = tx.Student.Delete().Where(student.ID(id)).Exec(ctx)
		if e != nil {
			tx.Rollback()
		}
		return nil, tx.Commit()
	})
	return err
}

func StudentUpdate(s domain.Student) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		class, err := client.Class.Query().Where(class.ID(s.Edges.Class.ID)).WithMajor().WithDepartment().Only(context.Background())
		if err != nil {
			return nil, err
		}
		return client.Student.Update().
			Where(student.ID(s.ID)).
			SetNillableClassID(&class.ID).
			SetNillableDepartmentID(&class.Edges.Department.ID).
			SetNillableMajorID(&class.Edges.Major.ID).
			SetAge(s.Age).
			SetAvatar([]byte{}).
			SetSex(s.Sex).
			SetCode(s.Code).
			SetName(s.Name).
			Save(context.Background())
	})
	return err
}

func StudentQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Student.Query().WithClass().WithDepartment().WithMajor().WithCertificate().WithFamilyInfo().WithPracticalExperience().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
