package abf

import (
	"OTUS_hws/Anti-BruteForce/internal/config"
	"OTUS_hws/Anti-BruteForce/internal/redisdb"
)

type Abf struct {
	RedisServer *redisdb.RedisClient

	LimitIP       int
	LimitLogin    int
	LimitPassword int
}

func New(r *redisdb.RedisClient, conf config.Config) *Abf {

	abf := &Abf{
		RedisServer:   r,
		LimitIP:       conf.Parameters.K,
		LimitLogin:    conf.Parameters.N,
		LimitPassword: conf.Parameters.M,
	}
	return abf
}

func (abf *Abf) CheckRequest() (bool, error) {

	var ok bool
	return ok, nil
}

func (abf *Abf) CheckIP() {

}

func (abf *Abf) CheckLogin() {

}

func (abf *Abf) CheckPassword() {

}

func (abf *Abf) CheckInBlackList() {

}

func (abf *Abf) CheckInWhiteList() {

}
