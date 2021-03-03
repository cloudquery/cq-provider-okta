package resources

import (
	"github.com/mitchellh/mapstructure"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"time"
)

type User struct {
	Activated             *time.Time   `json:"activated,omitempty"`
	Created               *time.Time   `json:"created,omitempty"`
	Id                    string       `json:"id,omitempty"`
	LastLogin             *time.Time   `json:"lastLogin,omitempty"`
	LastUpdated           *time.Time   `json:"lastUpdated,omitempty"`
	PasswordChanged       *time.Time   `json:"passwordChanged,omitempty"`
	Profile               *UserProfile `json:"profile,omitempty" gorm:"embedded;embeddedPrefix:profile_"`
	Status                string       `json:"status,omitempty"`
	StatusChanged         *time.Time   `json:"statusChanged,omitempty"`
	TransitioningToStatus string       `json:"transitioningToStatus,omitempty"`
	UserTypeId            string       `json:"userTypeId,omitempty"`
}

func (u User) TableName() string {
	return "okta_users"
}

func TransformUser(u *okta.User) *User {
	return &User{
		Activated:             u.Activated,
		Created:               u.Created,
		Id:                    u.Id,
		LastLogin:             u.LastLogin,
		LastUpdated:           u.LastUpdated,
		PasswordChanged:       u.PasswordChanged,
		Profile:               TransformUserProfile(u.Profile),
		Status:                u.Status,
		StatusChanged:         u.StatusChanged,
		TransitioningToStatus: u.TransitioningToStatus,
		UserTypeId:            u.Type.Id,
	}
}

func TransformUsers(uu []*okta.User) []*User {
	tuu := make([]*User, len(uu))
	for i, u := range uu {
		tuu[i] = TransformUser(u)
	}
	return tuu
}

type UserProfile struct {
	Login             string
	FirstName         string
	LastName          string
	MiddleName        string
	HonorificPrefix   string
	HonorificSuffix   string
	Email             string
	Title             string
	DisplayName       string
	NickName          string
	ProfileUrl        string
	SecondEmail       string
	MobilePhone       string
	PrimaryPhone      string
	StreetAddress     string
	City              string
	State             string
	ZipCode           string
	CountryCode       string
	PostalAddress     string
	PreferredLanguage string
	Locale            string
	Timezone          string
	UserType          string
	EmployeeNumber    string
	CostCenter        string
	Organization      string
	Division          string
	Department        string
	ManagerId         string
	Manager           string
}

func TransformUserProfile(profile *okta.UserProfile) *UserProfile {
	if profile == nil {
		return nil
	}
	var up UserProfile
	err := mapstructure.Decode(profile, &up)
	if err != nil {
		return nil
	}
	return &up
}

type UserType struct {
	Id            string     `json:"id,omitempty" gorm:"primaryKey"`
	Created       *time.Time `json:"created,omitempty"`
	CreatedBy     string     `json:"createdBy,omitempty"`
	Default       *bool      `json:"default,omitempty"`
	Description   string     `json:"description,omitempty"`
	DisplayName   string     `json:"displayName,omitempty"`
	LastUpdated   *time.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy string     `json:"lastUpdatedBy,omitempty"`
	Name          string     `json:"name,omitempty"`
}

func (UserType) TableName() string {
	return "okta_user_types"
}

func TransformUserType(userType *okta.UserType) *UserType {
	if userType == nil {
		return nil
	}
	return &UserType{
		Id:            userType.Id,
		Created:       userType.Created,
		CreatedBy:     userType.CreatedBy,
		Default:       userType.Default,
		Description:   userType.Description,
		DisplayName:   userType.DisplayName,
		LastUpdated:   userType.LastUpdated,
		LastUpdatedBy: userType.LastUpdatedBy,
		Name:          userType.Name,
	}
}

func TransformUserTypes(uu []*okta.UserType) []*UserType {
	tut := make([]*UserType, len(uu))
	for i, u := range uu {
		tut[i] = TransformUserType(u)
	}
	return tut
}

//
//import (
//	"context"
//	"fmt"
//	"github.com/mitchellh/mapstructure"
//	"github.com/okta/okta-sdk-golang/v2/okta"
//	"go.uber.org/zap"
//	"log"
//	"reflect"
//	"strings"
//	"time"
//)
//
//type User struct {
//	ID        uint   `gorm:"primarykey"`
//	Domain    string `neo:"unique"`
//	Activated *time.Time
//	Created   *time.Time
//
//	Groups            []*UserGroup `gorm:"constraint:OnDelete:CASCADE;"`
//	Login             string
//	FirstName         string
//	LastName          string
//	MiddleName        string
//	HonorificPrefix   string
//	HonorificSuffix   string
//	Email             string
//	Title             string
//	DisplayName       string
//	NickName          string
//	ProfileUrl        string
//	SecondEmail       string
//	MobilePhone       string
//	PrimaryPhone      string
//	StreetAddress     string
//	City              string
//	State             string
//	ZipCode           string
//	CountryCode       string
//	PostalAddress     string
//	PreferredLanguage string
//	Locale            string
//	Timezone          string
//	UserType          string
//	EmployeeNumber    string
//	CostCenter        string
//	Organization      string
//	Division          string
//	Department        string
//	ManagerId         string
//	Manager           string
//
//	CredentialsProviderName string
//	CredentialsProviderType string
//
//	ResourceID            string `neo:"unique"`
//	LastLogin             *time.Time
//	LastUpdated           *time.Time
//	PasswordChanged       *time.Time
//	Status                string
//	StatusChanged         *time.Time
//	TransitioningToStatus string
//}
//
//func (User) TableName() string {
//	return "okta_application_users"
//}
//
//type UserGroup struct {
//	UserGroupID           uint   `gorm:"primarykey"`
//	Domain                string `gorm:"-"`
//	UserID                uint   `neo:"ignore"`
//	Created               *time.Time
//	GroupID               string
//	LastMembershipUpdated *time.Time
//	LastUpdated           *time.Time
//	Name                  string
//	Description           string
//	Type                  string
//}
//
//func (UserGroup) TableName() string {
//	return "okta_application_user_groups"
//}
//
//func transformUserGroups(domain string values []*okta.Group) []*UserGroup {
//	var tValues []*UserGroup
//	for _, v := range values {
//		tValues = append(tValues, &UserGroup{
//			Domain:                domain,
//			Created:               v.Created,
//			GroupID:               v.Id,
//			LastMembershipUpdated: v.LastMembershipUpdated,
//			LastUpdated:           v.LastUpdated,
//			Type:                  v.Type,
//			Name:                  v.Profile.Name,
//			Description:           v.Profile.Description,
//		})
//	}
//	return tValues
//}
//
//func TransformUser(domain string, resource okta.UserResource, value *okta.User) *User {
//	res := User{
//		Domain:                domain,
//		Activated:             value.Activated,
//		Created:               value.Created,
//		ResourceID:            value.Id,
//		LastLogin:             value.LastLogin,
//		LastUpdated:           value.LastUpdated,
//		PasswordChanged:       value.PasswordChanged,
//		Status:                value.Status,
//		StatusChanged:         value.StatusChanged,
//		TransitioningToStatus: value.TransitioningToStatus,
//	}
//
//	if value.Credentials != nil {
//		if value.Credentials.Provider != nil {
//			res.CredentialsProviderName = value.Credentials.Provider.Name
//			res.CredentialsProviderType = value.Credentials.Provider.Type
//		}
//	}
//	if value.Profile != nil {
//		for key, value := range *value.Profile {
//			v := reflect.ValueOf(&res).Elem()
//			field := v.FieldByName(strings.Title(key))
//			field.SetString(fmt.Sprintf("%v", value))
//		}
//	}
//
//	groups, _, err := resource.ListUserGroups(context.Background(), value.Id)
//	if err != nil {
//		log.Fatal(err)
//	}
//	res.Groups = p.transformUserGroups(groups)
//	return &res
//}
//
//func (p *main.Provider) transformUsers(values []*okta.User) []*User {
//	var tValues []*User
//	for _, v := range values {
//		tValues = append(tValues, p.transformUser(v))
//	}
//	return tValues
//}
//
//type UserConfig struct {
//	Filter string
//}
//
//var userTables = []interface{}{
//	&User{},
//	&UserGroup{},
//}
//
//func (p *main.Provider) users(gConfig interface{}) error {
//	var config UserConfig
//	err := mapstructure.Decode(gConfig, &config)
//	if err != nil {
//		return err
//	}
//
//	//filter := query.NewQueryParams()
//	users, _, err := p.client.User.ListUsers(context.Background(), nil)
//	if err != nil {
//		return err
//	}
//
//	p.db.Where("domain", p.config.Domain).Delete(userTables...)
//	p.db.ChunkedCreate(p.transformUsers(users))
//	p.log.Info("Fetched resources", zap.Int("count", len(users)))
//
//	return nil
//}
