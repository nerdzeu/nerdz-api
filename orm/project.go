package orm

import (
	"errors"
	"net/url"
)

type ProjectInfo struct {
	Id          int64
	Owner       *User
	Members     []User
	Followers   []User
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
func (prj *Project) New(id int64) error {
	db.First(prj, id)

	if prj.Counter != id {
		return errors.New("Invalid id")
	}

	return nil
}

// GetInfo returns a ProjectInfo struct
func (prj *Project) GetProjectInfo() *ProjectInfo {
	var owner User
	db.First(&owner, prj.Owner)

	var mem []ProjectMember
	db.Find(&mem, ProjectMember{Group: prj.Counter})

	var members []User
	for _, elem := range mem {
		var member User
		member.New(elem.User)
		members = append(members, member)
	}

	var fol []ProjectFollower
	db.Find(&fol, ProjectFollower{Group: prj.Counter})

	var followers []User
	for _, elem := range fol {
		var follower User
		follower.New(elem.User)
		followers = append(followers, follower)
	}

	website, _ := url.Parse(prj.Website)
	photo, _ := url.Parse(prj.Photo.String)

	return &ProjectInfo{
		Id:          prj.Counter,
		Owner:       &owner,
		Members:     members,
		Followers:   followers,
		Description: prj.Description,
		Name:        prj.Name,
		Photo:       photo,
		Website:     website,
		Goal:        prj.Goal,
		Visible:     prj.Visible,
		Private:     prj.Private,
		Open:        prj.Open}
}
