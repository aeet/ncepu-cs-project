package service

import (
	"context"

	"github.com/devcui/ncepu-cs-project/domain"
	"github.com/devcui/ncepu-cs-project/domain/certificate"
)

func CertificateAdd(s domain.Certificate) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Certificate.Create().
			SetStudentID(s.Edges.Student.ID).
			SetName(s.Name).
			SetCertificateImage([]byte{}).
			SetAwardCategory(s.AwardCategory).
			SetCertificateLevel(s.CertificateLevel).
			SetCertificateType(s.CertificateType).
			SetCertificateType2(s.CertificateType2).
			SetCode(s.Code).
			SetDescription(s.Description).
			SetIssueDate(s.IssueDate).
			SetDepartment(s.Department).
			Save(context.Background())
	})
	return err
}

func CertificateDelete(id int) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Certificate.Delete().Where(certificate.ID(id)).Exec(context.Background())
	})
	return err
}

func CertificateUpdate(s domain.Certificate) error {
	_, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Certificate.Update().
			Where(certificate.ID(s.ID)).
			SetCertificateImage(s.CertificateImage).
			SetAwardCategory(s.AwardCategory).
			SetCertificateLevel(s.CertificateLevel).
			SetCertificateType(s.CertificateType).
			SetCertificateType2(s.CertificateType2).
			SetCode(s.Code).
			SetDescription(s.Description).
			SetIssueDate(s.IssueDate).
			SetDepartment(s.Department).
			Save(context.Background())
	})
	return err
}

func CertificateQuery() (interface{}, error) {
	res, err := HandleByClient(func(client *domain.Client) (interface{}, error) {
		return client.Certificate.Query().All(context.Background())
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
