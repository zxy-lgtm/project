package model

type Students struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IdentityCard string `json:"identity_card"`
	School       string `json:"school"`
	Major        string `json:"major"`
	Tel          string `json:"tel"`
	Grade        string `json:"grade"`
	Subject      string `json:"subject"`
	Type         int    `json:"type"`   //授课方式，1为线上，2为线下，3为两者皆可
	Gender       int    `json:"gender"` //1为男，2为女
	Notes        string `json:"notes"`
	Style        string `json:"style"`
	Money        int    `json:"money"`
}

type CommentInfo struct {
	ID int `json:"id"`
	//评级等级
	Name int `json:"name"`
	//评价总结
	All string `json:"all"`
	//1为心理素质，2为仪态仪表，3为语言表达，4为思维品质，5为教学设计，6为教学实验，7为教学评价
	C1       int    `json:"c1"`
	C2       int    `json:"c2"`
	C3       int    `json:"c3"`
	C4       int    `json:"c4"`
	C5       int    `json:"c5"`
	C6       int    `json:"c6"`
	C7       int    `json:"c7"`
	W1       string `json:"w1"`
	W2       string `json:"w2"`
	W3       string `json:"w3"`
	W4       string `json:"w4"`
	W5       string `json:"w5"`
	W6       string `json:"w6"`
	W7       string `json:"w7"`
	PersonID int    `json:"person_id"`
}

type Time struct {
	ID       int    `json:"id"`
	Name     string `json："name"`
	time     string `json:"time"`
	PersonID int    `json:"person_id"`
}

type Award struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PersonID int    `json:"person_id"`
}

type Teach struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Time     string `json:"time"`
	Result   string `json:"result"`
	PersonID int    `json:"person_id"`
}

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	//tips如果是教师则是对这门课的描述，如果是学生则是对这门课的要求
	Tips     string `json:"tips"`
	PersonID int    `json:"person_id"`
}

//将评级等级改变为A/B/C/D/F
func StrName(Name int) string {
	if Name > 80 && Name <= 100 {
		return "A"
	} else if Name > 60 && Name <= 80 {
		return "B"
	} else if Name > 40 && Name <= 60 {
		return "C"
	} else if Name > 20 && Name <= 40 {
		return "D"
	} else {
		return "F"
	}
}

func StrType(Type int) string {
	switch Type {
	case 1:
		return "线上授课"
	case 2:
		return "线下授课"
	case 3:
		return "两者皆可"
	default:
		return "还未选择授课方式"
	}
}

func StrGender(Gender int) string {
	switch Gender {
	case 1:
		return "男"
	case 2:
		return "女"
	default:
		return "未知"
	}
}

type Patriarchs struct {
	ID                 int    `json:"id"`
	NameParents        string `json:"name_p"`
	NameChild          string `json:"name_c"`
	IdentityCardParent string `json:"identity_card_p"`
	IdentityCardChild  string `json:"identity_card_c"`
	School             string `json:"school"`
	Tel                string `json:"tel"`
	Grade              string `json:"grade"`
	ScoreSituation     string `json:"score_situation"`
	GenderParents      int    `json:"gender_p"` //1为男，2为女
}

type ClassInfo struct {
	ID           int    `json:"id"`
	StudentsName string `json:"students_name"`
	Name         string `json:"name"`
	Money        int    `json:"money"`
	Tips         string `json:"tips"`
	Tips2        string `json:"tips2"`
	//状态码，1为精品课（收费），0为免费
	Status int `json:"status"`
}

type UsersInfo struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IdentityCard string `json:"identity_card"`
	Gender       int    `json:"gender"` //1为男，2为女
	Account      string `json:"account"`
	Password     string `json:"password"`
}
