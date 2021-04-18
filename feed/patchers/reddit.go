package patchers

// // Reddit ...
// type Reddit struct {
// }

// // NewRedditPatcher ....
// func NewRedditPatcher() Patcher {
// 	return &Reddit{}
// }

// // Patch ...
// func (r *Reddit) Patch(item *gofeed.Item) *models.ItemModel {

// 	thumb := ""
// 	if item.Extensions != nil {
// 		thumb = item.Extensions["media"]["thumbnail"][0].Attrs["url"]
// 	}

// 	_item := &models.ItemModel{
// 		ID:          item.GUID,
// 		Title:       item.Title,
// 		Link:        item.Link,
// 		Description: item.Content,
// 		Thumbnail:   thumb,
// 		Published:   item.PublishedParsed.Local(),
// 		New:         true,
// 		Provider:    "reddit",
// 		Type:        "unknown",
// 		Embed:       item.Link,
// 	}

// 	if id, ok := extractYoutube(_item.Description); ok {
// 		_item.Type = "video"
// 		_item.Embed = embedYoutube(id)
// 	}

// 	if img, ok := extractImage(_item.Description); ok {
// 		_item.Type = "image"
// 		_item.Embed = img
// 	}

// 	return _item
// }

// func extractImage(content string) (string, bool) {

// 	content = html.UnescapeString(content)

// 	re, err := regexp.Compile(`<a href="((?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+(png|jpg|jpeg|gif|svg))">\[link\]</a>`)

// 	if err != nil {
// 		panic(err)
// 	}

// 	found := re.FindAllStringSubmatch(content, 2)

// 	if len(found) > 0 {
// 		return fmt.Sprintf("<img src='%s'>", found[0][1]), true
// 	}

// 	return "", false

// }
