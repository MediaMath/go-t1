package t1

const (
	mediaTypeMashery = "text/xml"
)

type masheryResponse struct {
	Message string `xml:"h1"`
}
