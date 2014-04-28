package nerdz

// getNumericFollowers returns a slice containing the IDs of users that followed this project
func (prj *Project) getNumericFollowers() []int64 {
	var followers []int64
	db.Model(ProjectFollower{}).Where(ProjectFollower{Group: prj.Counter}).Pluck("\"user\"", &followers)
	return followers
}

// getNumericMembers returns a slice containing the IDs of users that are member of this project
func (prj *Project) getNumericMembers() []int64 {
	var members []int64
	db.Model(ProjectMember{}).Where(ProjectMember{Group: prj.Counter}).Pluck("\"user\"", &members)
	return members
}
