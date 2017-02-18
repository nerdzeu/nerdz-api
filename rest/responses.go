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
along with this program.  If not, see <http:// www.gnu.org/licenses/>.
*/

/* This file is generated as stated in the doc folder:
rts -server http://localhost:9090/v1 -pkg rest \
    -headers "Authorization: Bearer _uZV-FCsS3-ytssqZC6qLw" \
	-routes routes.txt -out ../rest/responses.go
and manually refined */

package rest

// ProjectsId is a response
//
// swagger:response ProjectsId
type ProjectsId struct {
	// in: body
	Body struct {
		Data struct {
			Board  string `json:"board"`
			Closed bool   `json:"closed"`
			ID     int64  `json:"id"`
			Image  string `json:"image"`
			Name   string `json:"name"`
			Owner  struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"owner"`
			Type     string `json:"type"`
			Username string `json:"username"`
			Website  string `json:"website"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdPosts is a response
//
// swagger:response ProjectsIdPosts
type ProjectsIdPosts struct {
	// in: body
	Body struct {
		Data []struct {
			Bookmarkers int64 `json:"bookmarkers"`
			CanBookmark bool  `json:"canBookmark"`
			CanComment  bool  `json:"canComment"`
			CanDelete   bool  `json:"canDelete"`
			CanEdit     bool  `json:"canEdit"`
			CanLurk     bool  `json:"canLurk"`
			Closed      bool  `json:"closed"`
			Comments    int64 `json:"comments"`
			From        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Lurkers   int64  `json:"lurkers"`
			Message   string `json:"message"`
			News      bool   `json:"news"`
			Pid       int64  `json:"pid"`
			Rate      int64  `json:"rate"`
			Revisions int64  `json:"revisions"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board  string `json:"board"`
				Closed bool   `json:"closed"`
				ID     int64  `json:"id"`
				Image  string `json:"image"`
				Name   string `json:"name"`
				Owner  struct {
					Board    string      `json:"board"`
					Closed   bool        `json:"closed"`
					ID       int64       `json:"id"`
					Image    string      `json:"image"`
					Name     string      `json:"name"`
					Owner    interface{} `json:"owner"`
					Type     string      `json:"type"`
					Username string      `json:"username"`
					Website  string      `json:"website"`
				} `json:"owner"`
				Type     string `json:"type"`
				Username string `json:"username"`
				Website  string `json:"website"`
			} `json:"to"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// Me is a response
//
// swagger:response Me
type Me struct {
	//  in: body
	Body struct {
		Data struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string   `json:"biography"`
				Birthday  string   `json:"birthday"`
				Gender    bool     `json:"gender"`
				Gravatar  string   `json:"gravatar"`
				ID        int64    `json:"id"`
				Interests []string `json:"interests"`
				Name      string   `json:"name"`
				Nation    string   `json:"nation"`
				Online    bool     `json:"online"`
				Quotes    []string `json:"quotes"`
				Surname   string   `json:"surname"`
				Timezone  string   `json:"timezone"`
				Username  string   `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdFollowers is a response
//
// swagger:response ProjectsIdFollowers
type ProjectsIdFollowers struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string   `json:"biography"`
				Birthday  string   `json:"birthday"`
				Gender    bool     `json:"gender"`
				Gravatar  string   `json:"gravatar"`
				ID        int64    `json:"id"`
				Interests []string `json:"interests"`
				Name      string   `json:"name"`
				Nation    string   `json:"nation"`
				Online    bool     `json:"online"`
				Quotes    []string `json:"quotes"`
				Surname   string   `json:"surname"`
				Timezone  string   `json:"timezone"`
				Username  string   `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePostsPidCommentsCidVotes is a response
//
// swagger:response MePostsPidCommentsCidVotes
type MePostsPidCommentsCidVotes struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// MePostsPidVotes is a response
//
// swagger:response MePostsPidVotes
type MePostsPidVotes struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// UsersIdPostsPidLurks is a response
//
// swagger:response UsersIdPostsPidLurks
type UsersIdPostsPidLurks struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// ProjectsIdPostsPidVotes is a response
//
// swagger:response ProjectsIdPostsPidVotes
type ProjectsIdPostsPidVotes struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// ProjectsIdPostsPidCommentsCidVotes is a response
//
// swagger:response ProjectsIdPostsPidCommentsCidVotes
type ProjectsIdPostsPidCommentsCidVotes struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// UsersIdWhitelist is a response
//
// swagger:response UsersIdWhitelist
type UsersIdWhitelist struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string   `json:"biography"`
				Birthday  string   `json:"birthday"`
				Gender    bool     `json:"gender"`
				Gravatar  string   `json:"gravatar"`
				ID        int64    `json:"id"`
				Interests []string `json:"interests"`
				Name      string   `json:"name"`
				Nation    string   `json:"nation"`
				Online    bool     `json:"online"`
				Quotes    []string `json:"quotes"`
				Surname   string   `json:"surname"`
				Timezone  string   `json:"timezone"`
				Username  string   `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePostsPidLurks is a response
//
// swagger:response MePostsPidLurks
type MePostsPidLurks struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// MeBlacklist is a response
//
// swagger:response MeBlacklist
type MeBlacklist struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdMembers is a response
//
// swagger:response ProjectsIdMembers
type ProjectsIdMembers struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersId is a response
//
// swagger:response UsersId
type UsersId struct {
	// in: body
	Body struct {
		Data struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string   `json:"biography"`
				Birthday  string   `json:"birthday"`
				Gender    bool     `json:"gender"`
				Gravatar  string   `json:"gravatar"`
				ID        int64    `json:"id"`
				Interests []string `json:"interests"`
				Name      string   `json:"name"`
				Nation    string   `json:"nation"`
				Online    bool     `json:"online"`
				Quotes    []string `json:"quotes"`
				Surname   string   `json:"surname"`
				Timezone  string   `json:"timezone"`
				Username  string   `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePostsPidLocks is a response
//
// swagger:response MePostsPidLocks
type MePostsPidLocks struct {
	// in: body
	Body struct {
		Data []struct {
			Counter   int64  `json:"counter"`
			Hpid      int64  `json:"hpid"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			User      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"user"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePostsPidCommentsCid is a response
//
// swagger:response MePostsPidCommentsCid
type MePostsPidCommentsCid struct {
	// in: body
	Body struct {
		Data struct {
			CanDelete bool `json:"canDelete"`
			CanEdit   bool `json:"canEdit"`
			From      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hcid      int64  `json:"hcid"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Message   string `json:"message"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdPostsPidBookmarks is a response
//
// swagger:response ProjectsIdPostsPidBookmarks
type ProjectsIdPostsPidBookmarks struct {
	// in: body
	Body struct {
		Data []struct {
			Counter int64 `json:"counter"`
			From    struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MeWhitelisting is a response
//
// swagger:response MeWhitelisting
type MeWhitelisting struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdPostsPidLurks is a response
//
// swagger:response ProjectsIdPostsPidLurks
type ProjectsIdPostsPidLurks struct {
	// in: body
	Body struct {
		Data []struct {
			Counter int64 `json:"counter"`
			From    struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board  string `json:"board"`
				Closed bool   `json:"closed"`
				ID     int64  `json:"id"`
				Image  string `json:"image"`
				Name   string `json:"name"`
				Owner  struct {
					Board    string      `json:"board"`
					Closed   bool        `json:"closed"`
					ID       int64       `json:"id"`
					Image    string      `json:"image"`
					Name     string      `json:"name"`
					Owner    interface{} `json:"owner"`
					Type     string      `json:"type"`
					Username string      `json:"username"`
					Website  string      `json:"website"`
				} `json:"owner"`
				Type     string `json:"type"`
				Username string `json:"username"`
				Website  string `json:"website"`
			} `json:"to"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePostsPidBookmarks is a response
//
// swagger:response MePostsPidBookmarks
type MePostsPidBookmarks struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// UsersIdPostsPidCommentsCid is a response
//
// swagger:response UsersIdPostsPidCommentsCid
type UsersIdPostsPidCommentsCid struct {
	// in: body
	Body struct {
		Data struct {
			CanDelete bool `json:"canDelete"`
			CanEdit   bool `json:"canEdit"`
			From      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hcid      int64  `json:"hcid"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Message   string `json:"message"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdPostsPidLocks is a response
//
// swagger:response ProjectsIdPostsPidLocks
type ProjectsIdPostsPidLocks struct {
	// in: body
	Body struct {
		Data []struct {
			Counter   int64  `json:"counter"`
			Hpid      int64  `json:"hpid"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			User      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"user"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdBlacklist is a response
//
// swagger:response UsersIdBlacklist
type UsersIdBlacklist struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MeFollowingUsers is a response
//
// swagger:response MeFollowingUsers
type MeFollowingUsers struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePmsOtherPmid is a response
//
// swagger:response MePmsOtherPmid
type MePmsOtherPmid struct {
	// in: body
	Body struct {
		Data struct {
			CanDelete bool `json:"canDelete"`
			CanEdit   bool `json:"canEdit"`
			From      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Lang      string `json:"lang"`
			Message   string `json:"message"`
			Pmid      int64  `json:"pmid"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
			ToRead bool `json:"toRead"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MeBlacklisting is a response
//
// swagger:response MeBlacklisting
type MeBlacklisting struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MeFollowingProjects is a response
//
// swagger:response MeFollowingProjects
type MeFollowingProjects struct {
	// in: body
	Body struct {
		Data []struct {
			Info struct {
				Board  string `json:"board"`
				Closed bool   `json:"closed"`
				ID     int64  `json:"id"`
				Image  string `json:"image"`
				Name   string `json:"name"`
				Owner  struct {
					Board    string      `json:"board"`
					Closed   bool        `json:"closed"`
					ID       int64       `json:"id"`
					Image    string      `json:"image"`
					Name     string      `json:"name"`
					Owner    interface{} `json:"owner"`
					Type     string      `json:"type"`
					Username string      `json:"username"`
					Website  string      `json:"website"`
				} `json:"owner"`
				Type     string `json:"type"`
				Username string `json:"username"`
				Website  string `json:"website"`
			} `json:"info"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdPostsPidLocks is a response
//
// swagger:response UsersIdPostsPidLocks
type UsersIdPostsPidLocks struct {
	// in: body
	Body struct {
		Data []struct {
			Counter   int64  `json:"counter"`
			Hpid      int64  `json:"hpid"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			User      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"user"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdPostsPidComments is a response
//
// swagger:response UsersIdPostsPidComments
type UsersIdPostsPidComments struct {
	// in: body
	Body struct {
		Data []struct {
			CanDelete bool `json:"canDelete"`
			CanEdit   bool `json:"canEdit"`
			From      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hcid      int64  `json:"hcid"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Message   string `json:"message"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdPostsPidBookmarks is a response
//
// swagger:response UsersIdPostsPidBookmarks
type UsersIdPostsPidBookmarks struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// UsersIdPostsPid is a response
//
// swagger:response UsersIdPostsPid
type UsersIdPostsPid struct {
	// in: body
	Body struct {
		Data struct {
			Bookmarkers int64 `json:"bookmarkers"`
			CanBookmark bool  `json:"canBookmark"`
			CanComment  bool  `json:"canComment"`
			CanDelete   bool  `json:"canDelete"`
			CanEdit     bool  `json:"canEdit"`
			CanLurk     bool  `json:"canLurk"`
			Closed      bool  `json:"closed"`
			Comments    int64 `json:"comments"`
			From        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Lurkers   int64  `json:"lurkers"`
			Message   string `json:"message"`
			News      bool   `json:"news"`
			Pid       int64  `json:"pid"`
			Rate      int64  `json:"rate"`
			Revisions int64  `json:"revisions"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdBlacklisting is a response
//
// swagger:response UsersIdBlacklisting
type UsersIdBlacklisting struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdFollowingProjects is a response
//
// swagger:response UsersIdFollowingProjects
type UsersIdFollowingProjects struct {
	// in: body
	Body struct {
		Data []struct {
			Info struct {
				Board  string `json:"board"`
				Closed bool   `json:"closed"`
				ID     int64  `json:"id"`
				Image  string `json:"image"`
				Name   string `json:"name"`
				Owner  struct {
					Board    string      `json:"board"`
					Closed   bool        `json:"closed"`
					ID       int64       `json:"id"`
					Image    string      `json:"image"`
					Name     string      `json:"name"`
					Owner    interface{} `json:"owner"`
					Type     string      `json:"type"`
					Username string      `json:"username"`
					Website  string      `json:"website"`
				} `json:"owner"`
				Type     string `json:"type"`
				Username string `json:"username"`
				Website  string `json:"website"`
			} `json:"info"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdPostsPidVotes is a response
//
// swagger:response UsersIdPostsPidVotes
type UsersIdPostsPidVotes struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// UsersIdWhitelisting is a response
//
// swagger:response UsersIdWhitelisting
type UsersIdWhitelisting struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePostsPidComments is a response
//
// swagger:response MePostsPidComments
type MePostsPidComments struct {
	// in: body
	Body struct {
		Data []struct {
			CanDelete bool `json:"canDelete"`
			CanEdit   bool `json:"canEdit"`
			From      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hcid      int64  `json:"hcid"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Message   string `json:"message"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MeFollowers is a response
//
// swagger:response MeFollowers
type MeFollowers struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdFriends is a response
//
// swagger:response UsersIdFriends
type UsersIdFriends struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdPostsPidCommentsCid is a response
//
// swagger:response ProjectsIdPostsPidCommentsCid
type ProjectsIdPostsPidCommentsCid struct {
	// in: body
	Body struct {
		Data struct {
			CanDelete bool `json:"canDelete"`
			CanEdit   bool `json:"canEdit"`
			From      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hcid      int64  `json:"hcid"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Message   string `json:"message"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board  string `json:"board"`
				Closed bool   `json:"closed"`
				ID     int64  `json:"id"`
				Image  string `json:"image"`
				Name   string `json:"name"`
				Owner  struct {
					Board    string      `json:"board"`
					Closed   bool        `json:"closed"`
					ID       int64       `json:"id"`
					Image    string      `json:"image"`
					Name     string      `json:"name"`
					Owner    interface{} `json:"owner"`
					Type     string      `json:"type"`
					Username string      `json:"username"`
					Website  string      `json:"website"`
				} `json:"owner"`
				Type     string `json:"type"`
				Username string `json:"username"`
				Website  string `json:"website"`
			} `json:"to"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MeWhitelist is a response
//
// swagger:response MeWhitelist
type MeWhitelist struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string   `json:"biography"`
				Birthday  string   `json:"birthday"`
				Gender    bool     `json:"gender"`
				Gravatar  string   `json:"gravatar"`
				ID        int64    `json:"id"`
				Interests []string `json:"interests"`
				Name      string   `json:"name"`
				Nation    string   `json:"nation"`
				Online    bool     `json:"online"`
				Quotes    []string `json:"quotes"`
				Surname   string   `json:"surname"`
				Timezone  string   `json:"timezone"`
				Username  string   `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdPostsPidCommentsCidVotes is a response
//
// swagger:response UsersIdPostsPidCommentsCidVotes
type UsersIdPostsPidCommentsCidVotes struct {
	// in: body
	Body struct {
		Data         []interface{} `json:"data"`
		HumanMessage string        `json:"humanMessage"`
		Message      string        `json:"message"`
		Status       int64         `json:"status"`
		Success      bool          `json:"success"`
	}
}

// UsersIdFollowingUsers is a response
//
// swagger:response UsersIdFollowingUsers
type UsersIdFollowingUsers struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdPostsPidComments is a response
//
// swagger:response ProjectsIdPostsPidComments
type ProjectsIdPostsPidComments struct {
	// in: body
	Body struct {
		Data []struct {
			CanDelete bool `json:"canDelete"`
			CanEdit   bool `json:"canEdit"`
			From      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hcid      int64  `json:"hcid"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Message   string `json:"message"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board  string `json:"board"`
				Closed bool   `json:"closed"`
				ID     int64  `json:"id"`
				Image  string `json:"image"`
				Name   string `json:"name"`
				Owner  struct {
					Board    string      `json:"board"`
					Closed   bool        `json:"closed"`
					ID       int64       `json:"id"`
					Image    string      `json:"image"`
					Name     string      `json:"name"`
					Owner    interface{} `json:"owner"`
					Type     string      `json:"type"`
					Username string      `json:"username"`
					Website  string      `json:"website"`
				} `json:"owner"`
				Type     string `json:"type"`
				Username string `json:"username"`
				Website  string `json:"website"`
			} `json:"to"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePostsPid is a response
//
// swagger:response MePostsPid
type MePostsPid struct {
	// in: body
	Body struct {
		Data struct {
			Bookmarkers int64 `json:"bookmarkers"`
			CanBookmark bool  `json:"canBookmark"`
			CanComment  bool  `json:"canComment"`
			CanDelete   bool  `json:"canDelete"`
			CanEdit     bool  `json:"canEdit"`
			CanLurk     bool  `json:"canLurk"`
			Closed      bool  `json:"closed"`
			Comments    int64 `json:"comments"`
			From        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Lurkers   int64  `json:"lurkers"`
			Message   string `json:"message"`
			News      bool   `json:"news"`
			Pid       int64  `json:"pid"`
			Rate      int64  `json:"rate"`
			Revisions int64  `json:"revisions"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdFollowers is a response
//
// swagger:response UsersIdFollowers
type UsersIdFollowers struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MeFriends is a response
//
// swagger:response MeFriends
type MeFriends struct {
	// in: body
	Body struct {
		Data []struct {
			Contacts struct {
				Facebook string `json:"facebook"`
				Github   string `json:"github"`
				Jabber   string `json:"jabber"`
				Skype    string `json:"skype"`
				Steam    string `json:"steam"`
				Twitter  string `json:"twitter"`
				Website  string `json:"website"`
				Yahoo    string `json:"yahoo"`
			} `json:"contacts"`
			Info struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"info"`
			Personal struct {
				Biography string      `json:"biography"`
				Birthday  string      `json:"birthday"`
				Gender    bool        `json:"gender"`
				Gravatar  string      `json:"gravatar"`
				ID        int64       `json:"id"`
				Interests interface{} `json:"interests"`
				Name      string      `json:"name"`
				Nation    string      `json:"nation"`
				Online    bool        `json:"online"`
				Quotes    []string    `json:"quotes"`
				Surname   string      `json:"surname"`
				Timezone  string      `json:"timezone"`
				Username  string      `json:"username"`
			} `json:"personal"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// ProjectsIdPostsPid is a response
//
// swagger:response ProjectsIdPostsPid
type ProjectsIdPostsPid struct {
	// in: body
	Body struct {
		Data struct {
			Bookmarkers int64 `json:"bookmarkers"`
			CanBookmark bool  `json:"canBookmark"`
			CanComment  bool  `json:"canComment"`
			CanDelete   bool  `json:"canDelete"`
			CanEdit     bool  `json:"canEdit"`
			CanLurk     bool  `json:"canLurk"`
			Closed      bool  `json:"closed"`
			Comments    int64 `json:"comments"`
			From        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Lurkers   int64  `json:"lurkers"`
			Message   string `json:"message"`
			News      bool   `json:"news"`
			Pid       int64  `json:"pid"`
			Rate      int64  `json:"rate"`
			Revisions int64  `json:"revisions"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board  string `json:"board"`
				Closed bool   `json:"closed"`
				ID     int64  `json:"id"`
				Image  string `json:"image"`
				Name   string `json:"name"`
				Owner  struct {
					Board    string      `json:"board"`
					Closed   bool        `json:"closed"`
					ID       int64       `json:"id"`
					Image    string      `json:"image"`
					Name     string      `json:"name"`
					Owner    interface{} `json:"owner"`
					Type     string      `json:"type"`
					Username string      `json:"username"`
					Website  string      `json:"website"`
				} `json:"owner"`
				Type     string `json:"type"`
				Username string `json:"username"`
				Website  string `json:"website"`
			} `json:"to"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePms is a response
//
// swagger:response MePms
type MePms struct {
	// in: body
	Body struct {
		Data []struct {
			From struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			LastMessage string `json:"lastMessage"`
			Time        string `json:"time"`
			Timestamp   int64  `json:"timestamp"`
			To          struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
			ToRead bool `json:"toRead"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePmsOther is a response
//
// swagger:response MePmsOther
type MePmsOther struct {
	// in: body
	Body struct {
		Data []struct {
			CanDelete bool `json:"canDelete"`
			CanEdit   bool `json:"canEdit"`
			From      struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Lang      string `json:"lang"`
			Message   string `json:"message"`
			Pmid      int64  `json:"pmid"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
			ToRead bool `json:"toRead"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MePosts is a response
//
// swagger:response MePosts
type MePosts struct {
	// in: body
	Body struct {
		Data []struct {
			Bookmarkers int64 `json:"bookmarkers"`
			CanBookmark bool  `json:"canBookmark"`
			CanComment  bool  `json:"canComment"`
			CanDelete   bool  `json:"canDelete"`
			CanEdit     bool  `json:"canEdit"`
			CanLurk     bool  `json:"canLurk"`
			Closed      bool  `json:"closed"`
			Comments    int64 `json:"comments"`
			From        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Lurkers   int64  `json:"lurkers"`
			Message   string `json:"message"`
			News      bool   `json:"news"`
			Pid       int64  `json:"pid"`
			Rate      int64  `json:"rate"`
			Revisions int64  `json:"revisions"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// MeHome is a response
//
// swagger:response MeHome
type MeHome struct {
	// in: body
	Body struct {
		Data []struct {
			Bookmarkers int64 `json:"bookmarkers"`
			CanBookmark bool  `json:"canBookmark"`
			CanComment  bool  `json:"canComment"`
			CanDelete   bool  `json:"canDelete"`
			CanEdit     bool  `json:"canEdit"`
			CanLurk     bool  `json:"canLurk"`
			Closed      bool  `json:"closed"`
			Comments    int64 `json:"comments"`
			From        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Lurkers   int64  `json:"lurkers"`
			Message   string `json:"message"`
			News      bool   `json:"news"`
			Pid       int64  `json:"pid"`
			Rate      int64  `json:"rate"`
			Revisions int64  `json:"revisions"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board  string `json:"board"`
				Closed bool   `json:"closed"`
				ID     int64  `json:"id"`
				Image  string `json:"image"`
				Name   string `json:"name"`
				Owner  struct {
					Board    string      `json:"board"`
					Closed   bool        `json:"closed"`
					ID       int64       `json:"id"`
					Image    string      `json:"image"`
					Name     string      `json:"name"`
					Owner    interface{} `json:"owner"`
					Type     string      `json:"type"`
					Username string      `json:"username"`
					Website  string      `json:"website"`
				} `json:"owner"`
				Type     string `json:"type"`
				Username string `json:"username"`
				Website  string `json:"website"`
			} `json:"to"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}

// UsersIdPosts is a response
//
// swagger:response UsersIdPosts
type UsersIdPosts struct {
	// in: body
	Body struct {
		Data []struct {
			Bookmarkers int64 `json:"bookmarkers"`
			CanBookmark bool  `json:"canBookmark"`
			CanComment  bool  `json:"canComment"`
			CanDelete   bool  `json:"canDelete"`
			CanEdit     bool  `json:"canEdit"`
			CanLurk     bool  `json:"canLurk"`
			Closed      bool  `json:"closed"`
			Comments    int64 `json:"comments"`
			From        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"from"`
			Hpid      int64  `json:"hpid"`
			Lang      string `json:"lang"`
			Lurkers   int64  `json:"lurkers"`
			Message   string `json:"message"`
			News      bool   `json:"news"`
			Pid       int64  `json:"pid"`
			Rate      int64  `json:"rate"`
			Revisions int64  `json:"revisions"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			To        struct {
				Board    string      `json:"board"`
				Closed   bool        `json:"closed"`
				ID       int64       `json:"id"`
				Image    string      `json:"image"`
				Name     string      `json:"name"`
				Owner    interface{} `json:"owner"`
				Type     string      `json:"type"`
				Username string      `json:"username"`
				Website  string      `json:"website"`
			} `json:"to"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"data"`
		HumanMessage string `json:"humanMessage"`
		Message      string `json:"message"`
		Status       int64  `json:"status"`
		Success      bool   `json:"success"`
	}
}
