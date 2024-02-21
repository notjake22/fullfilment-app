package upc

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"main/cache"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func InitializeUPCLookup(upc string, ctx context.Context, cacheRef cache.Cache) (*LookUp, error) {
	u := &LookUp{
		UPC:         upc,
		item:        ItemData{},
		shouldCache: false,
		Client:      http.Client{},
		Context:     ctx,
		cache:       cacheRef,
	}

	return u, nil
}

func (u *LookUp) proxyHandler() error {
	proxyU := getProxy()

	if proxyU != "" {
		parts := strings.Split(proxyU, ":")
		if len(parts) == 4 {
			newProxy := fmt.Sprintf("%s:%s@%s:%s", parts[2], parts[3], parts[0], strings.Replace(parts[1], "]", "", -1))
			proxyUrl, err := url.Parse("http://" + newProxy)
			if err != nil {
				return err
			}
			ipPort, err := url.Parse(fmt.Sprintf("http://%s:%s", proxyUrl.Hostname(), proxyUrl.Port()))
			if err != nil {
				return err
			}

			pw, _ := proxyUrl.User.Password()
			userpass := fmt.Sprintf("%s:%s", proxyUrl.User.Username(), pw)
			u.Client = http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(ipPort),
					ProxyConnectHeader: http.Header{
						"Proxy-Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(userpass)))},
					},
				},
			}
		} else if len(parts) == 2 {
			newProxy := fmt.Sprintf("%s:%s", parts[0], strings.Replace(parts[1], "]", "", -1))
			proxyUrl, err := url.Parse("http://" + newProxy)
			if err != nil {
				return err
			}
			ipPort, err := url.Parse(fmt.Sprintf("http://%s:%s", proxyUrl.Hostname(), proxyUrl.Port()))
			if err != nil {
				return err
			}

			u.Client = http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(ipPort),
				},
			}

		} else {
			log.Println("Bad proxy, using local")
		}
	} else {
		return nil
	}
	return nil
}

func getProxy() string {
	file, err := os.ReadFile("./proxies.txt")
	if err != nil {
		panic(err)
	}
	proxies := string(file)
	if proxies == "" {
		return ""
	}
	proxies = strings.Replace(proxies, "\r", "", -1)
	proxiesArray := strings.Split(proxies, "\n")

	return proxiesArray[rand.Intn(len(proxiesArray))]
}
