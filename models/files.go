package models

type FileInfo struct {
	Id                           string `json:"Id"`
	Name                         string `json:"Name"`
	Size                         string `json:"Size"`
	HotlinkId                    string `json:"HotlinkId"`
	ContentType                  string `json:"ContentType"`
	ExpireAt                     int64  `json:"ExpireAt"`
	SizeBytes                    int    `json:"SizeBytes"`
	ExpireAtString               string `json:"ExpireAtString"`
	DownloadsRemaining           int    `json:"DownloadsRemaining"`
	DownloadCount                int    `json:"DownloadCount"`
	UnlimitedDownloads           bool   `json:"UnlimitedDownloads"`
	UnlimitedTime                bool   `json:"UnlimitedTime"`
	RequiresClientSideDecryption bool   `json:"RequiresClientSideDecryption"`
	IsEncrypted                  bool   `json:"IsEncrypted"`
	IsPasswordProtected          bool   `json:"IsPasswordProtected"`
	IsSavedOnLocalStorage        bool   `json:"IsSavedOnLocalStorage"`
}

type FileAddResponse struct {
	Result            string   `json:"Result"`
	FileInfo          FileInfo `json:"FileInfo"`
	HotlinkURL        string   `json:"HotlinkUrl"`
	URL               string   `json:"Url"`
	GenericHotlinkURL string   `json:"GenericHotlinkUrl"`
}
