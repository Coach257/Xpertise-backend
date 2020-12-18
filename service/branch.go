package service

import (
	"xpertise-go/global"
	"xpertise-go/model"

	"github.com/jinzhu/gorm"
)

func CreateAComment(comment *model.Comment) (err error) {
	if err = global.DB.Create(&comment).Error; err != nil {
		return err
	}
	return
}

// 查询某条评论
func QueryAComment(commentID uint64) (comment model.Comment, notFound bool) {
	notFound = global.DB.First(&comment, commentID).RecordNotFound()
	return comment, notFound
}

// 列出某个文献的所有评论
func QueryAllComments(paperID string) (comments []model.Comment) {
	global.DB.Where("paper_id = ?", paperID).Find(&comments)
	return comments
}

// 删除某条评论
func DeleteAComment(CommentID uint64) (err error) {
	var comment model.Comment
	notFound := global.DB.First(&comment, CommentID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(&comment).Error
	return err
}

// 置顶某条评论
func PutCommentToTop(commentID uint64) (err error) {
	var comment model.Comment
	notFound := global.DB.First(&comment, commentID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	comment.OnTop = true
	err = global.DB.Save(&comment).Error
	return err
}

// 取消置顶某条评论
func CancelCommentToTop(commentID uint64) (err error) {
	var comment model.Comment
	notFound := global.DB.First(&comment, commentID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	comment.OnTop = false
	err = global.DB.Save(&comment).Error
	return err
}

// 评论点赞数加一/点踩数加一/点赞数减一/点踩数减一
func UpdateLikeOrDislike(comment *model.Comment, method uint64) (err error) {
	switch method {
	case 1:
		comment.Like += 1
	case 2:
		comment.Dislike += 1
	case 3:
		comment.Like -= 1
	case 4:
		comment.Dislike -= 1
	}
	err = global.DB.Save(comment).Error
	return err
}

// 在评论-点赞表中加入一项
func CreateACommentLike(userID uint64, comment *model.Comment, method uint64) (err error) {
	var likeOrDislike bool
	if method == 1 {
		likeOrDislike = true // 点赞
	} else {
		likeOrDislike = false // 点踩
	}
	commentLike := model.CommentLike{UserID: userID, CommentID: comment.CommentID, LikeOrDislike: likeOrDislike}
	if err = global.DB.Create(&commentLike).Error; err != nil {
		return err
	}
	// 创建完之后还要修改相应的评论条目，点赞数+1或点踩数+1
	err = UpdateLikeOrDislike(comment, method)
	return
}

// 查询评论-点赞表(CommentLike)的某一项
func QueryAnItemFromCommentLike(commentID uint64, userID uint64) (commentLike model.CommentLike, notFound bool) {
	notFound = global.DB.Where("comment_id = ?", commentID).Where("user_id = ?", userID).First(&commentLike).RecordNotFound()
	return commentLike, notFound
}

// 转换点赞为点踩/点踩为点赞
func TransferBetweenLikeAndDislike(commentLike *model.CommentLike, comment *model.Comment) error {
	var err1 error
	var err2 error
	if commentLike.LikeOrDislike == true {
		commentLike.LikeOrDislike = false
		err1 = UpdateLikeOrDislike(comment, 3) // 原评论点赞数减一
		if err1 != nil {
			return err1
		}
		err1 = UpdateLikeOrDislike(comment, 2) // 原评论点踩数加一
	} else {
		commentLike.LikeOrDislike = true
		err1 = UpdateLikeOrDislike(comment, 4) // 原评论点踩数减一
		if err1 != nil {
			return err1
		}
		err1 = UpdateLikeOrDislike(comment, 1) // 原评论点赞数加一
	}
	err2 = global.DB.Save(commentLike).Error
	if err1 != nil {
		return err1
	}
	return err2
}

// 找到authorid所属连通块的根
func GetFa(AuthorID string) (connection model.Connection, notFound bool) {
	notFound = global.DB.Where("author1_id = ?", AuthorID).Or("author2_id = ?", AuthorID).First(&connection).RecordNotFound()
	return connection, notFound
}

// 获得图
func GetAuthorConnectionGraph(FaID string) (connection []model.Connection) {
	global.DB.Where("father_id = ?", FaID).Find(&connection)
	return connection
}

// QueryAllReferences 列出某个文献的所有参考文献
func QueryAllReferences(paperID string) (references []model.PaperReference) {
	global.DB.Where(&model.PaperReference{PaperID: paperID}).Find(&references)
	return references
}
