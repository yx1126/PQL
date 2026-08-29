package live

import (
	"bytes"
	"crypto/md5"
	cryptorand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"pql/pkg/request"
	"pql/pkg/utils/tool"
	"regexp"
	"strconv"
	"strings"
	"time"

	"resty.dev/v3"
)

type HuyaResponse struct {
	Docs     []HuyaSearchInfo `json:"docs"`
	NumFound int              `json:"numFound"`
	Start    int              `json:"start"`
}

type Huya struct {
	userAgent string
	http      *request.Http
}

type HuyaRoomInfo struct {
	ExceptionType *int   `json:"exceptionType"`
	Description   string `json:"welcomeText"`
	RoomInfo      struct {
		LiveStatus int `json:"eLiveStatus"`
		LiveInfo   struct {
			RoomId      int    `json:"lProfileRoom"`
			RoomName    string `json:"sIntroduction"`
			OwnerName   string `json:"sNick"`
			OwnerAvatar string `json:"sAvatar180"`
			RoomPic     string `json:"sScreenshot"`

			StreamInfo struct {
				Streams struct {
					Value []struct {
						FlvURL     string `json:"sFlvUrl"`
						StreamName string `json:"sStreamName"`
						LineIndex  int    `json:"iLineIndex"`
					} `json:"value"`
				} `json:"vStreamInfo"`
				Rates struct {
					Value []struct {
						RateName string `json:"sDisplayName"`
						Rate     int    `json:"iBitRate"`
					} `json:"value"`
				} `json:"vBitRateInfo"`
			} `json:"tLiveStreamInfo"`
		} `json:"tLiveInfo"`

		ReplayInfo struct {
			RoomId      int    `json:"lProfileRoom"`
			RoomName    string `json:"sIntroduction"`
			OwnerName   string `json:"sNick"`
			OwnerAvatar string `json:"sAvatar180"`
			RoomPic     string `json:"sScreenshot"`
		} `json:"tReplayInfo"`
	} `json:"roomInfo"`
}

type HuyaSearchInfo struct {
	Rid      int    `json:"room_id"`
	NickName string `json:"game_nick"`
	Avatar   string `json:"game_imgUrl"`
	CateName string `json:"gameName"`
	RoomName string `json:"game_introduction"`
	RoomSrc  string `json:"game_screenshot"`
}

const (
	WUPURL  = "http://wup.huya.com"
	HYSDKUA = "HYSDK(Windows, 30000002)_APP(pc_exe&7060000&official)_SDK(trans&2.32.3.5646)"
)

var (
	globalInitRE = regexp.MustCompile(`(?s)window\.HNF_GLOBAL_INIT\s*=\s*(\{.*?\})\s*</script>`)
	functionRE   = regexp.MustCompile(`(?s)function.*?\(.*?\).\{.*?\}`)
	topSIDRE     = regexp.MustCompile(`lChannelId"\s*:\s*([0-9]+)`)
	subSIDRE     = regexp.MustCompile(`lSubChannelId"\s*:\s*([0-9]+)`)
)

func NewHuya(http *request.Http) *Huya {
	return &Huya{
		userAgent: "Mozilla/5.0 (Linux; Android 11; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.91 Mobile Safari/537.36 Edg/117.0.0.0",
		http:      http,
	}
}

func (d *Huya) GetHeaders(roomId int) map[string]string {
	return map[string]string{
		"accept":     "application/json, text/plain, */*",
		"referer":    fmt.Sprintf("https://www.douyu.com/%v", roomId),
		"user-agent": d.userAgent,
	}
}

func (d *Huya) GetRoomWebUri(roomId string) string {
	return "https://www.huya.com/" + roomId
}

func (d *Huya) GetInfo(roomId string) (*RoomInfo, error) {
	room, _, err := d.getBaseInfo(roomId)
	if err != nil {
		return nil, err
	}
	live := room.RoomInfo.LiveInfo
	replay := room.RoomInfo.ReplayInfo
	isReplay := room.RoomInfo.LiveStatus == 3
	return &RoomInfo{
		RoomId:    strconv.Itoa(tool.Flag(isReplay, replay.RoomId, live.RoomId)),
		RoomName:  tool.Flag(isReplay, replay.RoomName, live.RoomName),
		OwnerName: tool.Flag(isReplay, replay.OwnerName, live.OwnerName),
		// IsLive:      room.RoomInfo.LiveStatus == 2 || isReplay,
		IsLive:      room.RoomInfo.LiveStatus == 2,
		IsReplay:    false,
		OwnerAvatar: tool.Flag(isReplay, replay.OwnerAvatar, live.OwnerAvatar),
		RoomPic:     tool.Flag(isReplay, replay.RoomPic, live.RoomPic),
		Description: room.Description,
	}, nil
}

func (d *Huya) GetPlayUrl(roomId string, extra map[string]string) (*LiveInfo, error) {
	cdn := extra["cdn"]
	rate := extra["rate"]
	room, resp, err := d.getBaseInfo(roomId)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}

	presenterUID := regexpInt64(topSIDRE, resp.Bytes())
	if presenterUID <= 0 {
		presenterUID = regexpInt64(subSIDRE, resp.Bytes())
	}
	if presenterUID <= 0 {
		var raw any
		decoder := json.NewDecoder(bytes.NewReader(jsonBytes))
		decoder.UseNumber()
		if decoder.Decode(&raw) == nil {
			presenterUID = findPositiveInt(raw, map[string]bool{"lchannelid": true, "channelid": true}, 0)
			if presenterUID <= 0 {
				presenterUID = findPositiveInt(raw, map[string]bool{"lsubchannelid": true, "subchannelid": true}, 0)
			}
		}
	}
	if presenterUID <= 0 {
		return nil, errors.New("presenter UID was not found in Huya room data")
	}

	cnds := make([]CdnsWithName, 0)
	for _, v := range room.RoomInfo.LiveInfo.StreamInfo.Streams.Value {
		cnds = append(cnds, CdnsWithName{
			Name:       fmt.Sprintf("线路%v", v.LineIndex),
			Cdn:        strings.TrimRight(v.FlvURL, "/") + "/" + v.StreamName,
			StreamName: v.StreamName,
		})
	}

	multirates := make([]Multirate, 0)
	for _, v := range room.RoomInfo.LiveInfo.StreamInfo.Rates.Value {
		multirates = append(multirates, Multirate{
			Name: v.RateName,
			Rate: strconv.Itoa(v.Rate),
		})
	}

	if len(cnds) <= 0 || len(multirates) <= 0 {
		return nil, errors.New(ParserErr)
	}

	if strings.TrimSpace(cdn) == "" {
		cdn = cnds[0].Cdn
	}

	if strings.TrimSpace(rate) == "-1" {
		rate = multirates[0].Rate
	}

	cdnItem := tool.Find(cnds, func(item CdnsWithName, _ int) bool {
		return item.Cdn == cdn
	})

	rateItem := tool.Find(multirates, func(item Multirate, _ int) bool {
		return item.Rate == rate
	})

	if cdnItem == nil || rateItem == nil {
		return nil, errors.New(ParserErr)
	}

	antiCode, err := d.getCdnTokenInfoEx(cdnItem.StreamName, presenterUID)

	if err != nil {
		return nil, err
	}

	playURL := cdnItem.Cdn + ".flv?" + antiCode + "&codec=264"
	rateInt, err := strconv.Atoi(rate)
	if err != nil {
		return nil, err
	}
	if rateInt > 0 {
		playURL += "&ratio=" + rate
	}

	return &LiveInfo{
		RoomId:       strconv.Itoa(room.RoomInfo.LiveInfo.RoomId),
		VideoType:    "flv",
		RtmpCdn:      cdn,
		Rate:         rate,
		Url:          playURL,
		CdnsWithName: cnds,
		Multirates:   multirates,
	}, nil
}

func (d *Huya) Search(params LiveParams) (*HuyaResponse, error) {
	page, err := strconv.Atoi(params.Page)
	if err != nil {
		return nil, err
	}

	r := d.http.R()

	start := strconv.Itoa((page - 1) * 20)

	r.SetQueryParams(map[string]string{
		"m":         "Search",
		"do":        "getSearchContent",
		"q":         params.Keyword,
		"uid":       "0",
		"v":         "4",
		"typ":       "-5",
		"livestate": "0",
		"rows":      "20",
		"start":     start,
	})

	resp, err := r.Execute(http.MethodGet, "https://search.cdn.huya.com")
	if err != nil {
		return nil, err
	}

	var response struct {
		Response struct {
			Value HuyaResponse `json:"3"`
		} `json:"response"`
	}
	if err := json.Unmarshal(resp.Bytes(), &response); err != nil {
		return nil, err
	}
	return &response.Response.Value, nil
}

func (d *Huya) getBaseInfo(roomId string) (*HuyaRoomInfo, *resty.Response, error) {
	r := d.http.R()
	r.SetHeader("User-Agent", d.userAgent)
	resp, err := r.Execute(http.MethodGet, "https://m.huya.com/"+roomId)
	if err != nil {
		return nil, nil, err
	}
	match := globalInitRE.FindSubmatch(resp.Bytes())
	if len(match) != 2 {
		return nil, nil, errors.New(ParserErr)
	}
	jsonData := functionRE.ReplaceAll(match[1], []byte(`""`))
	var room HuyaRoomInfo
	if err := json.Unmarshal(jsonData, &room); err != nil {
		return nil, nil, errors.New(RoomErr)
	}
	if room.ExceptionType != nil && *room.ExceptionType == 0 {
		return nil, nil, errors.New(RoomErr)
	}
	return &room, resp, nil
}

func (d *Huya) getCdnTokenInfoEx(streamName string, presenterUID int64) (string, error) {
	body := buildWUPRequest(streamName)

	r := d.http.R()

	r.SetHeaders(map[string]string{
		"Content-Type": "application/x-wup",
		"Origin":       "https://m.huya.com/",
		"Referer":      "https://m.huya.com/",
		"User-Agent":   HYSDKUA,
	})

	r.SetBody(body)

	resp, err := r.Execute(http.MethodPost, WUPURL)
	if err != nil {
		return "", fmt.Errorf("create WUP request: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return "", fmt.Errorf("read Huya WUP response: %w", err)
	}
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return "", fmt.Errorf("Huya WUP returned %s: %s", resp.Status(), strings.TrimSpace(string(data)))
	}
	code, err := decodeWUPResponse(data)
	if err != nil {
		return "", err
	}

	antiCode, err := buildAntiCode(streamName, presenterUID, code)
	if err != nil {
		return "", err
	}
	return antiCode, nil
}

func buildAntiCode(stream string, presenterUID int64, antiCode string) (string, error) {
	p, err := url.ParseQuery(antiCode)
	if err != nil {
		return "", fmt.Errorf("parse anti code: %w", err)
	}
	if !p.Has("fm") {
		return antiCode, nil
	}
	ctype := p.Get("ctype")
	if ctype == "" {
		ctype = "huya_pc_exe"
	}
	platformID, err := strconv.ParseInt(defaultString(p.Get("t"), "0"), 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid t: %w", err)
	}
	wsTime := p.Get("wsTime")
	if wsTime == "" {
		return "", errors.New("anti code is missing wsTime")
	}

	seqID := presenterUID + time.Now().UnixMilli()
	secretHash := md5Hex(fmt.Sprintf("%d|%s|%d", seqID, ctype, platformID))
	convertedUID := rotateLow32Left8(presenterUID)
	calcUID := convertedUID
	isWAP := platformID == 103
	if isWAP {
		calcUID = presenterUID
	}

	fm := p.Get("fm")                                  // ParseQuery 已完成 Dart queryParametersAll 的解码。
	if decoded, e := url.QueryUnescape(fm); e == nil { // 对应额外的 decodeComponent。
		fm = decoded
	}
	fmBytes, err := base64.StdEncoding.DecodeString(fm)
	if err != nil {
		return "", fmt.Errorf("decode fm: %w", err)
	}
	secretPrefix := strings.SplitN(string(fmBytes), "_", 2)[0]
	wsSecret := md5Hex(fmt.Sprintf("%s_%d_%s_%s_%s", secretPrefix, calcUID, stream, secretHash, wsTime))

	// 手工拼接以保持 Dart 中字段顺序；QueryEscape 对这些字段与服务端兼容。
	fields := [][2]string{
		{"wsSecret", wsSecret}, {"wsTime", wsTime},
		{"seqid", strconv.FormatInt(seqID, 10)}, {"ctype", ctype},
		{"ver", "1"}, {"fs", p.Get("fs")}, {"fm", fm},
		{"t", strconv.FormatInt(platformID, 10)},
	}
	if isWAP {
		uuid, e := makeUUID(wsTime)
		if e != nil {
			return "", e
		}
		fields = append(fields, [2]string{"uid", strconv.FormatInt(presenterUID, 10)}, [2]string{"uuid", uuid})
	} else {
		fields = append(fields, [2]string{"u", strconv.FormatInt(convertedUID, 10)})
	}
	var out strings.Builder
	for i, field := range fields {
		if i != 0 {
			out.WriteByte('&')
		}
		out.WriteString(field[0])
		out.WriteByte('=')
		out.WriteString(url.QueryEscape(field[1]))
	}
	return out.String(), nil
}

func decodeWUPResponse(data []byte) (string, error) {
	if len(data) < 4 || int(binary.BigEndian.Uint32(data[:4])) > len(data) {
		return "", errors.New("invalid WUP response length")
	}
	packet := tarsReader{b: data[4:]}
	h, err := packet.field(7)
	if err != nil {
		return "", fmt.Errorf("find response buffer: %w", err)
	}
	attrs, err := packet.blob(h)
	if err != nil {
		return "", err
	}
	// newData map<string, byte[]>
	r := tarsReader{b: attrs}
	mh, err := r.field(0)
	if err != nil || mh.typ != tMap {
		return "", errors.New("invalid TUP3 response map")
	}
	lh, err := r.head()
	if err != nil {
		return "", err
	}
	count, err := r.integer(lh)
	if err != nil {
		return "", err
	}
	var code int64
	var rsp []byte
	for i := int64(0); i < count; i++ {
		kh, e := r.head()
		if e != nil {
			return "", e
		}
		k, e := r.string(kh)
		if e != nil {
			return "", e
		}
		vh, e := r.head()
		if e != nil {
			return "", e
		}
		v, e := r.blob(vh)
		if e != nil {
			return "", e
		}
		switch k {
		case "":
			cr := tarsReader{b: v}
			ch, e := cr.field(0)
			if e == nil {
				code, _ = cr.integer(ch)
			}
		case "tRsp":
			rsp = v
		}
	}
	if code != 0 {
		return "", fmt.Errorf("Huya WUP result code %d", code)
	}
	if len(rsp) == 0 {
		return "", errors.New("Huya WUP response has no tRsp")
	}
	// tRsp 是 tag 0 struct，内部 tag 0 为 sFlvToken。
	rr := tarsReader{b: rsp}
	sh, err := rr.field(0)
	if err != nil || sh.typ != tStructBegin {
		return "", errors.New("invalid tRsp struct")
	}
	th, err := rr.field(0)
	if err != nil {
		return "", fmt.Errorf("read sFlvToken: %w", err)
	}
	token, err := rr.string(th)
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("Huya returned an empty FLV token")
	}
	return token, nil
}

func rotateLow32Left8(v int64) int64 {
	u := uint64(v)
	low := uint32(u)
	return int64((u & 0xffffffff00000000) | uint64(low<<8|low>>24))
}

func md5Hex(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func defaultString(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

func randomFraction() (float64, error) {
	var b [8]byte
	if _, err := cryptorand.Read(b[:]); err != nil {
		return 0, err
	}
	return float64(binary.BigEndian.Uint64(b[:])>>11) / float64(uint64(1)<<53), nil
}

func makeUUID(wsTime string) (string, error) {
	v, err := strconv.ParseInt(wsTime, 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid wsTime: %w", err)
	}
	r1, err := randomFraction()
	if err != nil {
		return "", err
	}
	r2, err := randomFraction()
	if err != nil {
		return "", err
	}
	ct := int64((float64(v) + r1) * 1000)
	uuid := int64(math.Mod((float64(ct%10_000_000_000)+r2)*1000, float64(uint64(0xffffffff))))
	return strconv.FormatInt(uuid, 10), nil
}

// 以下是本接口所需的最小 TARS/TUP3 codec，不依赖第三方包。
const (
	tByte = iota
	tShort
	tInt
	tLong
	tFloat
	tDouble
	tString1
	tString4
	tMap
	tList
	tStructBegin
	tStructEnd
	tZero
	tSimpleList
)

type tarsWriter struct{ bytes.Buffer }

func (w *tarsWriter) head(typ, tag int) {
	if tag < 15 {
		w.WriteByte(byte(tag<<4 | typ))
		return
	}
	w.WriteByte(byte(0xf0 | typ))
	w.WriteByte(byte(tag))
}
func (w *tarsWriter) integer(v int64, tag int) {
	switch {
	case v == 0:
		w.head(tZero, tag)
	case v >= -128 && v <= 127:
		w.head(tByte, tag)
		w.WriteByte(byte(int8(v)))
	case v >= -32768 && v <= 32767:
		w.head(tShort, tag)
		_ = binary.Write(&w.Buffer, binary.BigEndian, int16(v))
	case v >= -2147483648 && v <= 2147483647:
		w.head(tInt, tag)
		_ = binary.Write(&w.Buffer, binary.BigEndian, int32(v))
	default:
		w.head(tLong, tag)
		_ = binary.Write(&w.Buffer, binary.BigEndian, v)
	}
}
func (w *tarsWriter) str(s string, tag int) {
	b := []byte(s)
	if len(b) <= 255 {
		w.head(tString1, tag)
		w.WriteByte(byte(len(b)))
	} else {
		w.head(tString4, tag)
		_ = binary.Write(&w.Buffer, binary.BigEndian, uint32(len(b)))
	}
	w.Write(b)
}
func (w *tarsWriter) blob(b []byte, tag int) {
	w.head(tSimpleList, tag)
	w.head(tByte, 0)
	w.integer(int64(len(b)), 0)
	w.Write(b)
}
func (w *tarsWriter) stringMap(m map[string]string, tag int) {
	w.head(tMap, tag)
	w.integer(int64(len(m)), 0)
	for k, v := range m {
		w.str(k, 0)
		w.str(v, 1)
	}
}
func (w *tarsWriter) bytesMap(key string, value []byte, tag int) {
	w.head(tMap, tag)
	w.integer(1, 0)
	w.str(key, 0)
	w.blob(value, 1)
}

func buildWUPRequest(stream string) []byte {
	// HuyaUserId
	var uid tarsWriter
	uid.integer(0, 0)
	uid.str("", 1)
	uid.str("", 2)
	uid.str("pc_exe&7060000&official", 3)
	uid.str("", 4)
	uid.integer(0, 5)
	uid.str("", 6)
	uid.str("", 7)
	// GetCdnTokenExReq
	var req tarsWriter
	req.str("", 0)
	req.str(stream, 1)
	req.integer(0, 2)
	req.head(tStructBegin, 3)
	req.Write(uid.Bytes())
	req.head(tStructEnd, 0)
	req.integer(66, 4)
	// TUP3 newData = {"tReq": encoded struct}
	var attrs tarsWriter
	var wrapped tarsWriter
	wrapped.head(tStructBegin, 0)
	wrapped.Write(req.Bytes())
	wrapped.head(tStructEnd, 0)
	attrs.bytesMap("tReq", wrapped.Bytes(), 0)
	// RequestPacket
	var packet tarsWriter
	packet.integer(3, 1)
	packet.integer(0, 2)
	packet.integer(0, 3)
	packet.integer(0, 4)
	packet.str("liveui", 5)
	packet.str("getCdnTokenInfoEx", 6)
	packet.blob(attrs.Bytes(), 7)
	packet.integer(0, 8)
	packet.stringMap(map[string]string{}, 9)
	packet.stringMap(map[string]string{}, 10)
	body := packet.Bytes()
	out := make([]byte, 4+len(body))
	binary.BigEndian.PutUint32(out, uint32(len(out)))
	copy(out[4:], body)
	return out
}

type tarsReader struct {
	b []byte
	p int
}

type tarsHead struct{ typ, tag int }

func (r *tarsReader) head() (tarsHead, error) {
	if r.p >= len(r.b) {
		return tarsHead{}, io.ErrUnexpectedEOF
	}
	v := r.b[r.p]
	r.p++
	h := tarsHead{typ: int(v & 0x0f), tag: int(v >> 4)}
	if h.tag == 15 {
		if r.p >= len(r.b) {
			return h, io.ErrUnexpectedEOF
		}
		h.tag = int(r.b[r.p])
		r.p++
	}
	return h, nil
}
func (r *tarsReader) integer(h tarsHead) (int64, error) {
	n := map[int]int{tByte: 1, tShort: 2, tInt: 4, tLong: 8}[h.typ]
	if h.typ == tZero {
		return 0, nil
	}
	if n == 0 || r.p+n > len(r.b) {
		return 0, errors.New("invalid TARS integer")
	}
	b := r.b[r.p : r.p+n]
	r.p += n
	switch n {
	case 1:
		return int64(int8(b[0])), nil
	case 2:
		return int64(int16(binary.BigEndian.Uint16(b))), nil
	case 4:
		return int64(int32(binary.BigEndian.Uint32(b))), nil
	default:
		return int64(binary.BigEndian.Uint64(b)), nil
	}
}
func (r *tarsReader) string(h tarsHead) (string, error) {
	var n int
	if h.typ == tString1 {
		if r.p >= len(r.b) {
			return "", io.ErrUnexpectedEOF
		}
		n = int(r.b[r.p])
		r.p++
	} else if h.typ == tString4 {
		if r.p+4 > len(r.b) {
			return "", io.ErrUnexpectedEOF
		}
		n = int(binary.BigEndian.Uint32(r.b[r.p:]))
		r.p += 4
	} else {
		return "", errors.New("invalid TARS string")
	}
	if n < 0 || r.p+n > len(r.b) {
		return "", io.ErrUnexpectedEOF
	}
	s := string(r.b[r.p : r.p+n])
	r.p += n
	return s, nil
}
func (r *tarsReader) blob(h tarsHead) ([]byte, error) {
	if h.typ != tSimpleList {
		return nil, errors.New("invalid TARS byte array")
	}
	elem, err := r.head()
	if err != nil || elem.typ != tByte {
		return nil, errors.New("invalid TARS byte array element")
	}
	lh, err := r.head()
	if err != nil {
		return nil, err
	}
	n, err := r.integer(lh)
	if err != nil || n < 0 || n > int64(len(r.b)-r.p) {
		return nil, errors.New("invalid TARS byte array length")
	}
	v := append([]byte(nil), r.b[r.p:r.p+int(n)]...)
	r.p += int(n)
	return v, nil
}
func (r *tarsReader) skip(h tarsHead) error {
	switch h.typ {
	case tZero, tStructEnd:
		return nil
	case tByte, tShort, tInt, tLong:
		_, e := r.integer(h)
		return e
	case tFloat:
		r.p += 4
	case tDouble:
		r.p += 8
	case tString1, tString4:
		_, e := r.string(h)
		return e
	case tSimpleList:
		_, e := r.blob(h)
		return e
	case tStructBegin:
		for {
			x, e := r.head()
			if e != nil {
				return e
			}
			if x.typ == tStructEnd {
				return nil
			}
			if e = r.skip(x); e != nil {
				return e
			}
		}
	case tList:
		lh, e := r.head()
		if e != nil {
			return e
		}
		n, e := r.integer(lh)
		if e != nil {
			return e
		}
		for i := int64(0); i < n; i++ {
			x, e := r.head()
			if e != nil {
				return e
			}
			if e = r.skip(x); e != nil {
				return e
			}
		}
	case tMap:
		lh, e := r.head()
		if e != nil {
			return e
		}
		n, e := r.integer(lh)
		if e != nil {
			return e
		}
		for i := int64(0); i < n*2; i++ {
			x, e := r.head()
			if e != nil {
				return e
			}
			if e = r.skip(x); e != nil {
				return e
			}
		}
	default:
		return fmt.Errorf("unsupported TARS type %d", h.typ)
	}
	if r.p > len(r.b) {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (r *tarsReader) field(tag int) (tarsHead, error) {
	for r.p < len(r.b) {
		h, e := r.head()
		if e != nil {
			return h, e
		}
		if h.tag == tag {
			return h, nil
		}
		if h.tag > tag {
			return h, fmt.Errorf("TARS field %d not found", tag)
		}
		if e = r.skip(h); e != nil {
			return h, e
		}
	}
	return tarsHead{}, io.EOF
}

func regexpInt64(re *regexp.Regexp, data []byte) int64 {
	m := re.FindSubmatch(data)
	if len(m) != 2 {
		return 0
	}
	v, _ := strconv.ParseInt(string(m[1]), 10, 64)
	return v
}

func findPositiveInt(value any, keys map[string]bool, depth int) int64 {
	if depth > 8 {
		return 0
	}
	switch v := value.(type) {
	case map[string]any:
		for key, item := range v {
			if keys[strings.ToLower(key)] {
				if n := anyPositiveInt(item); n > 0 {
					return n
				}
			}
		}
		for _, item := range v {
			if n := findPositiveInt(item, keys, depth+1); n > 0 {
				return n
			}
		}
	case []any:
		for _, item := range v {
			if n := findPositiveInt(item, keys, depth+1); n > 0 {
				return n
			}
		}
	}
	return 0
}

func anyPositiveInt(value any) int64 {
	var text string
	switch v := value.(type) {
	case json.Number:
		text = v.String()
	case string:
		text = strings.TrimSpace(v)
	default:
		return 0
	}
	n, _ := strconv.ParseInt(text, 10, 64)
	if n > 0 {
		return n
	}
	return 0
}
