package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"slices"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
)

type VidMeta struct {
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	Duration float64 `json:"duration"`
}

func GetVidMetaFromBytes(vidBytes []byte, ffmpegFormat string) (vidMate VidMeta, err error) {
	reader := BytesReaderPoolGet(vidBytes)
	defer BytesReaderPoolPut(reader)
	cmd := exec.Command(
		`ffprobe`,
		`-v`, `quiet`,
		`-print_format`, `json`,
		`-show_streams`,
		`-show_format`,
		`-f`, ffmpegFormat, // 必须指定！
		`-`, // 从 stdin 读取
	)
	cmd.Stdin = reader
	output, err := cmd.Output()
	if err != nil {
		return
	}
	outputJson := gjson.New(output)
	vidMate.Duration = outputJson.Get(`format.duration`).Float64()
	for _, v := range outputJson.Get(`streams`).Maps() {
		if gconv.String(v[`codec_type`]) == `video` {
			vidMate.Width = gconv.Int(v[`width`])
			vidMate.Height = gconv.Int(v[`height`])
			if vidMate.Width > 0 && vidMate.Height > 0 {
				break
			}
		}
		/* switch gconv.String(v[`codec_type`]) {
		case `video`:
			vidMate.Width = gconv.Int(v[`width`])
			vidMate.Height = gconv.Int(v[`height`])
		case `audio`:
		} */
	}
	if vidMate.Width == 0 || vidMate.Height == 0 || vidMate.Duration == 0 {
		err = errors.New(`获取视频元信息失败`)
	}
	return
}

var mimeToFFmpeg = map[string]string{
	`video/mp4`:         `mp4`,
	`video/webm`:        `webm`,
	`video/quicktime`:   `mov`,
	`video/x-quicktime`: `mov`,
	`video/avi`:         `avi`,
	`video/x-msvideo`:   `avi`,
	`video/mkv`:         `matroska`,
	`video/x-matroska`:  `matroska`,
	"video/x-ms-wmv":    "asf",
	"video/x-ms-asf":    "asf",
	"video/x-flv":       "flv",
	"video/mpeg":        "mpeg",
	"video/mp2t":        "mpegts",
	"video/3gpp":        "3gp",
	"video/3gpp2":       "3g2",
	"video/ogg":         "ogg",
	"video/ogv":         "ogg",
}

type VidOption struct {
	Width           int      `json:"width"`
	Height          int      `json:"height"`
	MinDuration     float64  `json:"min_duration"` //注意：转换目标格式是webm时，不支持填充时长
	MaxDuration     float64  `json:"max_duration"`
	EncodeFormatArr []string `json:"encode_format_arr"` //需要转换的格式：video/mp4, video/web
	TargerFormat    string   `json:"targer_format"`     //当EncodeFormatArr不为空，且需要格式转换时才有用，用于指定转换后的目标格式，默认：mp4
}

func VidHandle(vidBytesOfRaw []byte, vidOption VidOption) (vidBytes []byte, err error) {
	vidBytes = vidBytesOfRaw
	vidType := http.DetectContentType(vidBytes[:min(512, len(vidBytes))])
	vidType = strings.TrimSpace(strings.Split(vidType, `;`)[0])
	ffmpegFormat, ok := mimeToFFmpeg[vidType]
	if !ok {
		err = errors.New(`未识别的视频类型`)
		return
	}
	vidMeta, err := GetVidMetaFromBytes(vidBytes, ffmpegFormat)
	if err != nil {
		return
	}
	isHandle := false
	var argsOfVf []string
	var argsOfAf []string
	var argsOfT []string
	targerFormat := ffmpegFormat
	if vidOption.Width > 0 && vidOption.Height > 0 && !(vidMeta.Width == vidOption.Width && vidMeta.Height == vidOption.Height) {
		isHandle = true
		argsOfVf = append(argsOfVf,
			fmt.Sprintf(`scale=%d:%d:force_original_aspect_ratio=decrease`, vidOption.Width, vidOption.Height),
			fmt.Sprintf(`pad=%d:%d:(ow-iw)/2:(oh-ih)/2`, vidOption.Width, vidOption.Height),
		)
	}
	if vidOption.MinDuration > 0 && vidMeta.Duration < vidOption.MinDuration {
		if vidOption.MaxDuration > 0 && vidOption.MinDuration > vidOption.MaxDuration {
			err = errors.New(`最小时长不能大于最大时长`)
			return
		}
		isHandle = true
		padDur := vidOption.MinDuration - vidMeta.Duration
		argsOfVf = append(argsOfVf, fmt.Sprintf(`tpad=stop_mode=clone:stop_duration=%.3f`, padDur))
		argsOfAf = append(argsOfAf, fmt.Sprintf(`apad=pad_dur=%.3f`, padDur))
	}
	if vidOption.MaxDuration > 0 && vidMeta.Duration > vidOption.MaxDuration {
		isHandle = true
		argsOfT = append(argsOfT, strconv.FormatFloat(vidOption.MaxDuration, 'f', -1, 64))
	}
	if len(vidOption.EncodeFormatArr) > 0 && slices.Contains(vidOption.EncodeFormatArr, vidType) {
		isHandle = true
		targerFormat = vidOption.TargerFormat
		if targerFormat == `` {
			targerFormat = `mp4`
		}
	}
	if !isHandle {
		return
	}
	args := []string{`-f`, ffmpegFormat, `-i`, `pipe:0`}
	if len(argsOfVf) > 0 {
		args = append(args, `-vf`, strings.Join(argsOfVf, `,`))
	}
	if len(argsOfAf) > 0 {
		args = append(args, `-af`, strings.Join(argsOfAf, `,`))
	}
	if len(argsOfT) > 0 {
		args = append(args, `-t`, strings.Join(argsOfT, `,`))
	}
	targerFormat = `avi`
	switch targerFormat {
	case `webm`:
		args = append(args, `-c:v`, `libvpx-vp9`, `-crf`, `30`, `-b:v`, `0`, `-c:a`, `libopus`, `-b:a`, `128k`, `-cpu-used`, `4`, `-row-mt`, `1`, `-f`, targerFormat, `pipe:1`)
	case `avi`:
		args = append(args, `-c:v`, `msmpeg4v3`, `-c:a`, `libmp3lame`, `-b:a`, `128k`, `-f`, targerFormat, `pipe:1`)
	// case `mp4`, `mov`:
	default:
		args = append(args, `-c:v`, `libx264`, `-crf`, `23`, `-preset`, `fast`, `-c:a`, `aac`, `-b:a`, `128k`, `-f`, targerFormat, `-movflags`, `frag_keyframe+empty_moov`, `pipe:1`)
	}

	cmd := exec.Command(`ffmpeg`, args...)
	reader := BytesReaderPoolGet(vidBytes)
	defer BytesReaderPoolPut(reader)
	cmd.Stdin = reader
	vidBytes, err = cmd.Output()
	return
}
