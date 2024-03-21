package odinNode

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OdinNodeModel = (*customOdinNodeModel)(nil)

type (
	// OdinNodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOdinNodeModel.
	OdinNodeModel interface {
		odinNodeModel
		withSession(session sqlx.Session) OdinNodeModel
	}

	customOdinNodeModel struct {
		*defaultOdinNodeModel
	}
)

// NewOdinNodeModel returns a model for the database table.
func NewOdinNodeModel(conn sqlx.SqlConn) OdinNodeModel {
	return &customOdinNodeModel{
		defaultOdinNodeModel: newOdinNodeModel(conn),
	}
}

func (m *customOdinNodeModel) withSession(session sqlx.Session) OdinNodeModel {
	return NewOdinNodeModel(sqlx.NewSqlConnFromSession(session))
}
