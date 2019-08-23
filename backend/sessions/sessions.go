package sessions
//
//import (
//	"time"
//
//	"github.com/gorilla/sessions"
//)
//
//type Session struct {
//	key   [32]byte
//	store *sessions.CookieStore
//}
//
//type Cookie struct {
//	Name       string
//	Value      string
//	Path       string
//	Domain     string
//	Expires    time.Time
//	RawExpires string
//
//	// MaxAge=0 means no 'Max-Age' attribute specified.
//	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
//	// MaxAge>0 means Max-Age attribute present and given in seconds
//	MaxAge   int
//	Secure   bool
//	HttpOnly bool
//	Raw      string
////	Unparsed []string // Raw text of unparsed attribute-value pairs
//}
//
//func (s *Session) Login() bool {
//	return false
//}
//// first cookie of client
//func init() {
//	//session := Session{}
//	//
//	//// assigns a randomly generated cookie to client
//	//authKey := securecookie.GenerateRandomKey(32)
//	//copy(session.key[:], authKey)
//	//
//	//session.store.Options = &sessions.Options{
//	//	HttpOnly: true,
//
//}
