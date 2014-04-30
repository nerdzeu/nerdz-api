package nerdz

func getUsers(ids []int64) []*User {
	var users []*User
	for _, elem := range ids {
		user, _ := NewUser(elem)
		users = append(users, user)
	}

	return users
}

func getProjects(ids []int64) []*Project {
	var projects []*Project
	for _, elem := range ids {
		project, _ := NewProject(elem)
		projects = append(projects, project)
	}
	return projects
}

func idInSlice(id int64, slice []int64) bool {
	for _, e := range slice {
		if e == id {
			return true
		}
	}
	return false
}
