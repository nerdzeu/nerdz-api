package orm

import (
	"errors"
	"net/url"
)

type ProjectInfo struct {
	Id          int64
	Owner       *User
	Members     []*User
	Followers   []*User
	Description string
	Name        string
	Photo       *url.URL
	Website     *url.URL
	Goal        string
	Visible     bool
	Private     bool
	Open        bool
}

// New initializes a Project struct
func NewProject(id int64) (prj *Project, e error) {
	prj = new(Project)
	db.First(prj, id)

	if prj.Counter != id {
		return nil, errors.New("Invalid id")
	}

	return prj, nil
}

// GetFollowers returns a []*User that follows the project
func (prj *Project) GetFollowers() []*User {
	var fol []ProjectFollower
	db.Find(&fol, ProjectFollower{Group: prj.Counter})

	var followers []*User
	for _, elem := range fol {
		user, _ := NewUser(elem.User)
		followers = append(followers, user)
	}

	return followers
}

// GetProjectInfo returns a ProjectInfo struct
func (prj *Project) GetProjectInfo() *ProjectInfo {
	owner, _ := NewUser(prj.Owner)

	var mem []ProjectMember
	db.Find(&mem, ProjectMember{Group: prj.Counter})

	var members []*User
	for _, elem := range mem {
		user, _ := NewUser(elem.User)
		members = append(members, user)
	}

	website, _ := url.Parse(prj.Website)
	photo, _ := url.Parse(prj.Photo.String)

	return &ProjectInfo{
		Id:          prj.Counter,
		Owner:       owner,
		Members:     members,
		Followers:   prj.GetFollowers(),
		Description: prj.Description,
		Name:        prj.Name,
		Photo:       photo,
		Website:     website,
		Goal:        prj.Goal,
		Visible:     prj.Visible,
		Private:     prj.Private,
		Open:        prj.Open}
}

//GetInfo returns a *Info struct
func (prj *Project) GetInfo() *Info {

	website, _ := url.Parse(prj.Website)
	owner, _ := NewUser(prj.Owner)
	image, _ := url.Parse(prj.Photo.String)

	return &Info{
		Id:        prj.Counter,
		Owner:     owner,
		Followers: prj.GetFollowers(),
		Name:      prj.Name,
		Website:   website,
		Image:     image}
}
