package server

import (
	"fmt"
	"xpertise-go/dao"
)

// CreateAUser : Create a table user.
func CreateAUser(user *dao.User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}

func QueryAUerById(userid string)(user dao.User,Notfound bool)  {
	fmt.Println("userid:", userid)
	fmt.Println("QueryAUserByID")
	Notfound = dao.DB.First(&user,userid).RecordNotFound()
	return user, Notfound
}

func QueryAUserByUsername(username string) (user dao.User, Notfound bool) {
	fmt.Println("username:", username)
	fmt.Println("QueryAUserByUsername")
	Notfound = dao.DB.Where("username = ?", username).First(&user).RecordNotFound()
	return user, Notfound
}

func QueryAUserByEmail(email string) (user dao.User, Notfound bool) {
	fmt.Println("email", email)
	fmt.Println("QueryAUserByEmail")
	Notfound = dao.DB.Where("email = ?", email).First(&user).RecordNotFound()
	return user, Notfound
}

func UpdateAUserPassword(user *dao.User, newpassword string) error {
	user.Password = newpassword
	err := dao.DB.Save(user).Error
	return err
}

func UpdateAUser(user *dao.User,username string,email string,info string)error{
	user.Username=username
	user.Email=email
	user.BasicInfo=info
	err:=dao.DB.Save(user).Error
	return err
}

func CreateAFolder(foldername string,folderinfo string,user *dao.User)(error,uint64){
	folder := dao.Folder{Foldername: foldername,Folderinfo: folderinfo,UserID: user.UserID}
	err:=dao.DB.Create(&folder).Error
	return err,folder.FolderID
}

func QueryAFolderByID(folderId string)(folder dao.Folder,notFound bool){
	fmt.Println("folderId",folderId)
	fmt.Println("QueryAFolderByID")
	notFound=dao.DB.First(&folder,folderId).RecordNotFound()
	return folder,notFound
}

func CreateAFavorite(folderId uint64,docId uint64,docInfo string)(error,uint64){
	favorite :=dao.Favorite{FolderID: folderId,DocID: docId,Docinfo: docInfo}
	err :=dao.DB.Create(&favorite).Error
	return err,favorite.FavorID
}


func DeleteAStudentByID(StudentID uint64) {
	dao.DB.Where("ID = ?", StudentID).Delete(&dao.Student{})
	return
}

func UpdateAStudentByAge(student *dao.Student, age uint64) {
	dao.DB.Model(&student).Update("Age", age)
}

func QueryAllStudents() (students []*dao.Student) {
	dao.DB.Find(&students)
	return students
}

func QueryStudentByID(StudentID uint64) (student []*dao.Student) {
	dao.DB.First(&student, StudentID)
	return student
}

func QueryStudentsByAge(age uint64) (students []*dao.Student) {
	dao.DB.Where("Age = ?", age).Find(&students)
	return students
}
