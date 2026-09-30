package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHuangguoAIEpisodesExcludeRecommendations(t *testing.T) {
	body := `<a href="/video/6756/">推荐</a>
 <a href="/video/117/ep-10/" data-ep-id="10">10</a>
 <a href="/video/117/" data-ep-id="1">01</a>
 <a href="/video/117/ep-2/" data-ep-id="2">02</a>
 <a href="/video/117/ep-1/">重复第一集</a>
 <a href="/video/117/ep-3/" data-ep-id="4">不匹配</a>
 <a href="https://unrelated.invalid/video/117/ep-4/">外部推荐</a>`
	episodes := parseHuangguoAIEpisodes(body, "https://huangguoai.com/video/117/", "117")
	if len(episodes) != 3 {
		t.Fatalf("got %d episodes, want 3", len(episodes))
	}
	for i, number := range []int{1, 2, 10} {
		if episodes[i].Index != number || episodes[i].Key != fmt.Sprintf("ep-%d", number) {
			t.Fatalf("episode %d: %+v", i, episodes[i])
		}
	}
}

func TestHuangguoAIInitialDataSelectsRequestedEpisode(t *testing.T) {
	for _, scenario := range []struct{ name, data, page, want string }{
		{"first", `{"id":117,"ep":1,"videoSrc":"https://media.invalid/one.m3u8"}`, "/video/117/", "https://media.invalid/one.m3u8"},
		{"exact_map", `{"id":"117","ep":1,"videoSrc":"https://media.invalid/one.m3u8","epPlaySrcs":{"1":"https://media.invalid/one.m3u8","2":"https://media.invalid/two.m3u8"}}`, "/video/117/ep-2/", "https://media.invalid/two.m3u8"},
		{"current", `{"id":"117","ep":3,"videoSrc":"https://media.invalid/three.m3u8","epPlaySrcs":{"1":"https://media.invalid/one.m3u8"}}`, "/video/117/ep-3/", "https://media.invalid/three.m3u8"},
		{"missing", `{"id":"117","ep":1,"videoSrc":"https://media.invalid/one.m3u8","epPlaySrcs":{"1":"https://media.invalid/one.m3u8"}}`, "/video/117/ep-3/", ""},
		{"foreign", `{"id":"6756","ep":1,"videoSrc":"https://media.invalid/recommendation.m3u8"}`, "/video/117/", ""},
		{"invalid_json", `{broken`, "/video/117/", ""},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			body := `<script id="videoInitialData" type="application/json">` + scenario.data + `</script><div data-play-src="https://media.invalid/advertisement.m3u8"></div>`
			if got := parseAIVideoURL(body, "https://huangguoai.com"+scenario.page); got != scenario.want {
				t.Fatalf("got %q want %q", got, scenario.want)
			}
		})
	}
}

func TestHuangguoAIRejectsUnidentifiedMedia(t *testing.T) {
	body := `<div data-play-src="https://media.invalid/advertisement.m3u8"></div><script>videoSrc="https://media.invalid/recommendation.m3u8"</script>`
	if media := parseAIVideoURL(body, "https://huangguoai.com/video/117/"); media != "" {
		t.Fatal("unidentified page media accepted", media)
	}
}

func TestHuangguoAIIdentityRejectsStaleChapterAndRedirect(t *testing.T) {
	for _, scenario := range []struct{ name, page, data string }{
		{"stale_recommendation", "/video/6756/", `{"id":"6756","ep":1}`},
		{"wrong_episode", "/video/117/ep-2/", `{"id":"117","ep":2}`},
		{"foreign_data", "/video/117/", `{"id":"6756","ep":1}`},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			if err := validateAIPageIdentity(`<script id="videoInitialData">`+scenario.data+`</script>`, "https://huangguoai.com"+scenario.page, "117", 1); err == nil {
				t.Fatal("identity mismatch accepted")
			}
		})
	}
}

func TestHuangguoAIDetailAndMediaKeepDramaIdentity(t *testing.T) {
	var address string
	upstream := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) >= 2 && parts[0] == "detail" {
			http.Redirect(w, r, "/video/"+parts[1]+"/", http.StatusFound)
			return
		}
		if len(parts) >= 2 && parts[0] == "video" {
			id := parts[1]
			episode := 1
			if len(parts) == 3 {
				fmt.Sscanf(parts[2], "ep-%d", &episode)
			}
			fmt.Fprintf(w, `<h1>Fixture %s</h1><a href="/video/6756/">推荐</a><a href="/video/%s/" data-ep-id="1">01</a><a href="/video/%s/ep-2/" data-ep-id="2">02</a><script id="videoInitialData">{"id":"%s","ep":%d,"videoSrc":"%s/media/%s/%d.m3u8"}</script>`, id, id, id, id, episode, address, id, episode)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/media/") {
			fmt.Fprint(w, "#EXTM3U\n#EXTINF:2,\nsegment.ts\n#EXT-X-ENDLIST\n")
			return
		}
		http.NotFound(w, r)
	}))
	defer upstream.Close()
	address = upstream.URL
	previous := huangguoAIBaseURL
	huangguoAIBaseURL = address
	defer func() { huangguoAIBaseURL = previous }()
	downloader := sourceFixtureDownloader(t, nil)
	downloader.client = upstream.Client()
	for _, id := range []string{"117", "2181"} {
		drama, chapters, err := downloader.fetchHuangguoAIDetail(context.Background(), id)
		if err != nil || len(chapters) != 2 {
			t.Fatalf("detail %s: chapters=%d err=%v", id, len(chapters), err)
		}
		for index, chapter := range chapters {
			media, err := downloader.resolveProviderMedia(context.Background(), Task{DramaID: drama.ID, Chapter: chapter, Index: index + 1})
			want := fmt.Sprintf("%s/media/%s/%d.m3u8", address, id, index+1)
			if err != nil || media.URL != want {
				t.Fatalf("resolve %s episode %d: got %q want %q err=%v", id, index+1, media.URL, want, err)
			}
		}
		stale := Chapter{Source: sourceHuangguoAI, PageURL: address + "/video/6756/", CurrentEpisode: json.RawMessage(`1`)}
		if _, err := downloader.resolveProviderMedia(context.Background(), Task{DramaID: drama.ID, Chapter: stale, Index: 1}); err == nil {
			t.Fatal("stale foreign recommendation played")
		}
	}
}
