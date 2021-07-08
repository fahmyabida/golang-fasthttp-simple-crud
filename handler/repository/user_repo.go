package repository

import (
	"fasthttp_crud/handler"
	"fasthttp_crud/model"
	"fasthttp_crud/util/log"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	dbClient *pg.DB
}

func NewUserRepo(dbClient *pg.DB) handler.IUserRepo {
	return &UserRepo{dbClient}
}

func (r *UserRepo) GetAll(traceId string, limit, offset int) (list []model.User, total int, err error) {
	query := r.dbClient.Model(&list)
	if limit != 0 {
		query.Limit(limit)
	}
	if offset != 0 {
		query.Offset(offset)
	}
	total, err = query.SelectAndCount()
	if err != nil {
		log.Error(err, traceId)
	}
	return list, total, err
}

func (r *UserRepo) GetUserById(traceId, userId string) (user model.User, err error) {
	err = r.dbClient.Model(&user).Where("id = ?", userId).Select()
	if err != nil {
		log.Error(err, traceId)
	}
	return user, err
}

func (r *UserRepo) GetUserByUsername(traceId, username string) (user model.User, err error) {
	err = r.dbClient.Model(&user).Where("username = ?", username).Select()
	if err != nil {
		log.Error(err, traceId)
	}
	return user, err
}

func (r *UserRepo) AddNewUser(traceId string, user model.User) (inserted bool, err error) {
	inserted, err = r.dbClient.Model(&user).Where("username = ?", user.Username).SelectOrInsert()
	if err != nil {
		log.Error(err, traceId)
	} else if !inserted {
		err = fmt.Errorf("data already exist with id %v", user.Id)
		log.Error(err, traceId)
	}
	return inserted, err
}

func (r *UserRepo) UpdateUserById(traceId string, user model.User) (err error) {
	tx, err := r.dbClient.Begin()
	if err != nil {
		log.Error(err, traceId)
		return err
	}
	defer tx.Close()
	res, err := tx.Model(&user).Where("id = ?", user.Id).UpdateNotZero()
	if err != nil {
		log.Error(err, traceId)
		tx.Rollback()
		return err
	} else if res.RowsAffected() != 1 {
		err = fmt.Errorf("data update failed, row affected is %v", res.RowsAffected())
		log.Error(err, traceId)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *UserRepo) DeleteUserById(traceId, userId string) (err error) {
	tx, err := r.dbClient.Begin()
	if err != nil {
		log.Error(err, traceId)
		return err
	}
	defer tx.Close()
	res, err := tx.Model(new(model.User)).Where("id = ?", userId).Delete()
	if err != nil {
		log.Error(err, traceId)
		tx.Rollback()
		return err
	} else if res.RowsAffected() != 1 {
		err = fmt.Errorf("data deleted fail, row affected is %v", res.RowsAffected())
		log.Error(err, traceId)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
