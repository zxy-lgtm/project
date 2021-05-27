package model

import (
	"log"
)

//const dsn = "root:123456@/PICKUP?charset=utf8&parseTime=True&loc=Local"

func CreateU(tmpUser LoginInfo) bool {

	if err := Db.Self.Create(&tmpUser).Error; err != nil {
		log.Println("creat user err:", err)
		return false
	}

	return true
}

func CreateT(tmpUser UsersInfo) error {

	var tooluser = Students{
		ID: tmpUser.ID,
	}
	err2 := Db.Self.Where(&Students{ID: tmpUser.ID}).Find(&tooluser).Error
	if err2 == nil {
		log.Println("creat user err: ", err2)
		return err2
	}

	Db.Self.Create(&tmpUser)

	return nil
}

func CreateP(tmpUser UsersInfo) error {

	var tooluser = Patriarchs{
		ID:                 tmpUser.ID,
		IdentityCardParent: tmpUser.IdentityCard,
		GenderParents:      tmpUser.Gender,
	}
	err2 := Db.Self.Where(&Patriarchs{ID: tmpUser.ID}).Find(&tooluser).Error
	if err2 == nil {
		log.Println("creat user err: ", err2)
		return err2
	}
	Db.Self.Create(&tmpUser)

	return nil
}

//通过account寻找user中的用户，找得到则返回该用户信息和true
func FindUserAccount(account string) (UsersInfo, bool) {
	var user UsersInfo
	err := Db.Self.Where(&UsersInfo{Account: account}).First(&user).Error
	if err != nil {
		return user, false
	}
	return user, true
}

func FindStudentsID(id int) (Students, bool) {
	var student Students
	err := Db.Self.Where(&Students{ID: id}).Find(&student).Error
	if err != nil {
		return student, false
	}
	return student, true
}

func FindTimeUID(uid int) ([]Time, bool) {
	var times []Time
	err := Db.Self.Where(&Time{PersonID: uid}).Find(&times).Error
	if err != nil {
		return times, false
	}
	return times, true
}

func FindSubjectUID(uid int) ([]Subject, bool) {
	var subjects []Subject
	err := Db.Self.Where(&Subject{PersonID: uid}).Find(&subjects).Error
	if err != nil {
		return subjects, false
	}
	return subjects, true
}

func FindAwardUID(uid int) ([]Award, bool) {
	var awards []Award
	err := Db.Self.Where(&Award{PersonID: uid}).Find(&awards).Error
	if err != nil {
		return awards, false
	}
	return awards, true
}

func FindTeachUID(uid int) ([]Teach, bool) {
	var teaches []Teach
	err := Db.Self.Where(&Teach{PersonID: uid}).Find(&teaches).Error
	if err != nil {
		return teaches, false
	}
	return teaches, true
}

func FindCommentUID(uid int) (Comment, bool) {
	var comment Comment
	err := Db.Self.Where(&Comment{PersonID: uid}).Find(&comment).Error
	if err != nil {
		return comment, false
	}
	return comment, true
}

func UpdateStudent(tmpstudent Students) bool {

	if err := Db.Self.Model(&Students{}).Where(Students{ID: tmpstudent.ID}).Update(&tmpstudent).Error; err != nil {
		return false
	}
	return true

}

/*func FindUser(uid string) (Users, error) {
	var tmpUser Users
	if err := Db.Self.Model(&Users{}).Where(Users{Sid: uid}).Find(&tmpUser).Error; err != nil {
		return tmpUser, err
	}
	return tmpUser, nil

}

func Updateuser(tmpUser Users, uid string) error {
	if err := Db.Self.Model(&Users{}).Where(Users{Sid: uid}).Error; err != nil {
		return err
	}

	if err := Db.Self.Model(&Users{}).Where(Users{Sid: uid}).Update(&tmpUser).Error; err != nil {
		return err
	}
	return nil

}

func UpdatePwd(tmpChange UpdatePwdInfo, uid string) error {
	tmpUser, err := FindUser(uid)
	if err != nil {
		return err
	}
	num := strings.Compare(tmpChange.Old, tmpUser.Password)

	//fmt.Println(num, tmpChange.Old)
	var myerr handler.Error
	myerr.ErrorCode = "password is not corret"
	myerr.Message = "bad request!"
	if num != 0 {
		return &myerr
	}
	tmpUser.Password = tmpChange.New

	err = Db.Self.Model(&Users{}).Where(Users{Sid: uid}).Update(&tmpUser).Error
	if err != nil {
		return err
	}

	return nil

}

func UpdateCommentD(tmpComment CommentDriver, uid string) error {
	exist, err := FindCommentD(uid)
	if err != nil {
		if err := Db.Self.Model(&CommentDriver{}).Create(&tmpComment).Error; err != nil {
			return err
		}
		return nil
	}
	exist.DriverScore = (exist.DriverScore + tmpComment.DriverScore) / 2.0
	exist.Words = exist.Words + tmpComment.Words
	if err := Db.Self.Model(&CommentDriver{}).Where(CommentDriver{DriverID: uid}).Update(&exist).Error; err != nil {
		return err
	}

	return nil
}

func FindCommentD(uid string) (CommentDriver, error) {
	var tmpComment CommentDriver
	if err := Db.Self.Model(&CommentDriver{}).Where(CommentDriver{DriverID: uid}).First(&tmpComment).Error; err != nil {
		return tmpComment, err
	}
	return tmpComment, nil
}

func UpdateCommentP(tmpComment CommentPassenger, uid string) error {
	exist, err := FindCommentP(uid)
	if err != nil {
		if err := Db.Self.Model(&CommentPassenger{}).Create(&tmpComment).Error; err != nil {
			return err
		}
		return nil
	}
	exist.PassengerScore = (exist.PassengerScore + tmpComment.PassengerScore) / 2.0
	exist.Words = exist.Words + tmpComment.Words
	if err := Db.Self.Model(&CommentPassenger{}).Where(CommentPassenger{PassengerID: uid}).Update(&exist).Error; err != nil {
		return err
	}

	return nil
}

func FindCommentP(uid string) (CommentPassenger, error) {
	var tmpComment CommentPassenger
	if err := Db.Self.Model(&CommentPassenger{}).Where(CommentPassenger{PassengerID: uid}).First(&tmpComment).Error; err != nil {
		return tmpComment, err
	}
	return tmpComment, nil
}

func CreateDriverRt(tmpRt RequireDriver) error {
	if _, err := FindDriverRt(tmpRt.DriverID); err == nil {
		if ok := DeleteDriverRt(tmpRt.DriverID); ok != nil {
			return ok
		}
	}

	if err := Db.Self.Create(&tmpRt).Error; err != nil {
		return err
	}

	return nil
}

func FindDriverRt(uid string) (RequireDriver, error) {
	var tmpRt RequireDriver
	if err := Db.Self.Model(&RequireDriver{}).Where(RequireDriver{DriverID: uid}).First(&tmpRt).Error; err != nil {
		return tmpRt, err
	}

	return tmpRt, nil
}

func DeleteDriverRt(uid string) error {
	tmpRt, err := FindDriverRt(uid)
	if err != nil {
		return err
	}
	if err = Db.Self.Delete(&tmpRt).Error; err != nil {
		return err
	}
	return nil
}

func ConfirmP(uid string) error {
	tmpRt, err := FindPassengerRt(uid)
	if err != nil {
		return err
	}
	tmpRt.Status = 2
	if err := Db.Self.Model(&RequirePassenger{}).Where(RequirePassenger{PassengerID: uid}).Update(&tmpRt).Error; err != nil {
		return err
	}

	return nil
}
func ConfirmD(pid string, uid string) error {
	var tmpRt RequirePassenger
	var myerr handler.Error
	myerr.ErrorCode = "passenger does not confirm"
	myerr.Message = "wait!"
	err := Db.Self.Model(&RequirePassenger{}).Where(RequirePassenger{PassengerID: pid}).First(&tmpRt).Error
	fmt.Println(1)
	if err != nil {
		return err
	}
	if tmpRt.Status == 1 {
		return &myerr
	}
	fmt.Println(2)
	tmprt, err2 := FindDriverRt(uid)
	if err2 != nil {
		return err2
	}
	fmt.Println(3)
	tmprt.Status = 2
	if err := Db.Self.Model(&RequireDriver{}).Where(RequireDriver{DriverID: uid}).Update(&tmprt).Error; err != nil {
		return err
	}

	match := Match{
		DriverID:    tmprt.DriverID,
		PassengerID: tmpRt.PassengerID,
		StartTime:   tmprt.StartTime,
		EndTime:     tmprt.EndTime,
		StartSpot:   tmprt.StartSpot,
		EndSpot:     tmpRt.EndSpot,
		DriverPhone: GetPhone(tmprt.DriverID),
	}

	if err := CreateRoute(match); err != nil {
		return err
	}

	if err := DeleteDriverRt(tmprt.DriverID); err != nil {
		return err
	}
	if err := DeletePassengerRt(tmpRt.PassengerID); err != nil {
		return err
	}

	return nil
}

func GetPhone(uid string) string {
	user, err := FindUser(uid)
	if err != nil {
		return "null"
	}
	return user.Phone
}

func CreatePassengerRt(tmpRt RequirePassenger) error {
	if _, err := FindPassengerRt(tmpRt.PassengerID); err == nil {
		if ok := DeletePassengerRt(tmpRt.PassengerID); ok != nil {
			return ok
		}
	}

	if err := Db.Self.Create(&tmpRt).Error; err != nil {
		return err
	}

	return nil
}

func FindPassengerRt(uid string) (RequirePassenger, error) {
	var tmpRt RequirePassenger
	if err := Db.Self.Model(&RequirePassenger{}).Where(RequirePassenger{PassengerID: uid}).First(&tmpRt).Error; err != nil {
		return tmpRt, err
	}

	return tmpRt, nil
}

func DeletePassengerRt(uid string) error {
	tmpRt, err := FindPassengerRt(uid)
	if err != nil {
		return err
	}
	if err = Db.Self.Delete(&tmpRt).Error; err != nil {
		return err
	}
	return nil
}

func CreateRoute(tmpRoute Match) error {
	err := Db.Self.Model(&Match{}).Create(&tmpRoute).Error
	if err != nil {
		return err
	}
	return nil
}

func FindRoute(uid string) ([]Match, error) {
	var routes []Match
	err := Db.Self.Model(&Match{}).Where(Match{UserID: uid}).Find(&routes).Error
	if err != nil {
		return routes, err
	}
	return routes, nil
}

func FindRoute2(id int) (Match, error) {
	var route Match
	err := Db.Self.Model(&Match{}).Where(Match{ID: id}).Find(&route).Error
	if err != nil {
		return route, err
	}
	return route, nil
}

func DeleteRoute(id int) error {
	var route Match
	err := Db.Self.Model(&Match{}).Where(Match{ID: id}).Delete(&route).Error
	if err != nil {
		return err
	}

	return nil
}
*/
