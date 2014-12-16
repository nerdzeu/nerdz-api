package nerdz

import (
	"errors"
	"net/url"
)

// ProjectInfo is the struct that contains all the project's informations
type ProjectInfo struct {
	Id               uint64
	Owner            *User
	Members          []*User
	NumericMembers   []uint64
	Followers        []*User
	NumericFollowers []uint64
	Description      string
	Name             string
	Photo            *url.URL
	Website          *url.URL
	Goal             string
	Visible          bool
	Private          bool
	Open             bool
}

// New initializes a Project struct
func NewProject(id uint64) (prj *Project, e error) {
	prj = new(Project)
	db.First(prj, id)

	if prj.Counter != id {
		return nil, errors.New("Invalid id")
	}

	return prj, nil
}

// Begin *Numeric* Methods

// NumericFollowers returns a slice containing the IDs of users that followed this project
func (prj *Project) NumericFollowers() (followers []uint64) {
	db.Model(ProjectFollower{}).Where(ProjectFollower{To: prj.Counter}).Pluck("\"from\"", &followers)
	return
}

// NumericMembers returns a slice containing the IDs of users that are member of this project
func (prj *Project) NumericMembers() (members []uint64) {
	db.Model(ProjectMember{}).Where(ProjectMember{To: prj.Counter}).Pluck("\"from\"", &members)
	return
}

// Followers returns a []*User that follows the project
func (prj *Project) Followers() []*User {
	return Users(prj.NumericFollowers())
}

// End *Numeric* Methods

// Members returns a slice of Users members of the project
func (prj *Project) Members() []*User {
	return Users(prj.NumericMembers())
}

// NumericOwner returns the Id of the owner of the project
func (prj *Project) NumericOwner() uint64 {
	owners := make([]uint64, 1)
	db.Model(ProjectOwner{}).Where(ProjectOwner{To: prj.Counter}).Pluck("\"from\"", &owners)
	return owners[0]
}

// Owner returns the *User owner of the project
func (prj *Project) Owner() (owner *User) {
	owner, _ = NewUser(prj.NumericOwner())
	return
}

// ProjectInfo returns a ProjectInfo struct
func (prj *Project) ProjectInfo() *ProjectInfo {
	website, _ := url.Parse(prj.Website.String)
	photo, _ := url.Parse(prj.Photo.String)

	return &ProjectInfo{
		Id:               prj.Counter,
		Owner:            prj.Owner(),
		Members:          prj.Members(),
		NumericMembers:   prj.NumericMembers(),
		Followers:        prj.Followers(),
		NumericFollowers: prj.NumericFollowers(),
		Description:      prj.Description,
		Name:             prj.Name,
		Photo:            photo,
		Website:          website,
		Goal:             prj.Goal,
		Visible:          prj.Visible,
		Private:          prj.Private,
		Open:             prj.Open}
}

// Implements Board interface

//Info returns a *Info struct
func (prj *Project) Info() *Info {
	website, _ := url.Parse(prj.Website.String)
	image, _ := url.Parse(prj.Photo.String)

	return &Info{
		Id:               prj.Counter,
		Owner:            prj.Owner(),
		NumericOwner:     prj.NumericOwner(),
		Followers:        prj.Followers(),
		NumericFollowers: prj.NumericFollowers(),
		Name:             prj.Name,
		Website:          website,
		Image:            image,
		Closed:           !prj.Open}
}

// Postlist returns the specified posts on the project
func (prj *Project) Postlist(options *PostlistOptions) interface{} {
	var posts []ProjectPost
	var projectPost ProjectPost
	projectPosts := projectPost.TableName()
	users := new(User).TableName()

	query := db.Model(projectPost).Order("hpid DESC").
		Joins("JOIN "+users+" ON "+users+".counter = "+projectPosts+".to"). //PostListOptions.Language support
		Where("(\"to\" = ?)", prj.Counter)
	if options != nil {
		options.User = false
	} else {
		options = new(PostlistOptions)
		options.User = false
	}
	query = postlistQueryBuilder(query, options)
	query.Find(&posts)
	return posts
}
