package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"teamotp/internal/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xlzd/gotp"
)

func init() {
	loadConfig("data/config.toml")
	init_db("data/teamotp.db")
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	session := sessions.Default(c)

	if LdapAuth(username, password) {
		session.Set("username", username)
		session.Save()
		c.Redirect(http.StatusSeeOther, "/otp")
	} else {
		session.Set("error", "Invalid username or password")
		session.Save()
		c.Redirect(http.StatusFound, "/")
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/")
}

func AddOtp(c *gin.Context) {
	if Unauthorized_404(c) {
		return
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	otpURL := string(bodyBytes)
	otp, err := models.OTPCodeFromURL(otpURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := otp.SaveToDB(); err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusSeeOther, "/otp")
}

func GetOtpValue(c *gin.Context) {
	if Unauthorized_404(c) {
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	otp, err := models.GetOtpById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get OTP code"})
		return
	}

	totp := gotp.NewDefaultTOTP(otp.Secret)
	c.String(http.StatusOK, totp.Now())
}

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	r.Use(sessions.Sessions("session", store))

	r.Use(Default_404())

	r.Static("/static", "web/static")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", func(c *gin.Context) {
		if IsAuthenticated(c) {
			c.Redirect(http.StatusSeeOther, "/otp")
			return
		}

		session := sessions.Default(c)
		errMsg := session.Get("error")
		session.Delete("error")
		session.Save()

		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": errMsg,
		})
	})

	r.GET("/otp", func(c *gin.Context) {
		if Unauthorized_404(c) {
			return
		}

		var otps []models.OTPCode
		o := orm.NewOrm()
		_, err := o.QueryTable("otp_code").All(&otps)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.HTML(http.StatusOK, "otp.html", gin.H{
			"otps": otps,
		})
	})

	r.GET("/otp/qrscan", func(c *gin.Context) {
		if Unauthorized_404(c) {
			return
		}

		c.HTML(http.StatusOK, "qrscan.html", nil)
	})

	r.POST("/login", Login)
	r.GET("/logout", Logout)

	r.POST("/otp/add", AddOtp)
	r.GET("/otp/:id", GetOtpValue)

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	err := r.RunTLS(fmt.Sprintf(":%d", 6443), "data/cert.pem", "data/key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
