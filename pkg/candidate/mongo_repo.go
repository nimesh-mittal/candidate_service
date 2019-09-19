package candidate

import (
	"candidate_service/pkg/commons"
	"context"
	"errors"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

// MongoDBContext holds db object
type MongoDBContext struct {
	DB *mongo.Client
}

// NewCandidateMongoDBContext creates repo object
func NewCandidateMongoDBContext(url string) (*MongoDBContext, error) {
	ctx1, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx1, url)

	if err != nil {
		logrus.Fatal("failed to connect database", url)
		return nil, err
	}

	// Check the connection
	ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx2, nil)

	if err != nil {
		logrus.Fatal("failed to ping database", err)
	}

	defer logrus.Info("mongo database setup completed")
	return &MongoDBContext{DB: client}, nil
}

// SafeClose gets called when program terminates
func (ctx *MongoDBContext) SafeClose() {
	err := ctx.DB.Disconnect(context.TODO())

	if err != nil {
		logrus.Fatal("failed to close database connection", err)
	}
}

// ListCandidates gets candidates from db
func (ctx *MongoDBContext) ListCandidates(fCtx *commons.FlowContext, limit int, offset int) (*[]Candidate, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	collection := ctx.DB.Database(commons.CandidateDb).Collection(commons.CandidateColl)

	// TODO: take filter as parameter
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter, findOptions)

	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	var candidates []Candidate
	for cursor.Next(context.TODO()) {
		var c Candidate
		err := cursor.Decode(&c)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, c)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &candidates, nil
}

// TODO: make it efficient, current implementation is very hacky
// GetCandidate gets candidates from db
func (ctx *MongoDBContext) GetCandidate(fCtx *commons.FlowContext, cid string) (*Candidate, error) {
	findOptions := options.Find()
	findOptions.SetLimit(1)
	findOptions.SetSkip(0)

	collection := ctx.DB.Database(commons.CandidateDb).Collection(commons.CandidateColl)
	filter := bson.D{{Key: "id", Value: cid}}
	cursor, err := collection.Find(context.TODO(), filter, findOptions)

	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	var candidates []Candidate
	for cursor.Next(context.TODO()) {
		var c Candidate
		err := cursor.Decode(&c)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, c)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(candidates) <= 0 {
		return nil, errors.New("record not found")
	}

	return &candidates[0], nil
}

// CreateCandidate inserts a candidate in db
func (ctx *MongoDBContext) CreateCandidate(fCtx *commons.FlowContext, candidate *Candidate) (*Candidate, error) {
	collection := ctx.DB.Database(commons.CandidateDb).Collection(commons.CandidateColl)
	mctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := collection.InsertOne(mctx, candidate)

	if err != nil {
		return nil, err
	} else {
		logrus.WithField(commons.TrackingID, fCtx).Info("created record with id ", resp.InsertedID)
	}

	return candidate, nil
}

// GetBSON converts a candidate in to BSON type
func GetBSON(entity *Candidate) *bson.D {
	// TODO: use reflect
	name := bson.E{Key: "name", Value: entity.Name}
	email := bson.E{Key: "email", Value: entity.Email}
	mobile := bson.E{Key: "mobile", Value: entity.Mobile}
	rollNumber := bson.E{Key: "roll_number", Value: entity.RollNumber}
	age := bson.E{Key: "age", Value: entity.Age}

	return &bson.D{name, email, mobile, rollNumber, age}
}

// UpdateCandidate updates candidate in db by candidate id
func (ctx *MongoDBContext) UpdateCandidate(fCtx *commons.FlowContext, cid string, entity *Candidate) (string, error) {
	collection := ctx.DB.Database(commons.CandidateDb).Collection(commons.CandidateColl)

	filter := bson.D{{Key: "id", Value: cid}}
	new := bson.D{{Key: "$set", Value: GetBSON(entity)}}
	res, err := collection.UpdateOne(context.TODO(), filter, new)

	if err != nil {
		return commons.Empty, err
	} else {
		logrus.WithField(commons.TrackingID, fCtx).
			Info("updated MatchedCount, ModifiedCount, UpsertedCount, UpsertedID ",
				res.MatchedCount, res.ModifiedCount, res.UpsertedCount, res.UpsertedID)
	}

	return res.UpsertedID.(string), nil
}

// DeleteCandidate deletes candidate in db by candidate id
func (ctx *MongoDBContext) DeleteCandidate(fCtx *commons.FlowContext, cid string) (*Candidate, error) {
	collection := ctx.DB.Database(commons.CandidateDb).Collection(commons.CandidateColl)

	if cid == commons.Empty {
		logrus.Error("candidate id is empty")
		return nil, errors.New("candidate id is empty")
	}

	candidate, err := ctx.GetCandidate(fCtx, cid)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	filter := bson.D{{Key: "id", Value: cid}}
	ctx1, _ := context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx1, filter)

	if err != nil {
		logrus.Error(err)
		return nil, err
	} else {
		logrus.WithField(commons.TrackingID, fCtx).Info("number of candidate deleted is ", deleteResult.DeletedCount)
	}
	return candidate, nil
}
