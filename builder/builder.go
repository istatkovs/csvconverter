package builder

type Employee struct {
	Title          string `json:"title"`
	FirstName      string `json:"firstName"`
	Sn             string `json:"sn"`
	FullName       string `json:"fullName"`
	Uid            string `json:"uid"`
	PersonalNumber string `json:"personalNumber"`
	Mail           string `json:"mail"`
	Company        string `json:"company"`
	Manager        string `json:"manager"`
	Function       string `json:"function"`
}

func NewEmployee() Employee {
	return Employee{}
}

func (e Employee) WithTitle(title string) Employee {
	e.Title = title
	return e
}

func (e Employee) WithFirstName(firstname string) Employee {
	e.FirstName = firstname
	return e
}

func (e Employee) WithSn(sn string) Employee {
	e.Sn = sn
	return e
}

func (e Employee) WithFullName(fullname string) Employee {
	e.FullName = fullname
	return e
}

func (e Employee) WithUid(uid string) Employee {
	e.Uid = uid
	return e
}

func (e Employee) WithPersonalNumber(persnr string) Employee {
	e.PersonalNumber = persnr
	return e
}

func (e Employee) WithMail(mail string) Employee {
	e.Mail = mail
	return e
}

func (e Employee) WithCompany(company string) Employee {
	e.Company = company
	return e
}

func (e Employee) WithManager(manager string) Employee {
	e.Manager = manager
	return e
}

func (e Employee) WithFunction(function string) Employee {
	e.Function = function
	return e
}
