package mongohook

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
)

// Execer write the logrus entry to the database
type Execer interface {
	Exec(entry *logrus.Entry) error
}

// NewExec create an exec instance
func NewExec(sess *mgo.Session, dbName, cName string) Execer {
	return &defaultExec{sess, dbName, cName}
}

type defaultExec struct {
	sess   *mgo.Session
	dbName string
	cName  string
}

func (e *defaultExec) Exec(entry *logrus.Entry) error {
	item := make(bson.M)

	for k, v := range entry.Data {
		item[k] = v
	}

	item["level"] = entry.Level
	item["message"] = entry.Message
	item["created"] = entry.Time.Unix()

	sess := e.sess.Clone()
	defer sess.Close()
	return sess.DB(e.dbName).C(e.cName).Insert(item)
}
