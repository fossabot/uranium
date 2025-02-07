package postgres

import "github.com/go-uranium/uranium/model/user"

var SQLUserByUID = `SELECT uid, username, lowercase, electrons, admin, created, deleted FROM "user" WHERE uid = $1;`

func (pg *Postgres) UserByUID(uid int32) (*user.User, error) {
	u := &user.User{}
	err := pg.db.QueryRow(SQLUserByUID, uid).
		Scan(&u.UID, &u.Username, &u.Lowercase, &u.Electrons, &u.Admin, &u.Created, &u.Deleted)
	return u, err
}

func (pg *Postgres) UserByEmail(email string) (*user.User, error) {
	return nil, nil

}

func (pg *Postgres) UserByUsername(username string) (*user.User, error) {
	return nil, nil

}

var SQLUserAuthByUID = `SELECT uid, email, password, security_email, two_factor, locked, locked_till, disabled FROM user_auth WHERE uid = $1;`

func (pg *Postgres) UserAuthByUID(uid int32) (*user.Auth, error) {
	au := &user.Auth{}
	err := pg.db.QueryRow(SQLUserAuthByUID, uid).
		Scan(&au.UID, &au.Email, &au.Password, &au.SecurityEmail, &au.TwoFactor,
			&au.Locked, &au.LockedTill, &au.Disabled)
	return au, err
}

var SQLUserProfileByUID = `SELECT uid, name, bio, location, birthday, email, social FROM user_profile WHERE uid = $1;`

func (pg *Postgres) UserProfileByUID(uid int32) (*user.Profile, error) {
	pf := &user.Profile{}
	err := pg.db.QueryRow(SQLUserProfileByUID, uid).
		Scan(&pf.UID, &pf.Name, &pf.Bio, &pf.Location, &pf.Birthday,
			&pf.Email, &pf.Social)
	return pf, err
}

var SQLUserBasicByUID = `SELECT uid, username, admin FROM "user" WHERE uid = $1;`

func (pg *Postgres) UserBasicByUID(uid int32) (*user.Basic, error) {
	bc := &user.BasicCore{}
	err := pg.db.QueryRow(SQLUserBasicByUID, uid).
		Scan(&bc.UID, &bc.Username, &bc.Admin)
	return user.NewBasicFromCore(bc), err
}

func (pg *Postgres) UserUIDByLowercase(lowercase string) (int32, error) {
	return 0, nil

}

//var (
//	SQLUserByUID      = `SELECT uid, name, username, email, avatar, bio, created_at, artifact FROM users WHERE uid = $1;`
//	SQLUserByEmail    = `SELECT uid, name, username, email, avatar, bio, created_at, artifact FROM users WHERE email = $1;`
//	SQLUserByUsername = `SELECT uid, name, username, email, avatar, bio, created_at, artifact FROM users WHERE username = $1;`
//	SQLUserAuthByUID  = `SELECT uid, password, locked, security_email FROM user_auth WHERE uid = $1;`
//
//	SQLInsertUser     = `INSERT INTO users(name, username, email, avatar, bio, created_at, artifact) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING uid;`
//	SQLInsertUserAuth = `INSERT INTO user_auth(uid, password, locked, security_email) VALUES ($1, $2, $3, $4);`
//
//	SQLUpdateUser     = `UPDATE users SET name = $2, username = $3, email = $4, avatar = $5, bio = $6, created_at = $7, artifact = $8 WHERE uid = $1;`
//	SQLUpdateUserAuth = `UPDATE user_auth SET password = $2, locked = $3, security_email = $4 WHERE uid = $1;`
//
//	SQLAddArtifact = `UPDATE users SET artifact = artifact + $2 WHERE uid = $1;`
//
//	SQLDeleteUser = `DELETE FROM user_auth WHERE uid = $1;DELETE FROM session WHERE uid = $1; DELETE FROM users WHERE uid = $1; UPDATE post_info SET creator = 0 WHERE creator = $1; UPDATE comments SET creator = 0 WHERE creator = $1;`
//
//	SQLUsernameExists = `SELECT EXISTS(SELECT uid FROM users WHERE username = $1);`
//	SQLEmailExists    = `SELECT EXISTS(SELECT uid FROM users WHERE email = $1);`
//
//	//SQLUserFollow      = `UPDATE "user" SET following = array_append(following, $2) WHERE uid = $1;`
//	//SQLUserUnFollow    = `UPDATE "user" SET following = array_remove(following, $2) WHERE uid = $1;`
//	//SQLAlreadyFollow   = `SELECT $2 = ANY ((SELECT following FROM "user" WHERE uid = $1)::int[]);`
//	//SQLFollowings      = `SELECT uid, name, username, email, avatar, bio, created_at, artifact, following FROM user WHERE uid = ANY ((SELECT following FROM "user" WHERE uid = 2)::int[]);`
//	//SQLFollowers       = `SELECT uid, name, username, email, avatar, bio, created_at, artifact, following FROM "user" WHERE $1 = ANY(following);`
//)
//
//func (pg *Postgres) UserByUID(uid int64) (*user.User, error) {
//	row := pg.db.QueryRow(SQLUserByUID, uid)
//	u, err := user.ScanUser(row)
//	if err != nil {
//		//if err == sql.ErrNoRows {
//		//	return nil, nil
//		//}
//		return &user.User{}, err
//	}
//	u.Tidy()
//	return u, nil
//}
//
//func (pg *Postgres) UserByEmail(email string) (*user.User, error) {
//	email = clean.String(email)
//	row := pg.db.QueryRow(SQLUserByEmail, email)
//	u, err := user.ScanUser(row)
//	if err != nil {
//		//if err == sql.ErrNoRows {
//		//	return nil, nil
//		//}
//		return &user.User{}, err
//	}
//	u.Tidy()
//	return u, nil
//}
//
//func (pg *Postgres) UserByUsername(username string) (*user.User, error) {
//	username = clean.String(username)
//	row := pg.db.QueryRow(SQLUserByUsername, username)
//	u, err := user.ScanUser(row)
//	if err != nil {
//		//if err == sql.ErrNoRows {
//		//	return nil, nil
//		//}
//		return &user.User{}, err
//	}
//	u.Tidy()
//	return u, nil
//}
//
//func (pg *Postgres) UserAuthByUID(uid int64) (*user.Auth, error) {
//	row := pg.db.QueryRow(SQLUserAuthByUID, uid)
//	auth, err := user.ScanAuth(row)
//	if err != nil {
//		//if err == sql.ErrNoRows {
//		//	return nil, nil
//		//}
//		return &user.Auth{}, err
//	}
//	return auth, nil
//}
//
//func (pg *Postgres) InsertUser(u *user.User) (int64, error) {
//	u.Tidy()
//	var uid int64
//	err := pg.db.QueryRow(SQLInsertUser, u.Name,
//		u.Username, u.Email, u.Avatar, u.Bio,
//		u.CreatedAt, u.Artifact).Scan(&uid)
//	if err != nil {
//		return 0, err
//	}
//	return uid, nil
//}
//
//func (pg *Postgres) InsertUserAuth(auth *user.Auth) error {
//	_, err := pg.db.Exec(SQLInsertUserAuth, auth.UID,
//		auth.Password, auth.Locked, auth.SecurityEmail)
//	return err
//}
//
//func (pg *Postgres) UpdateUser(u *user.User) error {
//	u.Tidy()
//	_, err := pg.db.Exec(SQLUpdateUser, u.UID, u.Name,
//		u.Username, u.Email, u.Avatar, u.Bio, u.CreatedAt,
//		u.Artifact)
//	return err
//}
//
//func (pg *Postgres) UpdateUserAuth(auth *user.Auth) error {
//	_, err := pg.db.Exec(SQLUpdateUserAuth, auth.UID,
//		auth.Password, auth.Locked, auth.SecurityEmail)
//	return err
//}
//
//func (pg *Postgres) AddArtifact(uid, add int64) error {
//	_, err := pg.db.Exec(SQLAddArtifact, uid, add)
//	return err
//}
//
//func (pg *Postgres) DeleteUser(uid int64) error {
//	_, err := pg.db.Exec(SQLDeleteUser, uid)
//	return err
//}
//
//func (pg *Postgres) UsernameExists(username string) (bool, error) {
//	username = strings.ToLower(username)
//	row := pg.db.QueryRow(SQLUsernameExists, username)
//	e := true
//	err := row.Scan(&e)
//	if err != nil {
//		return true, err
//	}
//	return e, nil
//}
//
//func (pg *Postgres) EmailExists(email string) (bool, error) {
//	email = strings.ToLower(email)
//	row := pg.db.QueryRow(SQLEmailExists, email)
//	e := true
//	err := row.Scan(&e)
//	if err != nil {
//		return true, err
//	}
//	return e, nil
//}

//func (pg *Postgres) UserFollow(op, target int64) error {
//	_, err := pg.db.Exec(SQLUserFollow, op, target)
//	return err
//}
//
//func (pg *Postgres) UserUnFollow(op, target int64) error {
//	_, err := pg.db.Exec(SQLUserUnFollow, op, target)
//	return err
//}
//
//func (pg *Postgres) AlreadyFollow(op, target int64) (bool, error) {
//	already := true
//	err := pg.db.QueryRow(SQLAlreadyFollow, op, target).Scan(&already)
//	if err != nil {
//		return true, err
//	}
//	return already, nil
//}
//
//func (pg *Postgres) Followings(uid int64) ([]*user.User, error) {
//	rows, err := pg.db.Query(SQLFollowings, uid)
//	if err != nil {
//		return nil, err
//	}
//	var followings []*user.User
//	for rows.Next() {
//		u, err := user.ScanUser(rows)
//		if err != nil {
//			return nil, err
//		}
//		followings = append(followings, u)
//	}
//	return followings, nil
//}
//
//func (pg *Postgres) Followers(uid int64) ([]*user.User, error) {
//	rows, err := pg.db.Query(SQLFollowers, uid)
//	if err != nil {
//		return nil, err
//	}
//	var followers []*user.User
//	for rows.Next() {
//		u, err := user.ScanUser(rows)
//		if err != nil {
//			return nil, err
//		}
//		followers = append(followers, u)
//	}
//	return followers, nil
//}
