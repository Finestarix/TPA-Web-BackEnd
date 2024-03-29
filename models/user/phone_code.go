package user

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type PhoneCode struct {
	ID        int    `gorm:"PRIMARY_KEY"`
	Country   string `gorm:"VARCHAR(100); NOT NULL"`
	Code      string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.AutoMigrate(&PhoneCode{})
	log.Println("Initialize PhoneCode Success")
}

func DropTablePhoneCode() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&PhoneCode{})
	database.AutoMigrate(&PhoneCode{})
	log.Println("Drop Database Success")
}

func GetAllPhoneCode() []PhoneCode {
	database := connection.GetConnection()
	defer database.Close()

	var phoneCode []PhoneCode
	if core.ValidateAPIKey() == false {
		return phoneCode
	}
	database.Find(&phoneCode)
	return phoneCode
}

func GetPhoneCode(searchPhoneCode string) PhoneCode {
	database := connection.GetConnection()
	defer database.Close()

	var phoneCode PhoneCode
	if core.ValidateAPIKey() == false {
		return phoneCode
	}
	database.
		Where("code = ?", searchPhoneCode).
		First(&phoneCode)

	return phoneCode
}

func InsertPhoneCode(country string, code string) *PhoneCode {
	database := connection.GetConnection()
	defer database.Close()

	newPhoneCode := &PhoneCode{
		Country: country,
		Code:    code,
	}
	database.Save(newPhoneCode)

	log.Println("Insert New Phone Code Success")
	return newPhoneCode
}

func GeneratePhoneCode() {
	InsertPhoneCode("Afghanistan", "+93")
	InsertPhoneCode("Albania", "+355")
	InsertPhoneCode("Algeria", "+213")
	InsertPhoneCode("American Samoa", "+1684")
	InsertPhoneCode("Andorra", "+376")
	InsertPhoneCode("Angola", "+244")
	InsertPhoneCode("Anguilla", "+1264")
	InsertPhoneCode("Antigua and Barbuda", "+1268")
	InsertPhoneCode("Argentina", "+54")
	InsertPhoneCode("Armenia", "+374")
	InsertPhoneCode("Aruba", "+297")
	InsertPhoneCode("Australia", "+61")
	InsertPhoneCode("Austria", "+43")
	InsertPhoneCode("Azerbaijan", "+994")
	InsertPhoneCode("Bahamas", "+1242")
	InsertPhoneCode("Bahrain", "+973")
	InsertPhoneCode("Bangladesh", "+880")
	InsertPhoneCode("Barbados", "+1246")
	InsertPhoneCode("Belarus", "+375")
	InsertPhoneCode("Belgium", "+32")
	InsertPhoneCode("Belize", "+501")
	InsertPhoneCode("Benin", "+229")
	InsertPhoneCode("Bermuda", "+1441")
	InsertPhoneCode("Bhutan", "+975")
	InsertPhoneCode("Bolivia", "+591")
	InsertPhoneCode("Bosnia and Herzegovina", "+387")
	InsertPhoneCode("Botswana", "+267")
	InsertPhoneCode("Brazil", "+55")
	InsertPhoneCode("British Indian Ocean Territory", "+246")
	InsertPhoneCode("Bulgaria", "+359")
	InsertPhoneCode("Burkina Faso", "+226")
	InsertPhoneCode("Burundi", "+257")
	InsertPhoneCode("Cambodia", "+855")
	InsertPhoneCode("Cameroon", "+237")
	InsertPhoneCode("Canada", "+1")
	InsertPhoneCode("Cayman Islands", "+1345")
	InsertPhoneCode("Central African Republic", "+236")
	InsertPhoneCode("Chad", "+235")
	InsertPhoneCode("Chile", "+56")
	InsertPhoneCode("China", "+86")
	InsertPhoneCode("Christmas Island", "+61")
	InsertPhoneCode("Cocos (Keeling) Islands", "+61")
	InsertPhoneCode("Colombia", "+57")
	InsertPhoneCode("Comoros", "+269")
	InsertPhoneCode("Congo", "+242")
	InsertPhoneCode("Cook Islands", "+682")
	InsertPhoneCode("Costa Rica", "+506")
	InsertPhoneCode("Croatia", "+385")
	InsertPhoneCode("Cuba", "+53")
	InsertPhoneCode("Cyprus", "+357")
	InsertPhoneCode("Czech Republic", "+420")
	InsertPhoneCode("Denmark", "+45")
	InsertPhoneCode("Djibouti", "+253")
	InsertPhoneCode("Dominica", "+1767")
	InsertPhoneCode("Dominican Republic", "+1849")
	InsertPhoneCode("East Timor", "+670")
	InsertPhoneCode("Ecuador", "+593")
	InsertPhoneCode("Egypt", "+20")
	InsertPhoneCode("El Salvador", "+503")
	InsertPhoneCode("Equatorial Guinea", "+240")
	InsertPhoneCode("Eritrea", "+291")
	InsertPhoneCode("Estonia", "+372")
	InsertPhoneCode("Ethiopia", "+251")
	InsertPhoneCode("Falkland Islands", "+500")
	InsertPhoneCode("Faroe Islands", "+298")
	InsertPhoneCode("Finland", "+358")
	InsertPhoneCode("France", "+33")
	InsertPhoneCode("French Guiana", "+594")
	InsertPhoneCode("French Polynesia", "+689")
	InsertPhoneCode("Gabon", "+241")
	InsertPhoneCode("Gambia", "+220")
	InsertPhoneCode("Georgia", "+995")
	InsertPhoneCode("Germany", "+49")
	InsertPhoneCode("Ghana", "+233")
	InsertPhoneCode("Gibraltar", "+350")
	InsertPhoneCode("Greece", "+30")
	InsertPhoneCode("Greenland", "+299")
	InsertPhoneCode("Grenada", "+1473")
	InsertPhoneCode("Guadeloupe", "+590")
	InsertPhoneCode("Guam", "+1671")
	InsertPhoneCode("Guatemala", "+502")
	InsertPhoneCode("Guinea", "+224")
	InsertPhoneCode("Guinea-Bissau", "+245")
	InsertPhoneCode("Guyana", "+592")
	InsertPhoneCode("Haiti", "+509")
	InsertPhoneCode("Honduras", "+504")
	InsertPhoneCode("Hong Kong", "+852")
	InsertPhoneCode("Hungary", "+36")
	InsertPhoneCode("Iceland", "+354")
	InsertPhoneCode("India", "+91")
	InsertPhoneCode("Indonesia", "+62")
	InsertPhoneCode("Iran", "+98")
	InsertPhoneCode("Iraq", "+964")
	InsertPhoneCode("Ireland", "+353")
	InsertPhoneCode("Israel", "+972")
	InsertPhoneCode("Italy", "+39")
	InsertPhoneCode("Ivory Coast", "+225")
	InsertPhoneCode("Jamaica", "+1876")
	InsertPhoneCode("Japan", "+81")
	InsertPhoneCode("Jordan", "+962")
	InsertPhoneCode("Kazakhstan", "+7")
	InsertPhoneCode("Kenya", "+254")
	InsertPhoneCode("Kiribati", "+686")
	InsertPhoneCode("Kuwait", "+965")
	InsertPhoneCode("Kyrgyzstan", "+996")
	InsertPhoneCode("Laos", "+856")
	InsertPhoneCode("Latvia", "+371")
	InsertPhoneCode("Lebanon", "+961")
	InsertPhoneCode("Lesotho", "+266")
	InsertPhoneCode("Liberia", "+231")
	InsertPhoneCode("Liechtenstein", "+423")
	InsertPhoneCode("Lithuania", "+370")
	InsertPhoneCode("Luxembourg", "+352")
	InsertPhoneCode("North Macedonia", "+389")
	InsertPhoneCode("Madagascar", "+261")
	InsertPhoneCode("Malawi", "+265")
	InsertPhoneCode("Malaysia", "+60")
	InsertPhoneCode("Maldives", "+960")
	InsertPhoneCode("Mali", "+223")
	InsertPhoneCode("Malta", "+356")
	InsertPhoneCode("Marshall Islands", "+692")
	InsertPhoneCode("Martinique", "+596")
	InsertPhoneCode("Mauritania", "+222")
	InsertPhoneCode("Mauritius", "+230")
	InsertPhoneCode("Mayotte", "+262")
	InsertPhoneCode("Mexico", "+52")
	InsertPhoneCode("Micronesia, Federated States of", "+691")
	InsertPhoneCode("Moldova", "+373")
	InsertPhoneCode("Monaco", "+377")
	InsertPhoneCode("Mongolia", "+976")
	InsertPhoneCode("Montserrat", "+1664")
	InsertPhoneCode("Morocco", "+212")
	InsertPhoneCode("Mozambique", "+258")
	InsertPhoneCode("Myanmar", "+95")
	InsertPhoneCode("Namibia", "+264")
	InsertPhoneCode("Nauru", "+674")
	InsertPhoneCode("Nepal", "+977")
	InsertPhoneCode("Netherlands", "+31")
	InsertPhoneCode("New Caledonia", "+687")
	InsertPhoneCode("New Zealand", "+64")
	InsertPhoneCode("Nicaragua", "+505")
	InsertPhoneCode("Niger", "+227")
	InsertPhoneCode("Nigeria", "+234")
	InsertPhoneCode("Niue", "+683")
	InsertPhoneCode("Norfolk Island", "+672")
	InsertPhoneCode("Northern Mariana Islands", "+1670")
	InsertPhoneCode("Norway", "+47")
	InsertPhoneCode("Oman", "+968")
	InsertPhoneCode("Pakistan", "+92")
	InsertPhoneCode("Palau", "+680")
	InsertPhoneCode("Panama", "+507")
	InsertPhoneCode("Papua New Guinea", "+675")
	InsertPhoneCode("Paraguay", "+595")
	InsertPhoneCode("Peru", "+51")
	InsertPhoneCode("Philippines", "+63")
	InsertPhoneCode("Poland", "+48")
	InsertPhoneCode("Portugal", "+351")
	InsertPhoneCode("Puerto Rico", "+1939")
	InsertPhoneCode("Qatar", "+974")
	InsertPhoneCode("Reunion", "+262")
	InsertPhoneCode("Romania", "+40")
	InsertPhoneCode("Russian Federation", "+7")
	InsertPhoneCode("Rwanda", "+250")
	InsertPhoneCode("Saint Helena", "+290")
	InsertPhoneCode("Saint Kitts and Nevis", "+1869")
	InsertPhoneCode("Saint Lucia", "+1758")
	InsertPhoneCode("Saint Pierre and Miquelon", "+508")
	InsertPhoneCode("Saint Vincent and the Grenadines", "+1784")
	InsertPhoneCode("Samoa", "+685")
	InsertPhoneCode("San Marino", "+378")
	InsertPhoneCode("Sao Tome and Principe", "+239")
	InsertPhoneCode("Saudi Arabia", "+966")
	InsertPhoneCode("Senegal", "+221")
	InsertPhoneCode("Seychelles", "+248")
	InsertPhoneCode("Sierra Leone", "+232")
	InsertPhoneCode("Singapore", "+65")
	InsertPhoneCode("Slovakia", "+421")
	InsertPhoneCode("Slovenia", "+386")
	InsertPhoneCode("Solomon Islands", "+677")
	InsertPhoneCode("Somalia", "+252")
	InsertPhoneCode("South Africa", "+27")
	InsertPhoneCode("South Georgia and the South Sandwich Islands", "+500")
	InsertPhoneCode("South Sudan", "+211")
	InsertPhoneCode("Spain", "+34")
	InsertPhoneCode("Sudan", "+249")
	InsertPhoneCode("Suriname", "+597")
	InsertPhoneCode("Swaziland", "+268")
	InsertPhoneCode("Sweden", "+46")
	InsertPhoneCode("Switzerland", "+41")
	InsertPhoneCode("Syria", "+963")
	InsertPhoneCode("Tajikistan", "+992")
	InsertPhoneCode("Tanzania", "+255")
	InsertPhoneCode("Thailand", "+66")
	InsertPhoneCode("Togo", "+228")
	InsertPhoneCode("Tokelau", "+690")
	InsertPhoneCode("Tonga", "+676")
	InsertPhoneCode("Trinidad and Tobago", "+1868")
	InsertPhoneCode("Tunisia", "+216")
	InsertPhoneCode("Turkey", "+90")
	InsertPhoneCode("Turkmenistan", "+993")
	InsertPhoneCode("Turks and Caicos Islands", "+1649")
	InsertPhoneCode("Tuvalu", "+688")
	InsertPhoneCode("Uganda", "+256")
	InsertPhoneCode("Ukraine", "+380")
	InsertPhoneCode("United Arab Emirates", "+971")
	InsertPhoneCode("United Kingdom", "+44")
	InsertPhoneCode("United States", "+1")
	InsertPhoneCode("Uruguay", "+598")
	InsertPhoneCode("Uzbekistan", "+998")
	InsertPhoneCode("Vanuatu", "+678")
	InsertPhoneCode("Venezuela", "+58")
	InsertPhoneCode("Vietnam", "+84")
	InsertPhoneCode("Wallis and Futuna", "+681")
	InsertPhoneCode("Yemen", "+967")
	InsertPhoneCode("Zambia", "+260")
	InsertPhoneCode("Zimbabwe", "+263")
}
