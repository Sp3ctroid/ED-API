# ED-API

This is a small API project that I'm making for educational purposes.

Current end-points:
-   "/" - home page. Says "Hello World"
-   "/albums" - Has two methods assigned to it. GET gives you all albums stored. POST is for creating new records in storage, you send a JSON to it, it responds with a JSON
-   "/albums/id" - change "id" to any number. Has a GET and a PUT request assigned to it. GET gives you album by ID from storage. PUT changes a record in storage based on passed album in form of JSON.

---

# Form of JSON for POST and PUT requests:

```json
{
    "id": Your_id,
    "title": "Your_title",
    "artist": "Your_artist",
    "price": Your_price
}
```

---

# Storage

Storage can be easily swapped, since it is injected into a ```type AlbumHandler struct{}``` as an interface ```type AlbumStore interface{}```. This also helps with testing the API, since we can always make a mock storage that satisfies ```AlbumStore``` interface.

Currently API supports only slice-based storage.
