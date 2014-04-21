package orm

// getNumericFollowers returns a slice containing the IDs of users that followed this project
func (prj *Project) getNumericFollowers() []int64 {
	var followers []int64
	var fol []ProjectFollower

	db.Find(&fol, ProjectFollower{Group: prj.Counter})
	for _, elem := range fol {
		followers = append(followers, elem.User)
	}

	return followers
}
