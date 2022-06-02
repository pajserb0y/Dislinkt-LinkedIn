package repo

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"requests-service/model"

	pb "github.com/MihajloMarjanski/xws-project/common/proto/user_service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RequestsRepository struct {
	db *gorm.DB
}

type IsPublicModel struct {
	Id uint
}

func New() (*RequestsRepository, error) {

	repo := &RequestsRepository{}

	//dsn := "host=requestdb user=XML password=ftn dbname=XML_REQUESTS port=5432 sslmode=disable"
	dsn := "host=localhost user=XML password=ftn dbname=XML_REQUESTS port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	repo.db = db
	repo.db.AutoMigrate(&model.Request{}, &model.Connection{}, &model.Message{})

	return repo, nil
}

func (repo *RequestsRepository) Close() error {
	db, err := repo.db.DB()
	if err != nil {
		return err
	}

	db.Close()
	return nil
}

func (repo *RequestsRepository) GetAllByRecieverId(rid uint) []model.Request {
	var request model.Request
	var requests []model.Request
	repo.db.Model(&request).Where("receiver_id  = ?", rid).Find(&requests)
	return requests
}

func (repo *RequestsRepository) AcceptRequest(sid, rid uint) {
	request := model.Request{
		SenderID:   sid,
		ReceiverID: rid,
	}
	repo.db.Delete(&request)

	connection := model.Connection{
		UserOne: sid,
		UserTwo: rid,
	}

	repo.db.Create(&connection)
}

func (repo *RequestsRepository) DeclineRequest(sid, rid uint) {
	request := model.Request{
		SenderID:   sid,
		ReceiverID: rid,
	}
	repo.db.Delete(&request)
}

func (repo *RequestsRepository) SendRequest(sid, rid uint) {
	conn, err := grpc.Dial("localhost:8100", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	response, err := client.GetPrivateStatusForUserId(context.Background(), &pb.GetPrivateStatusForUserIdRequest{Id: int64(rid)})
	if err != nil {
		panic(err)
	}

	isPrivate := response.IsPrivate
	if isPrivate {
		request := model.Request{
			SenderID:   sid,
			ReceiverID: rid,
		}

		repo.db.Create(&request)
	} else {
		connection := model.Connection{
			UserOne: sid,
			UserTwo: rid,
		}

		repo.db.Create(&connection)
	}
}

func (repo *RequestsRepository) SendMessage(senderID, receiverID uint, text string) {
	message := model.Message{
		Text:       text,
		SenderId:   senderID,
		ReceiverId: receiverID,
	}
	repo.db.Create(&message)
}

func (repo *RequestsRepository) AreConnected(senderID, receiverID uint) bool {
	var connection model.Connection
	repo.db.Model(&connection).Where("user_one = ? AND user_two = ? OR user_one = ? AND user_two = ?", senderID, receiverID, receiverID, senderID).Find(&connection)
	if connection.UserTwo == 0 {
		return false
	}
	return true
}

func (repo *RequestsRepository) AreRequested(u1 uint, u2 uint) bool {
	var request model.Request
	repo.db.Model(&request).Where("sender_id = ? AND receiver_id = ? OR sender_id = ? AND receiver_id = ?", u1, u2, u2, u1).Find(&request)
	if request.ReceiverID == 0 {
		return false
	}
	return true
}
