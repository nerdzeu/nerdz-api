package nerdz

import (
	"errors"
	"net/url"
)

// ProjectInfo is the struct that contains all the project's informations
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

// Begin *Numeric* Methods

// GetNumericFollowers returns a slice containing the IDs of users that followed this project
func (prj *Project) GetNumericFollowers() []int64 {
	var followers []int64
	db.Model(ProjectFollower{}).Where(ProjectFollower{Group: prj.Counter}).Pluck("\"user\"", &followers)
	return followers
}

// GetNumericMembers returns a slice containing the IDs of users that are member of this project
func (prj *Project) GetNumericMembers() []int64 {
	var members []int64
	db.Model(ProjectMember{}).Where(ProjectMember{Group: prj.Counter}).Pluck("\"user\"", &members)
	return members
}

// GetFollowers returns a []*User that follows the project
func (prj *Project) GetFollowers() []*User {
	return getUsers(prj.GetNumericFollowers())
}

// End *Numeric* Methods

// GetMembers returns a slice of Users members of the project
func (prj *Project) GetMembers() []*User {
	return getUsers(prj.GetNumericMembers())
}

// GetProjectInfo returns a ProjectInfo struct
func (prj *Project) GetProjectInfo() *ProjectInfo {
	owner, _ := NewUser(prj.Owner)
	website, _ := url.Parse(prj.Website)
	photo, _ := url.Parse(prj.Photo.String)

	return &ProjectInfo{
		Id:          prj.Counter,
		Owner:       owner,
		Members:     prj.GetMembers(),
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

// Implements Board interface

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

// GetPostlist returns the specified posts on the project
func (prj *Project) GetPostlist(options *PostlistOptions) interface{} {
	var posts []ProjectPost
	var projectPost ProjectPost
	projectPosts := projectPost.TableName()
	users := new(User).TableName()

	query := db.Model(projectPost).Order("hpid DESC").
		Joins("JOIN "+users+" ON "+users+".counter = "+projectPosts+".to"). //PostListOptions.Language support
		Where("(\"to\" = ?)", prj.Counter)
	query = postlistQueryBuilder(query, options)
	query.Find(&posts)
	return posts
}

// IsClosed returns a boolean indicating if the board is closed
func (prj *Project) IsClosed() bool {
	return !prj.Open
}
