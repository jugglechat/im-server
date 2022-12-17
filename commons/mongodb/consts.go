package mongodb

// const (
// 	MongoDbName = "zhaomimongo"
// )

var MongoDbName string

type MongoError struct {
	Msg string
}

func (e *MongoError) Error() string {
	return e.Msg
}

//0:待审核；1:审核通过；2:审核未通过;3:严重违规被审核系统删除；
const (
	AuditStatus_PENDING     = 0
	AuditStatus_PASS        = 1
	AuditStatus_NOPASS      = 2
	AuditStatus_AUDITDELETE = 3
)
