package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/liju-github/EcommerceAdminService/config"
	"github.com/liju-github/EcommerceAdminService/proto/admin"
	"github.com/liju-github/EcommerceAdminService/proto/content"
	"github.com/liju-github/EcommerceAdminService/proto/user"
	"google.golang.org/grpc"
)

type AdminService struct {
	admin.UnimplementedAdminServiceServer
	userClient    user.UserServiceClient
	contentClient content.ContentServiceClient
}

func NewAdminService(userConn *grpc.ClientConn, contentConn *grpc.ClientConn) *AdminService {
	return &AdminService{
		userClient:    user.NewUserServiceClient(userConn),
		contentClient: content.NewContentServiceClient(contentConn),
	}
}
func (as *AdminService) AdminLogin(ctx context.Context, req *admin.AdminLoginRequest) (*admin.AdminLoginResponse, error) {
	cred := config.LoadConfig()
	log.Printf("Expected Username: %s, Expected Password: %s", cred.AdminUsername, cred.AdminPassword)
	log.Printf("Received Username: %s, Received Password: %s", req.Username, req.Password)
	if req.Username != cred.AdminUsername || req.Password != cred.AdminPassword {
		return &admin.AdminLoginResponse{Success: false}, errors.New("login failed")
	}

	return &admin.AdminLoginResponse{Success: true}, nil
}

func (as *AdminService) BanUser(ctx context.Context, req *admin.BanUserRequest) (*admin.BanUserResponse, error) {
	resp, err := as.userClient.BanUser(ctx, &user.BanUserRequest{UserId: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("failed to ban user: %v", err)
	}
	return &admin.BanUserResponse{Success: resp.Success, Message: resp.Message}, nil
}

func (as *AdminService) UnBanUser(ctx context.Context, req *admin.UnBanUserRequest) (*admin.UnBanUserResponse, error) {
	resp, err := as.userClient.UnBanUser(ctx, &user.UnBanUserRequest{UserId: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("failed to unban user: %v", err)
	}
	return &admin.UnBanUserResponse{Success: resp.Success, Message: resp.Message}, nil
}

func (as *AdminService) GetAllFlaggedAnswers(ctx context.Context, req *admin.GetFlaggedAnswersRequest) (*admin.GetFlaggedAnswersResponse, error) {
	log.Println("reached admin service")
	resp, err := as.contentClient.GetFlaggedAnswers(ctx, &content.GetFlaggedAnswersRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get flagged answers: %v", err)
	}

	// Convert []*content.Answer to []*admin.Answer
	flaggedAnswers := make([]*admin.Answer, len(resp.FlaggedAnswers))
	for i, answer := range resp.FlaggedAnswers {
		flaggedAnswers[i] = &admin.Answer{
			Id:         answer.Id,
			QuestionId: answer.QuestionId,
			UserId:     answer.UserId,
			AnswerText: answer.AnswerText,
			Upvotes:    answer.Upvotes,
			Downvotes:  answer.Downvotes,
			IsFlagged:  answer.IsFlagged,
			CreatedAt:  answer.CreatedAt,
			UpdatedAt:  answer.UpdatedAt,
		}
	}

	return &admin.GetFlaggedAnswersResponse{
		FlaggedAnswers:      flaggedAnswers,
		TotalFlaggedAnswers: resp.TotalFlaggedAnswers,
	}, nil
}

func (as *AdminService) GetAllFlaggedQuestions(ctx context.Context, req *admin.GetFlaggedQuestionsRequest) (*admin.GetFlaggedQuestionsResponse, error) {
	resp, err := as.contentClient.GetFlaggedQuestions(ctx, &content.GetFlaggedQuestionsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get flagged questions: %v", err)
	}

	// Convert []*content.Question to []*admin.Question
	flaggedQuestions := make([]*admin.Question, len(resp.FlaggedQuestions))
	for i, question := range resp.FlaggedQuestions {
		flaggedQuestions[i] = &admin.Question{
			QuestionID: question.QuestionID,
			Question:   question.Question,
			UserID:     question.UserID,
			CreatedAt:  question.CreatedAt,
			Tags:       question.Tags,
			IsAnswered: question.IsAnswered,
			Details:    question.Details,
		}
	}

	return &admin.GetFlaggedQuestionsResponse{
		FlaggedQuestions:      flaggedQuestions,
		TotalFlaggedQuestions: resp.TotalFlaggedQuestions,
	}, nil
}

func (as *AdminService) GetAllUsers(ctx context.Context, req *admin.GetAllUsersRequest) (*admin.GetAllUsersResponse, error) {
    // Call GetAllUsers in UserService
    userResp, err := as.userClient.GetAllUsers(ctx, &user.GetAllUsersRequest{})
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve users: %v", err)
    }

    // Map the response to AdminService's GetAllUsersResponse format
    adminUsers := make([]*admin.User, len(userResp.Users))
    for i, u := range userResp.Users {
        adminUsers[i] = &admin.User{
            Id:               u.Id,
            Email:            u.Email,
            PasswordHash:     u.PasswordHash,
            Name:             u.Name,
            StreetName:       u.StreetName,
            Locality:         u.Locality,
            State:            u.State,
            Pincode:          u.Pincode,
            PhoneNumber:      u.PhoneNumber,
            Reputation:       u.Reputation,
            VerificationCode: u.VerificationCode,
            IsBanned:         u.IsBanned,
            IsVerified:       u.IsVerified,
        }
    }

    // Return the response with success status and the list of users
    return &admin.GetAllUsersResponse{
        Success: true,
        Users:   adminUsers,
    }, nil
}

