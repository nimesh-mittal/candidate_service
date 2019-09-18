package candidate

import (
	"encoding/json"

	"github.com/dunzoit/dunzo-microservices/mapping_service/src/main/mappers/constants"

	"github.com/nitishm/go-rejson"

	"github.com/gomodule/redigo/redis"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

type CandidateRedisContext struct {
	DB *redis.Pool
}

func NewCandidateRedisContext(address string, password string, dbName string) (*CandidateRedisContext, error) {
	pool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	defer logrus.Info("redis setup completed")
	return &CandidateRedisContext{DB: pool}, nil
}

func (ctx *CandidateRedisContext) SafeClose() {
	err := ctx.DB.Close()

	if err != nil {
		logrus.Fatal("failed to close redis connection pool", err)
	}
}

func (ctx *CandidateRedisContext) ListCandidates(limit int64, offset int64) (*[]Candidate, error) {
	// TODO: pending implementation

	return nil, nil
}

// TODO: make it efficient, current implementation is very hacky
func (ctx *CandidateRedisContext) GetCandidate(cid string) (*Candidate, error) {
	conn := ctx.DB.Get()
	defer conn.Close()

	rh := rejson.NewReJSONHandler()
	rh.SetRedigoClient(conn)
	studentJSON, err := redis.Bytes(rh.JSONGet(cid, "."))

	candidate := &Candidate{}
	err = json.Unmarshal(studentJSON, candidate)

	if err != nil {
		return nil, err
	}

	return candidate, nil
}

func (ctx *CandidateRedisContext) CreateCandidate(candidate *Candidate) (*Candidate, error) {
	conn := ctx.DB.Get()
	defer conn.Close()

	rh := rejson.NewReJSONHandler()
	rh.SetRedigoClient(conn)
	res, err := rh.JSONSet(candidate.ID, ".", candidate)

	if err != nil {
		return nil, err
	}

	if res.(string) == "OK" {
		return candidate, nil
	}

	return nil, redis.Error("unable to perform set")
}

func (ctx *CandidateRedisContext) UpdateCandidate(cid string, key string, value string) (string, error) {
	conn := ctx.DB.Get()
	defer conn.Close()

	rh := rejson.NewReJSONHandler()
	rh.SetRedigoClient(conn)
	res, err := rh.JSONSet(cid, "."+key, value)

	if err != nil {
		return constants.Empty, err
	}

	if res.(string) == "OK" {
		return cid, nil
	}

	return constants.Empty, redis.Error("unable to perform set")
}

func (ctx *CandidateRedisContext) DeleteCandidate(cid string) (*Candidate, error) {
	conn := ctx.DB.Get()
	defer conn.Close()

	candidate, err := ctx.GetCandidate(cid)

	if err != nil {
		return nil, err
	}

	rh := rejson.NewReJSONHandler()
	rh.SetRedigoClient(conn)
	res, err := rh.JSONDel(cid, ".")

	if err != nil {
		return nil, err
	}

	if res.(string) == "OK" {
		return candidate, nil
	}

	return nil, redis.Error("unable to perform del")
}
