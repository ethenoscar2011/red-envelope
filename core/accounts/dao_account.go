package accounts

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type AccountDao struct {
	runner *dbx.TxRunner
}

//查询数据库持久化对象的单实例，获取一行数据
func (dao *AccountDao) GetOne(accountNo string) *Account {
	a := &Account{AccountNo: accountNo}
	ok, err := dao.runner.GetOne(a)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return a
}

//通过用户id和账户类型查询账户信息
func (dao *AccountDao) GetByUserId(userId string, accountType int) *Account {
	a := &Account{}
	sql := "select * from account where user_id=? and account_type=?"
	ok, err := dao.runner.Get(a, sql, userId, accountType)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return a
}

//账户数据的插入
func (dao *AccountDao) Insert(a *Account) (id int64, err error) {
	//rs有两个属性，id和影响行数
	rs, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

//账户余额的更新
func (dao *AccountDao) UpdateBalance(accountNO string, amount decimal.Decimal) (rows int64, err error) {
	sql := "update account set balance=balance+CAST(? AS DECIMAL(30,6))" +
		"where account_no=? and balance>=-1*CAST(? AS DECIMAL(30,6))"
	rs, err := dao.runner.Exec(sql, amount.String(), accountNO, amount.String())
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
//账户状态更新
func (dao *AccountDao) UpdateStatus(accountNO string,status int)(row int64,err error)  {
	sql := "update account set status=? where account_no=?"
	rs, err := dao.runner.Exec(sql,accountNO,status)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

