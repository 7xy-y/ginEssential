package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"xietong.me/ginessential/common"
	"xietong.me/ginessential/model"
	"xietong.me/ginessential/response"
	"xietong.me/ginessential/util"
	//"github.com/giorgisio/goav/avformat"
)

func Test(ctx *gin.Context) {
	response.Response(ctx, http.StatusAccepted, 200, nil, "上传成功")
	return
}

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	password := requestUser.Password
	//数据验证
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称为空给一个随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Password: string(hasePassowrd),
	}
	DB.Create(&newUser)
	//发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func Signup(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)
	//获取参数
	email := requestUser.Email
	name := requestUser.Name
	password := requestUser.Password
	//数据验证
	if len(name) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能少于6位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称为空给一个随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if isEmailExist(DB, email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱已存在")
		return
	}
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Password: string(hasePassowrd),
		Email:    email,
	}
	DB.Create(&newUser)
	//发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func UpLoad(ctx *gin.Context) {
	DB := common.GetDB()
	var requestMission = model.Mission{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMission)
	ctx.Bind(&requestMission)
	//获取参数
	file := requestMission.File
	tag := requestMission.Tag
	username := requestMission.Username
	/*width := 640
	height := 360
	cmd := exec.Command("ffmpeg", "-i", file, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		panic("could not generate frame")
	}*/

	//数据验证
	newMission := model.Mission{
		File:     file,
		Tag:      tag,
		Username: username,
	}
	DB.Create(&newMission)
	//发送token
	/*token, err := common.ReleaseToken(newMission)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}*/

	//返回结果
	response.Response(ctx, http.StatusAccepted, 200, nil, "上传成功")
}

func Send(ctx *gin.Context) {
	DB := common.GetDB()
	var requestMission_tag = model.Mission_tag{}

	ctx.Bind(&requestMission_tag)

	tag := requestMission_tag.Tag
	mission_id := requestMission_tag.Mission_ID
	publisher := requestMission_tag.Publisher
	solver := requestMission_tag.Solver
	what := requestMission_tag.What

	newMission_tag := model.Mission_tag{
		Tag:        tag,
		Mission_ID: mission_id,
		Solver:     solver,
		Publisher:  publisher,
		What:       what,
	}

	DB.Create(&newMission_tag)

	response.Response(ctx, http.StatusAccepted, 200, nil, "上传成功")
}

func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	name := c.PostForm("Name")
	password := c.PostForm("Password")
	//数据验证
	if len(password) < 6 {
		response.Response(c, http.StatusAccepted, 421, nil, "密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	db.Where("name=?", name).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusAccepted, 423, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusAccepted, 401, nil, "密码错误")
		//c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//发送token
	/*token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		//log.Printf("token generate error:%v", err)
	}*/

	//返回结果
	response.Response(c, http.StatusAccepted, 200, nil, "登录成功")
}

func Myinfo(c *gin.Context) {
	db := common.GetDB()

	username := c.PostForm("Username")

	var mission_tag []model.Mission_tag
	db.Where("Publisher=?", username).Find(&mission_tag)

	response.Response(c, http.StatusAccepted, 200, gin.H{"mission_tag": mission_tag}, "请求成功")
}

func Info(c *gin.Context) {
	db := common.GetDB()

	var mission []model.Mission
	db.Find(&mission)
	//判断密码是否正

	//发送token
	/*token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		//log.Printf("token generate error:%v", err)
	}*/

	//返回结果
	response.Response(c, http.StatusAccepted, 200, gin.H{"mission": mission}, "请求成功")
}

/*func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}*/

/*func Info(ctx *gin.Context) {
	mission, _ := ctx.Get("user")

	ctx.JSON(http.StatusAccepted, gin.H{"code": 200, "data": gin.H{"mission": dto.ToMissionDto(mission.(model.Mission))}})
}*/

func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email=?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
