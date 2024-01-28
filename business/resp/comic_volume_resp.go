package resp

type ComicVolumeResp struct {
	ComicVolumeId int64              `json:"comicVolumeId"`
	VolumeName    string             `json:"volumeName"`
	Flag          int8               `json:"flag"`
	SeqNo         int64              `json:"seqNo"`
	Chapters      []ComicChapterResp `json:"chapters"`
}
