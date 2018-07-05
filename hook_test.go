package mongohook_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/LyricTian/logrus-mongo-hook"
	"github.com/Sirupsen/logrus"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	mgoURL = "mongodb://travis:test@127.0.0.1:27017/admin"
	dbName = "mydb_test"
)

func TestHook(t *testing.T) {
	sess, err := mgo.Dial(mgoURL)
	if err != nil {
		t.Error(err)
		return
	}
	defer sess.Close()

	var filter = func(entry *logrus.Entry) *logrus.Entry {
		if _, ok := entry.Data["foo2"]; ok {
			delete(entry.Data, "foo2")
		}

		return entry
	}

	cName := "t_log"
	hook := mongohook.Default(sess, dbName, cName,
		mongohook.SetExtra(map[string]interface{}{"foo": "bar"}),
		mongohook.SetFilter(filter),
	)

	defer sess.DB(dbName).C(cName).DropCollection()

	log := logrus.New()
	log.AddHook(hook)

	log.WithField("foo2", "bar").Infof("test foo")
	hook.Flush()

	var item map[string]interface{}
	err = sess.DB(dbName).C(cName).Find(nil).Select(bson.M{"_id": 0}).One(&item)
	if err != nil {
		t.Error(err)
		return
	} else if item == nil {
		t.Error("Not expected value:nil")
		return
	}

	if reflect.DeepEqual(item["level"], logrus.InfoLevel) {
		t.Errorf("Not expected value:%v", item["level"])
		return
	}

	if item["message"].(string) != "test foo" {
		t.Errorf("Not expected value:%v", item["message"])
		return
	}

	if item["foo"].(string) != "bar" {
		t.Errorf("Not expected value:%v", item["foo"])
		return
	}

	if time.Unix(item["created"].(int64), 0).IsZero() {
		t.Errorf("Not expected value:%v", item["created"])
	}
}

func ExampleHook() {
	sess, err := mgo.Dial(mgoURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sess.Close()

	cName := "e_log"
	hook := mongohook.Default(sess, dbName, cName)
	defer sess.DB(dbName).C(cName).DropCollection()

	log := logrus.New()
	log.AddHook(hook)
	log.WithField("foo", "bar").Info("foo test")
	hook.Flush()

	var item struct {
		Message string `bson:"message"`
	}
	err = sess.DB(dbName).C(cName).Find(nil).Select(bson.M{"_id": 0, "message": 1}).One(&item)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item.Message)

	// Output: foo test
}
