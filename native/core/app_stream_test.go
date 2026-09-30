package core

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNativeHLSSegmentPreservesBytesRangeAndContentType(t *testing.T) {
	for _, scenario := range []struct {
		name  string
		key   []byte
		query string
	}{
		{name: "encrypted", key: []byte("0123456789abcdef")},
		{name: "encrypted_with_playlist_query", key: []byte("0123456789abcdef"), query: "?format=m3u8"},
		{name: "playlist_query", query: "?format=m3u8"},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			segment := bytes.Repeat([]byte{0xa5}, 1024)
			upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/index.m3u8":
					w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
					io.WriteString(w, "#EXTM3U\n#EXT-X-KEY:METHOD=AES-128,URI=\"key\"\n#EXTINF:2,\nsegment.ts"+scenario.query+"\n#EXT-X-ENDLIST\n")
				case "/key":
					io.WriteString(w, "0123456789abcdef")
				case "/segment.ts":
					w.Header().Set("Content-Type", "video/mp2t")
					if r.Header.Get("Range") == "bytes=0-1023" {
						w.Header().Set("Content-Range", "bytes 0-1023/1024")
						w.WriteHeader(http.StatusPartialContent)
					}
					if r.Method != http.MethodHead {
						w.Write(segment)
					}
				default:
					http.NotFound(w, r)
				}
			}))
			defer upstream.Close()
			downloader := sourceFixtureDownloader(t, nil)
			downloader.client = upstream.Client()
			stream, err := newNativeStreamServer(downloader)
			if err != nil {
				t.Fatal(err)
			}
			defer stream.server.Close()
			address, token := stream.nativeOpen(providerMedia{URL: upstream.URL + "/index.m3u8", HLSKey: scenario.key})
			defer stream.nativeRelease(token)
			response, err := http.Get(address)
			if err != nil {
				t.Fatal(err)
			}
			playlist, err := io.ReadAll(response.Body)
			response.Body.Close()
			if err != nil || response.StatusCode != http.StatusOK {
				t.Fatalf("playlist failed: status=%d err=%v", response.StatusCode, err)
			}
			key := nativePlaylistURI.FindStringSubmatch(string(playlist))
			if len(key) != 2 {
				t.Fatal("rewritten key missing")
			}
			response, err = http.Get(key[1])
			if err != nil {
				t.Fatal(err)
			}
			keyBytes, err := io.ReadAll(response.Body)
			response.Body.Close()
			if err != nil || string(keyBytes) != "0123456789abcdef" {
				t.Fatal("player key changed", err)
			}
			mediaURL := ""
			for _, line := range strings.Split(string(playlist), "\n") {
				if strings.HasPrefix(line, "http://") {
					mediaURL = line
					break
				}
			}
			if mediaURL == "" {
				t.Fatal("rewritten segment missing")
			}
			for _, method := range []string{http.MethodGet, http.MethodHead} {
				request, _ := http.NewRequest(method, mediaURL, nil)
				request.Header.Set("Range", "bytes=0-1023")
				response, err = http.DefaultClient.Do(request)
				if err != nil {
					t.Fatal(err)
				}
				body, err := io.ReadAll(response.Body)
				response.Body.Close()
				if err != nil || response.StatusCode != http.StatusPartialContent || response.Header.Get("Content-Type") != "video/mp2t" || response.Header.Get("Content-Range") != "bytes 0-1023/1024" {
					t.Fatalf("%s segment failed: status=%d type=%q range=%q body=%q err=%v", method, response.StatusCode, response.Header.Get("Content-Type"), response.Header.Get("Content-Range"), body, err)
				}
				if method == http.MethodGet && !bytes.Equal(body, segment) {
					t.Fatal("encrypted segment bytes changed")
				}
			}
		})
	}
}
