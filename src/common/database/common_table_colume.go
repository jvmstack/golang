package database

import "time"

type CommonTableData struct {
	CreateTime time.Time `json:"createTime" description:"创建时间"`
}
