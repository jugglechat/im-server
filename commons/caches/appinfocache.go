package caches

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jugglechat/im-server/commons/dbcommons"
)

var appInfoCache *LruCache

type AppInfo struct {
	AppKey       string    `default:"-"`
	AppSecret    string    `default:"-"`
	AppSecureKey string    `default:"-"`
	AppStatus    string    `default:"-"`
	CreatedTime  time.Time `default:"-"`

	TestItem  string
	TestInt   int
	TestBool  bool  `default:"true"`
	TestInt64 int64 `default:"10"`
}

func init() {
	appInfoCache = NewLruCache(100, func(key, value interface{}) {})
	appInfoCache.AddTimeoutAfterRead(5 * time.Minute)
	appInfoCache.AddTimeoutAfterCreate(10 * time.Minute)
	appInfoCache.AddValueCreator(func(key interface{}) interface{} {
		appTable := dbcommons.AppTable{}
		app := appTable.FindByAppkey(key.(string))
		if app != nil {
			appInfo := &AppInfo{
				AppKey:       app.AppKey,
				AppSecret:    app.AppSecret,
				AppSecureKey: app.AppSecureKey,
				AppStatus:    app.AppStatus,
				CreatedTime:  app.CreatedTime,
			}

			appExtTable := dbcommons.AppExtTable{}
			appExtList := appExtTable.FindListByAppkey(key.(string))

			extMap := make(map[string]string)
			if len(appExtList) > 0 {
				for _, appExt := range appExtList {
					extMap[strings.ToLower(appExt.AppItemKey)] = appExt.AppItemValue
				}
			}
			filledExtValue(appInfo, extMap)
			return appInfo
		}
		return &AppInfo{
			AppKey: key.(string),
		}
	})
}
func filledExtValue(appInfo *AppInfo, extMap map[string]string) {
	appInfoVal := reflect.ValueOf(appInfo).Elem()
	for i := 0; i < appInfoVal.NumField(); i++ {
		fieldName := appInfoVal.Type().Field(i).Name
		fieldType := appInfoVal.Type().Field(i).Type
		fieldTag := appInfoVal.Type().Field(i).Tag
		defaultStr := strings.TrimSpace(fieldTag.Get("default"))
		if defaultStr != "-" {
			setFieldValue(appInfoVal.FieldByName(fieldName), fieldType, defaultStr)
		}
		lowerFieldName := strings.ToLower(fieldName)
		if mapVal, ok := extMap[lowerFieldName]; ok {
			afterTrimMapVal := strings.TrimSpace(mapVal)
			setFieldValue(appInfoVal.FieldByName(fieldName), fieldType, afterTrimMapVal)
		}
	}
}
func setFieldValue(field reflect.Value, typ reflect.Type, val string) {
	typeStr := typ.String()
	if typeStr == "string" {
		field.Set(reflect.ValueOf(val))
	} else {
		if val != "" {
			if typeStr == "int" {
				intVal, err := strconv.Atoi(val)
				if err == nil {
					field.Set(reflect.ValueOf(intVal))
				}
			} else if typeStr == "int64" {
				int64Val, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					field.Set(reflect.ValueOf(int64Val))
				}
			} else if typeStr == "bool" {
				boolVal, err := strconv.ParseBool(val)
				if err == nil {
					field.Set(reflect.ValueOf(boolVal))
				}
			}
		}
	}
}

func GetAppInfo(appkey string) *AppInfo {
	val, ok := appInfoCache.GetByCreator(appkey)
	if ok {
		return val.(*AppInfo)
	} else {
		return nil
	}
}
