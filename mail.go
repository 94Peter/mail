package mail

type MailLog interface {
	Info(string)
	Debug(string)
}

type MailConf interface {
	NewMailServ(l MailLog) MailServ
}

type MailServ interface {
	Subject(s string) MailServ
	PlaintText(p string) MailServ
	Html(h string) MailServ
	SendSingle(name, mail string) error
	SendMulti(tos []*To) error
}
