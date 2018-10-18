package mongohook

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
)

// ExecCloser write the logrus entry to the database and close the database
type ExecCloser interface {
	Exec(entry *logrus.Entry) error
	Close() error
}

// NewExec create an exec instance
func NewExec(sess *mgo.Session, dbName, cName string) ExecCloser {
	return &defaultExec{
		sess:   sess,
		dbName: dbName,
		cName:  cName,
	}
}

// NewExecWithURL create an exec instance
func NewExecWithURL(url, dbName, cName string) ExecCloser {
	sess, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	return &defaultExec{
		sess:     sess,
		dbName:   dbName,
		cName:    cName,
		canClose: true,
	}
}

type defaultExec struct {
	sess     *mgo.Session
	dbName   string
	cName    string
	canClose bool
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

func (e *defaultExec) Close() error {
	if e.canClose {
		e.sess.Close()
	}
	return nil
}
