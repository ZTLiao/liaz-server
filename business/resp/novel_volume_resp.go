package resp

type NovelVolumeResp struct {
	NovelVolumeId int64              `json:"novelVolumeId"`
	VolumeName    string             `json:"volumeName"`
	SeqNo         int64              `json:"seqNo"`
	Chapters      []NovelChapterResp `json:"chapters"`
}
