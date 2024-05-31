package my_enum_example

type HttpStatusCode int

const (
	OK HttpStatusCode = 200 
	Created HttpStatusCode = 201
	Accepted HttpStatusCode = 202
	NonAuthoritativeInformation HttpStatusCode = 203
)

var httpStatusMap = map[HttpStatusCode]string {
		OK: "OK",
		Created: "Created",
		Accepted: "Accepted",
		NonAuthoritativeInformation: "Non-Authoritative Information",
}

func (hsc HttpStatusCode) String() string {
	return httpStatusMap[hsc]
}