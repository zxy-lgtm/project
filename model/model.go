package model

type Students struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	IdentityCard   string `json:"identity_card"`
	School         string `json:"school"`
	Major          string `json:"major"`
	Tel            string `json:"tel"`
	Grade          string `json:"grade"`
	Subject        string `json:"subject"`
	AwardSituation string `json:"award_situation"`
	Type           int    `json:"type"`   //授课方式，1为线上，2为线下，3为两者皆可
	Gender         int    `json:"gender"` //1为男，2为女
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
