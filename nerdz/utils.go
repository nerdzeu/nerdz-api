package nerdz

// Users returns a slice of pointer to User, fetched from its Ids
func Users(ids []uint64) []*User {
	var users []*User
	for _, elem := range ids {
		user, _ := NewUser(elem)
		users = append(users, user)
	}

	return users
}

// Projects returns a slice of pointer to Project, fetched from its Ids
func Projects(ids []uint64) []*Project {
	var projects []*Project
	for _, elem := range ids {
		project, _ := NewProject(elem)
		projects = append(projects, project)
	}
	return projects
}
