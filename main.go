package main

import (
	"io"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 업로드된 파일을 저장할 디렉터리
	uploadDir := "./uploads"

	// 대용량 파일 업로드를 처리하는 엔드포인트
	r.POST("/upload", func(c *gin.Context) {
		// 업로드된 파일 가져오기
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"message": "파일 업로드 실패"})
			return
		}
		defer file.Close()

		targetFileName := "uploads/" + header.Filename
		// 대상 파일 생성
		outFile, err := os.OpenFile(targetFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			c.JSON(500, gin.H{"message": "파일 열기 또는 생성 실패"})
			return
		}

		defer outFile.Close()

		// 파일 복사
		_, err = io.Copy(outFile, file)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "파일 복사 실패"})
			return
		}

		c.JSON(200, gin.H{"message": "파일 업로드 완료"})
	})

	// 업로드 디렉터리 생성
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	r.Run(":8080")
}
