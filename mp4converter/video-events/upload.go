package videoevents

type UploadVideoEvent struct {
	VideoFileId string `json:"video_file_id"`
	AudioFileId string `json:"audio_file_id"`
	UserId      int64  `json:"user_id"`
	Email       string `json:"email"`
}
