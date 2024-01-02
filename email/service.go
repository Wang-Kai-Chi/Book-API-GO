package email

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	. "iknowbook.com/app/data"
	. "iknowbook.com/app/user"
)

type EmailService struct {
	userRepo UserRepository
}

func NewEmailService(userRepo UserRepository) EmailService {
	return EmailService{
		userRepo: userRepo,
	}
}

func handleInternalError(err error, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": "ERROR: " + err.Error(),
		})
	}
}

func mustGetUserFromBody(ctx *gin.Context) User {
	body, err := io.ReadAll(ctx.Request.Body)
	handleInternalError(err, ctx)

	var us User
	err = json.Unmarshal(body, &us)
	handleInternalError(err, ctx)

	return us
}

func (serv EmailService) SendVerificationEmail(ctx *gin.Context) {
	sendVerificationMail := func(receiver string) {
		form := EMail{
			Sender:   "ericwangcatch@gmail.com",
			Subject:  "Iknowbook email Verification",
			HTMLBody: string(MustReadEmailForm()),
		}
		form.Receiver = receiver

		err := SendMail(form)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "A verification mail has sended, please check your email and verify it.",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Sending a verification mail failed, please try again.",
			})
		}
	}

	VerifyUserEmail(ctx,
		func(ctx *gin.Context, user User) {
			users := serv.userRepo.FindUserByEmail(user)
			if len(users) > 0 {
				sendVerificationMail(user.Email)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Response": "Not registered user.",
				})
			}
		},
	)
}

func VerifyUserEmail(ctx *gin.Context, authOp func(*gin.Context, User)) {
	us := mustGetUserFromBody(ctx)

	res := VerifyEmail(us.Email)
	if res != nil && res.Syntax.Valid {
		resStr, _ := json.Marshal(res)
		log.Println(string(resStr))
		authOp(ctx, us)
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": "Email invaild.",
		})
	}
}

func (serv EmailService) VerifyEmail(ctx *gin.Context) {
	VerifyUserEmail(ctx,
		func(ctx *gin.Context, user User) {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "Email verified.",
			})
		},
	)
}
