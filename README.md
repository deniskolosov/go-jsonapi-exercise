# go-jsonapi-exercise
Implementation of JSONAPI protocol in Golang

### Usage:

`./go-jsonapi-exercise -p :8081`

where `-p` is `port` option. You can omit it, default value is :8080.

You can test this app using CURL:

>`$ curl -i -H "Accept: application/vnd.api+json" -H "Content-Type:application/vnd.api+json" http://localhost:<port>/<endpoint>`

and reply should be something like:

```json
{
	"data": [
		{
			"type": "comments",
			"id": "1",
			"attributes": {
				"content": "Very interesting post!",
				"created_at": 11111112,
				"post_id": 1,
				"updated_at": 2323233
			}
		},
		{
			"type": "comments",
			"id": "2",
			"attributes": {
				"content": "lol!",
				"created_at": 11111113,
				"post_id": 1,
				"updated_at": 232323233
			}
		},
		{
			"type": "comments",
			"id": "3",
			"attributes": {
				"content": "You're an idiot!",
				"created_at": 1111112,
				"post_id": 2,
				"updated_at": 343444444
			}
		}
	]
}
```

Application endpoints are:

`/posts`

`/comments`

`/tags`

They represent following BD structure:

![screenshot.png](https://raw.githubusercontent.com/thefivekey/go-jsonapi-exercise/master/screenshot.png)
