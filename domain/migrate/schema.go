// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthorizationsColumns holds the columns for the "authorizations" table.
	AuthorizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "client_id", Type: field.TypeString, Unique: true},
		{Name: "client_secret", Type: field.TypeString},
		{Name: "client_name", Type: field.TypeString, Unique: true},
		{Name: "grant_type", Type: field.TypeJSON},
		{Name: "scope", Type: field.TypeJSON},
		{Name: "redirect_url", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
	}
	// AuthorizationsTable holds the schema information for the "authorizations" table.
	AuthorizationsTable = &schema.Table{
		Name:       "authorizations",
		Columns:    AuthorizationsColumns,
		PrimaryKey: []*schema.Column{AuthorizationsColumns[0]},
	}
	// CampusColumns holds the columns for the "campus" table.
	CampusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
	}
	// CampusTable holds the schema information for the "campus" table.
	CampusTable = &schema.Table{
		Name:       "campus",
		Columns:    CampusColumns,
		PrimaryKey: []*schema.Column{CampusColumns[0]},
	}
	// CertificatesColumns holds the columns for the "certificates" table.
	CertificatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "department", Type: field.TypeString},
		{Name: "issue_date", Type: field.TypeString},
		{Name: "certificate_type", Type: field.TypeString},
		{Name: "certificate_level", Type: field.TypeString},
		{Name: "certificate_type2", Type: field.TypeString},
		{Name: "award_category", Type: field.TypeString},
		{Name: "certificate_image", Type: field.TypeBytes},
		{Name: "certificate_student", Type: field.TypeInt, Nullable: true},
	}
	// CertificatesTable holds the schema information for the "certificates" table.
	CertificatesTable = &schema.Table{
		Name:       "certificates",
		Columns:    CertificatesColumns,
		PrimaryKey: []*schema.Column{CertificatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "certificates_students_student",
				Columns:    []*schema.Column{CertificatesColumns[11]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ClassesColumns holds the columns for the "classes" table.
	ClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "campus_class", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "class_major", Type: field.TypeInt, Nullable: true},
		{Name: "class_department", Type: field.TypeInt, Nullable: true},
		{Name: "class_leader_class", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "major_direction_class", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "tutor_class", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ClassesTable holds the schema information for the "classes" table.
	ClassesTable = &schema.Table{
		Name:       "classes",
		Columns:    ClassesColumns,
		PrimaryKey: []*schema.Column{ClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "classes_campus_class",
				Columns:    []*schema.Column{ClassesColumns[5]},
				RefColumns: []*schema.Column{CampusColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_majors_major",
				Columns:    []*schema.Column{ClassesColumns[6]},
				RefColumns: []*schema.Column{MajorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_departments_department",
				Columns:    []*schema.Column{ClassesColumns[7]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_class_leaders_class",
				Columns:    []*schema.Column{ClassesColumns[8]},
				RefColumns: []*schema.Column{ClassLeadersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_major_directions_class",
				Columns:    []*schema.Column{ClassesColumns[9]},
				RefColumns: []*schema.Column{MajorDirectionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_tutors_class",
				Columns:    []*schema.Column{ClassesColumns[10]},
				RefColumns: []*schema.Column{TutorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ClassLeadersColumns holds the columns for the "class_leaders" table.
	ClassLeadersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ClassLeadersTable holds the schema information for the "class_leaders" table.
	ClassLeadersTable = &schema.Table{
		Name:       "class_leaders",
		Columns:    ClassLeadersColumns,
		PrimaryKey: []*schema.Column{ClassLeadersColumns[0]},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
	}
	// EducationLevelsColumns holds the columns for the "education_levels" table.
	EducationLevelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "education_level_student", Type: field.TypeInt, Nullable: true},
	}
	// EducationLevelsTable holds the schema information for the "education_levels" table.
	EducationLevelsTable = &schema.Table{
		Name:       "education_levels",
		Columns:    EducationLevelsColumns,
		PrimaryKey: []*schema.Column{EducationLevelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "education_levels_students_student",
				Columns:    []*schema.Column{EducationLevelsColumns[2]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EnrollmentStatusColumns holds the columns for the "enrollment_status" table.
	EnrollmentStatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "enrollment_status_student", Type: field.TypeInt, Nullable: true},
	}
	// EnrollmentStatusTable holds the schema information for the "enrollment_status" table.
	EnrollmentStatusTable = &schema.Table{
		Name:       "enrollment_status",
		Columns:    EnrollmentStatusColumns,
		PrimaryKey: []*schema.Column{EnrollmentStatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "enrollment_status_students_student",
				Columns:    []*schema.Column{EnrollmentStatusColumns[2]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FamilyInfosColumns holds the columns for the "family_infos" table.
	FamilyInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "relationship", Type: field.TypeString},
		{Name: "id_card", Type: field.TypeString},
		{Name: "age", Type: field.TypeString},
		{Name: "occupation", Type: field.TypeString},
		{Name: "post", Type: field.TypeString},
		{Name: "work_unit", Type: field.TypeString},
		{Name: "contact_number", Type: field.TypeString},
		{Name: "health", Type: field.TypeString},
		{Name: "family_info_student", Type: field.TypeInt, Nullable: true},
	}
	// FamilyInfosTable holds the schema information for the "family_infos" table.
	FamilyInfosTable = &schema.Table{
		Name:       "family_infos",
		Columns:    FamilyInfosColumns,
		PrimaryKey: []*schema.Column{FamilyInfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "family_infos_students_student",
				Columns:    []*schema.Column{FamilyInfosColumns[10]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MajorsColumns holds the columns for the "majors" table.
	MajorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "special_type", Type: field.TypeString},
		{Name: "enrollment_type", Type: field.TypeString},
		{Name: "is_major_category", Type: field.TypeBool},
		{Name: "major_category", Type: field.TypeString},
		{Name: "major_department", Type: field.TypeInt, Nullable: true},
	}
	// MajorsTable holds the schema information for the "majors" table.
	MajorsTable = &schema.Table{
		Name:       "majors",
		Columns:    MajorsColumns,
		PrimaryKey: []*schema.Column{MajorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "majors_departments_department",
				Columns:    []*schema.Column{MajorsColumns[8]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MajorDirectionsColumns holds the columns for the "major_directions" table.
	MajorDirectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// MajorDirectionsTable holds the schema information for the "major_directions" table.
	MajorDirectionsTable = &schema.Table{
		Name:       "major_directions",
		Columns:    MajorDirectionsColumns,
		PrimaryKey: []*schema.Column{MajorDirectionsColumns[0]},
	}
	// PracticalExperiencesColumns holds the columns for the "practical_experiences" table.
	PracticalExperiencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "unit", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeString},
		{Name: "end_time", Type: field.TypeString},
		{Name: "describe", Type: field.TypeString},
		{Name: "practical_experience_student", Type: field.TypeInt, Nullable: true},
	}
	// PracticalExperiencesTable holds the schema information for the "practical_experiences" table.
	PracticalExperiencesTable = &schema.Table{
		Name:       "practical_experiences",
		Columns:    PracticalExperiencesColumns,
		PrimaryKey: []*schema.Column{PracticalExperiencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "practical_experiences_students_student",
				Columns:    []*schema.Column{PracticalExperiencesColumns[6]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ResourcesColumns holds the columns for the "resources" table.
	ResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "resource_name", Type: field.TypeString},
		{Name: "resource_value", Type: field.TypeString},
	}
	// ResourcesTable holds the schema information for the "resources" table.
	ResourcesTable = &schema.Table{
		Name:       "resources",
		Columns:    ResourcesColumns,
		PrimaryKey: []*schema.Column{ResourcesColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_name", Type: field.TypeString, Unique: true},
		{Name: "role_value", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "sex", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeBytes},
		{Name: "class_leader_student", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "student_department", Type: field.TypeInt, Nullable: true},
		{Name: "student_major", Type: field.TypeInt, Nullable: true},
		{Name: "student_class", Type: field.TypeInt, Nullable: true},
		{Name: "tutor_student", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "students_class_leaders_student",
				Columns:    []*schema.Column{StudentsColumns[6]},
				RefColumns: []*schema.Column{ClassLeadersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "students_departments_department",
				Columns:    []*schema.Column{StudentsColumns[7]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "students_majors_major",
				Columns:    []*schema.Column{StudentsColumns[8]},
				RefColumns: []*schema.Column{MajorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "students_classes_class",
				Columns:    []*schema.Column{StudentsColumns[9]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "students_tutors_student",
				Columns:    []*schema.Column{StudentsColumns[10]},
				RefColumns: []*schema.Column{TutorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TutorsColumns holds the columns for the "tutors" table.
	TutorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// TutorsTable holds the schema information for the "tutors" table.
	TutorsTable = &schema.Table{
		Name:       "tutors",
		Columns:    TutorsColumns,
		PrimaryKey: []*schema.Column{TutorsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "account", Type: field.TypeString},
		{Name: "passwd", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeBytes},
		{Name: "email", Type: field.TypeString},
		{Name: "student_user", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_students_user",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AuthorizationResourceColumns holds the columns for the "authorization_resource" table.
	AuthorizationResourceColumns = []*schema.Column{
		{Name: "authorization_id", Type: field.TypeInt},
		{Name: "resource_id", Type: field.TypeInt},
	}
	// AuthorizationResourceTable holds the schema information for the "authorization_resource" table.
	AuthorizationResourceTable = &schema.Table{
		Name:       "authorization_resource",
		Columns:    AuthorizationResourceColumns,
		PrimaryKey: []*schema.Column{AuthorizationResourceColumns[0], AuthorizationResourceColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "authorization_resource_authorization_id",
				Columns:    []*schema.Column{AuthorizationResourceColumns[0]},
				RefColumns: []*schema.Column{AuthorizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "authorization_resource_resource_id",
				Columns:    []*schema.Column{AuthorizationResourceColumns[1]},
				RefColumns: []*schema.Column{ResourcesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoleResourceColumns holds the columns for the "role_resource" table.
	RoleResourceColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt},
		{Name: "resource_id", Type: field.TypeInt},
	}
	// RoleResourceTable holds the schema information for the "role_resource" table.
	RoleResourceTable = &schema.Table{
		Name:       "role_resource",
		Columns:    RoleResourceColumns,
		PrimaryKey: []*schema.Column{RoleResourceColumns[0], RoleResourceColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_resource_role_id",
				Columns:    []*schema.Column{RoleResourceColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_resource_resource_id",
				Columns:    []*schema.Column{RoleResourceColumns[1]},
				RefColumns: []*schema.Column{ResourcesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserRoleColumns holds the columns for the "user_role" table.
	UserRoleColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UserRoleTable holds the schema information for the "user_role" table.
	UserRoleTable = &schema.Table{
		Name:       "user_role",
		Columns:    UserRoleColumns,
		PrimaryKey: []*schema.Column{UserRoleColumns[0], UserRoleColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_role_user_id",
				Columns:    []*schema.Column{UserRoleColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_role_role_id",
				Columns:    []*schema.Column{UserRoleColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserResourceColumns holds the columns for the "user_resource" table.
	UserResourceColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "resource_id", Type: field.TypeInt},
	}
	// UserResourceTable holds the schema information for the "user_resource" table.
	UserResourceTable = &schema.Table{
		Name:       "user_resource",
		Columns:    UserResourceColumns,
		PrimaryKey: []*schema.Column{UserResourceColumns[0], UserResourceColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_resource_user_id",
				Columns:    []*schema.Column{UserResourceColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_resource_resource_id",
				Columns:    []*schema.Column{UserResourceColumns[1]},
				RefColumns: []*schema.Column{ResourcesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthorizationsTable,
		CampusTable,
		CertificatesTable,
		ClassesTable,
		ClassLeadersTable,
		DepartmentsTable,
		EducationLevelsTable,
		EnrollmentStatusTable,
		FamilyInfosTable,
		MajorsTable,
		MajorDirectionsTable,
		PracticalExperiencesTable,
		ResourcesTable,
		RolesTable,
		StudentsTable,
		TutorsTable,
		UsersTable,
		AuthorizationResourceTable,
		RoleResourceTable,
		UserRoleTable,
		UserResourceTable,
	}
)

func init() {
	CertificatesTable.ForeignKeys[0].RefTable = StudentsTable
	ClassesTable.ForeignKeys[0].RefTable = CampusTable
	ClassesTable.ForeignKeys[1].RefTable = MajorsTable
	ClassesTable.ForeignKeys[2].RefTable = DepartmentsTable
	ClassesTable.ForeignKeys[3].RefTable = ClassLeadersTable
	ClassesTable.ForeignKeys[4].RefTable = MajorDirectionsTable
	ClassesTable.ForeignKeys[5].RefTable = TutorsTable
	EducationLevelsTable.ForeignKeys[0].RefTable = StudentsTable
	EnrollmentStatusTable.ForeignKeys[0].RefTable = StudentsTable
	FamilyInfosTable.ForeignKeys[0].RefTable = StudentsTable
	MajorsTable.ForeignKeys[0].RefTable = DepartmentsTable
	PracticalExperiencesTable.ForeignKeys[0].RefTable = StudentsTable
	StudentsTable.ForeignKeys[0].RefTable = ClassLeadersTable
	StudentsTable.ForeignKeys[1].RefTable = DepartmentsTable
	StudentsTable.ForeignKeys[2].RefTable = MajorsTable
	StudentsTable.ForeignKeys[3].RefTable = ClassesTable
	StudentsTable.ForeignKeys[4].RefTable = TutorsTable
	UsersTable.ForeignKeys[0].RefTable = StudentsTable
	AuthorizationResourceTable.ForeignKeys[0].RefTable = AuthorizationsTable
	AuthorizationResourceTable.ForeignKeys[1].RefTable = ResourcesTable
	RoleResourceTable.ForeignKeys[0].RefTable = RolesTable
	RoleResourceTable.ForeignKeys[1].RefTable = ResourcesTable
	UserRoleTable.ForeignKeys[0].RefTable = UsersTable
	UserRoleTable.ForeignKeys[1].RefTable = RolesTable
	UserResourceTable.ForeignKeys[0].RefTable = UsersTable
	UserResourceTable.ForeignKeys[1].RefTable = ResourcesTable
}
