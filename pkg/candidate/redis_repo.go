package candidate

import (
	"encoding/json"

	"github.com/dunzoit/dunzo-microservices/mapping_service/src/main/mappers/constants"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson"
	"github.com/sirupsen/logrus"
)

// RedisContext holds redis db object
type RedisContext struct {
	DB *redis.Pool
}

// NewCandidateRedisContext creates Redis context for candidate
func NewCandidateRedisContext(address string, password string, dbName string) (*RedisContext, error) {
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
	return &RedisContext{DB: pool}, nil
}

// SafeClose gets called on program termination
func (ctx *RedisContext) SafeClose() {
	err := ctx.DB.Close()

	if err != nil {
		logrus.Fatal("failed to close redis connection pool", err)
	}
}

// ListCandidates list all candidates from redis
func (ctx *RedisContext) ListCandidates(limit int64, offset int64) (*[]Candidate, error) {
	// TODO: pending implementation

	return nil, nil
}

// GetCandidate get candidate by id
// TODO: make it efficient, current implementation is very hacky
func (ctx *RedisContext) GetCandidate(cid string) (*Candidate, error) {
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

// CreateCandidate create candidate in redis
func (ctx *RedisContext) CreateCandidate(candidate *Candidate) (*Candidate, error) {
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

// UpdateCandidate update candidate
func (ctx *RedisContext) UpdateCandidate(cid string, key string, value string) (string, error) {
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

// DeleteCandidate delete candidate by id
func (ctx *RedisContext) DeleteCandidate(cid string) (*Candidate, error) {
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
