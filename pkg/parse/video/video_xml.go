package video

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var ErrVideoNotFound = errors.New("video URL not found")

type ResolveOptions struct {
	UserAgent string
	Referer   string
	Timeout   time.Duration
	Headless  bool
}

type VideoSource struct {
	URL      string
	MIMEType string
	Referer  string
}

type resolver struct {
	result chan VideoSource
	once   sync.Once
}

func (r *resolver) publish(source VideoSource) {
	if !isVideoURL(source.URL) {
		return
	}

	r.once.Do(func() {
		r.result <- source
	})
}

func ResolveVideoURL(parent context.Context, url string, opts ResolveOptions) (VideoSource, error) {
	if opts.Timeout <= 0 {
		opts.Timeout = 20 * time.Second
	}
	if opts.UserAgent == "" {
		opts.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/127.0.0.0 Safari/537.36"
	}

	if err := validatePageURL(url); err != nil {
		return VideoSource{}, err
	}

	allocOptions := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", opts.Headless),
		chromedp.Flag("autoplay-policy", "no-user-gesture-required"),
		chromedp.Flag("disable-background-networking", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent(opts.UserAgent),
	)

	allocCtx, cancelAlloc :=
		chromedp.NewExecAllocator(parent, allocOptions...)
	defer cancelAlloc()

	browserCtx, cancelBrowser := chromedp.NewContext(allocCtx)
	defer cancelBrowser()

	ctx, cancel := context.WithTimeout(browserCtx, opts.Timeout)
	defer cancel()

	r := &resolver{
		result: make(chan VideoSource, 1),
	}

	// 必须在 Navigate 之前注册，否则可能漏掉页面早期发起的请求。
	chromedp.ListenTarget(ctx, func(event any) {
		switch e := event.(type) {
		case *network.EventRequestWillBeSent:
			requestURL := e.Request.URL

			if isObviousVideoURL(requestURL) {
				r.publish(VideoSource{
					URL:     requestURL,
					Referer: opts.Referer,
				})
			}

		case *network.EventResponseReceived:
			responseURL := e.Response.URL
			mimeType := strings.ToLower(e.Response.MimeType)

			if isMediaMIME(mimeType) ||
				isObviousVideoURL(responseURL) ||
				e.Type == network.ResourceTypeMedia {

				r.publish(VideoSource{
					URL:      responseURL,
					MIMEType: mimeType,
					Referer:  opts.Referer,
				})
			}
		}
	})

	headers := network.Headers{
		"User-Agent": opts.UserAgent,
	}

	if opts.Referer != "" {
		headers["Referer"] = opts.Referer
	}

	err := chromedp.Run(
		ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(headers),
		chromedp.Navigate(url),
	)
	if err != nil {
		return VideoSource{}, fmt.Errorf("load page: %w", err)
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case source := <-r.result:
			return source, nil

		case <-ticker.C:
			source, err := inspectDOM(ctx, url)
			if err == nil && source != "" {
				r.publish(VideoSource{
					URL:     source,
					Referer: opts.Referer,
				})
			}

		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return VideoSource{}, ErrVideoNotFound
			}
			return VideoSource{}, ctx.Err()
		}
	}
}

// 扫描 video、source 和旧式 iframe。
func inspectDOM(ctx context.Context, pageURL string) (string, error) {
	const script = `
		(() => {
			const invalid = value =>
				!value ||
				value.startsWith("blob:") ||
				value.includes("googleads") ||
				value.includes("googlesyndication");

			for (const video of document.querySelectorAll("video")) {
				const values = [
					video.currentSrc,
					video.src,
					video.getAttribute("src")
				];

				for (const value of values) {
					if (!invalid(value)) {
						return new URL(value, document.baseURI).href;
					}
				}

				for (const source of video.querySelectorAll("source")) {
					const value = source.src || source.getAttribute("src");

					if (!invalid(value)) {
						return new URL(value, document.baseURI).href;
					}
				}
			}

			// 检查 iframe URL 的查询参数里是否藏着 m3u8/mp4。
			for (const iframe of document.querySelectorAll("iframe")) {
				const value = iframe.src || iframe.getAttribute("src");
				if (!value) continue;
				try {
					const iframeURL = new URL(value, document.baseURI);
					for (const parameter of iframeURL.searchParams.values()) {
						let decoded = parameter;
						try {
							decoded = decodeURIComponent(parameter);
						} catch (_) {}
						if(/^https?:\/\/.+\.(m3u8|mp4)(?:[?#].*)?$/i.test(decoded)) {
							return decoded;
						}
					}
				} catch (_) {}
			}

			return "";
		})()
		`

	var result string
	err := chromedp.Run(
		ctx,
		chromedp.Evaluate(script, &result),
	)
	if err != nil {
		return "", err
	}

	if result == "" {
		return "", nil
	}

	return resolveURL(pageURL, result), nil
}

func resolveURL(baseURL, raw string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return raw
	}

	target, err := url.Parse(raw)
	if err != nil {
		return raw
	}

	return base.ResolveReference(target).String()
}

func isMediaMIME(mimeType string) bool {
	return strings.Contains(mimeType, "mpegurl") ||
		strings.HasPrefix(mimeType, "video/") ||
		mimeType == "application/vnd.apple.mpegurl" ||
		mimeType == "application/x-mpegurl"
}

func isObviousVideoURL(raw string) bool {
	parsed, err := url.Parse(raw)
	if err != nil {
		return false
	}

	path := strings.ToLower(parsed.Path)

	return strings.HasSuffix(path, ".m3u8") ||
		strings.HasSuffix(path, ".mp4") ||
		strings.HasSuffix(path, ".flv") ||
		strings.HasSuffix(path, ".webm") ||
		strings.HasSuffix(path, ".mkv")
}

func isVideoURL(raw string) bool {
	if strings.HasPrefix(raw, "blob:") ||
		strings.HasPrefix(raw, "data:") ||
		strings.Contains(raw, "googleads") ||
		strings.Contains(raw, "googlesyndication") {
		return false
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		return false
	}

	return (parsed.Scheme == "http" || parsed.Scheme == "https") &&
		parsed.Host != ""
}

func validatePageURL(raw string) error {
	parsed, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("invalid page URL: %w", err)
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("only HTTP/HTTPS URLs are allowed")
	}

	if parsed.Host == "" {
		return errors.New("page URL has no host")
	}

	return nil
}

// func main() {
// 	url := "https://www.bmmdmm.com/play/520299-0-0.html"

// 	source, err := ResolveVideoURL(
// 		context.Background(),
// 		url,
// 		ResolveOptions{
// 			Headless:  true,
// 			Timeout:   20 * time.Second,
// 			UserAgent: "",
// 			Referer:   url,
// 		},
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("视频地址:", source.URL)
// 	fmt.Println("MIME:", source.MIMEType)
// 	fmt.Println("Referer:", source.Referer)
// }
