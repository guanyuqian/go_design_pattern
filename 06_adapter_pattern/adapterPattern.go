package _6_adapter_pattern

// 适配器模式

//步骤 1
//为媒体播放器和更高级的媒体播放器创建接口。
type MediaPlayer interface {
	play(audioType, fileName string) string
}
type AdvanceMediaPlayer interface {
	PlayVlc(fileName string) string
	PlayMp4(fileName string) string
}

//步骤 2
//创建实现了 AdvancedMediaPlayer 接口的实体类。
type VlcPlayer struct{}
type Mp4Player struct{}

func (receiver VlcPlayer) PlayVlc(fileName string) string {
	return "support"
}
func (receiver Mp4Player) PlayVlc(fileName string) string {
	return "not support"
}

func (receiver VlcPlayer) PlayMp4(fileName string) string {
	return "not support"
}
func (receiver Mp4Player) PlayMp4(fileName string) string {
	return "support"
}

//步骤 3
//创建实现了 MediaPlayer 接口的适配器类。

type MediaAdapter struct {
	advancedMusicPlayer AdvanceMediaPlayer
}

func (receiver MediaAdapter) MediaAdapter(audioType string) *MediaAdapter {
	switch audioType {
	case "vlc":
		receiver.advancedMusicPlayer = new(VlcPlayer)
	case "mp4":
		receiver.advancedMusicPlayer = new(Mp4Player)
	}
	return &receiver
}

func (receiver MediaAdapter) play(audioType, fileName string) string {
	switch audioType {
	case "vlc":
		return receiver.advancedMusicPlayer.PlayVlc(fileName)
	case "mp4":
		return receiver.advancedMusicPlayer.PlayMp4(fileName)
	default:
		return "not support"
	}
}

//步骤 4
//创建实现了 MediaPlayer 接口的实体类。
type AudioPlayer struct {
	*MediaAdapter
}

func (receiver AudioPlayer) Play(audioType, fileName string) string {
	if audioType == "mp3" {
		// 播放 mp3 因为文件的内置支持
		return "support"
	} else {
		// mediaAdapter 提供了播放其他文件格式
		receiver.MediaAdapter = new(MediaAdapter).MediaAdapter(audioType)
		return receiver.MediaAdapter.play(audioType, fileName)

	}
}
