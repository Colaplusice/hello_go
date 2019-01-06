package matchers

import "encoding/xml"

// 搜索rss源

type item struct {
	XMLName xml.Name `xml:"item"`
	PubDate string `xml:"pubDate"`
	Title string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	GUID string `xml:"guid"`
	GeoRssPoint string `xml:"georss:point"`


}