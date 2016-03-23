/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package nerdz

import "net/url"

// ProjectInfo is the struct that contains all the project's informations
type ProjectInfo struct {
	ID               uint64
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

// NewProject initializes a Project struct
func NewProject(id uint64) (prj *Project, err error) {
	prj = new(Project)
	err = Db().First(prj, id)
	return
}

// Begin *Numeric* Methods

// NumericFollowers returns a slice containing the IDs of users that followed this project
func (prj *Project) NumericFollowers() (followers []uint64) {
	Db().Model(ProjectFollower{}).Where(ProjectFollower{To: prj.Counter}).Pluck(`"from"`, &followers)
	return
}

// NumericMembers returns a slice containing the IDs of users that are member of this project
func (prj *Project) NumericMembers() (members []uint64) {
	Db().Model(ProjectMember{}).Where(ProjectMember{To: prj.Counter}).Pluck(`"from"`, &members)
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
func (prj *Project) NumericOwner() (owner uint64) {
	Db().Model(ProjectOwner{}).Select(`"from"`).Where(ProjectOwner{To: prj.Counter}).Scan(&owner)
	return
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
		ID:               prj.Counter,
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

//Info returns a *info struct
func (prj *Project) Info() *Info {
	website, _ := url.Parse(prj.Website.String)
	image, _ := url.Parse(prj.Photo.String)
	boardURL, _ := url.Parse(Configuration.NERDZHost)
	boardURL.Path = prj.Name + ":"

	return &Info{
		ID:          prj.Counter,
		Owner:       prj.Owner().Info(),
		Name:        prj.Name,
		Username:    "",
		Website:     website,
		Image:       image,
		Closed:      !prj.Open,
		BoardString: boardURL.String(),
		Type:        PROJECT}
}

//Postlist returns the specified posts on the project
func (prj *Project) Postlist(options PostlistOptions) *[]ExistingPost {
	var posts []ProjectPost
	var projectPost ProjectPost
	projectPosts := projectPost.TableName()
	users := new(User).TableName()

	query := Db().Model(projectPost).Order("hpid DESC").
		Joins("JOIN "+users+" ON "+users+".counter = "+projectPosts+".to"). //PostListOptions.Language support
		Where(`"to" = ?`, prj.Counter)

	options.User = false
	query = postlistQueryBuilder(query, options)
	query.Scan(&posts)

	var retPosts []ExistingPost
	for _, p := range posts {
		projectPost := p
		retPosts = append(retPosts, ExistingPost(&projectPost))
	}

	return &retPosts
}

// Implements Reference interface

// ID returns the project ID
func (prj *Project) ID() uint64 {
	return prj.Counter
}

// Language returns the project language
func (prj *Project) Language() string {
	return prj.Owner().Language()
}
