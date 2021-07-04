# Nyaa-Api-Go

This API is an **Unofficial Nyaa API** used for searching torrents.

Made for fun, you can use [Nyaa API](https://github.com/Vivek-Kolhe/Nyaa-API) by [@Vivek-Kolhe](https://github.com/Vivek-Kolhe/) instead if you want more endpoints.

## Usage

- `username` and `id` are required parameters if using `/id/{id}` and `/user/{username}` endpoints.

- If no parameters are specified in other endpoints like `/anime`, `/manga`, etc. It will return the latest uploaded torrents in the respective category.

- For Filters, input `filter=1` for *No Remakes* and `filter=2` for *Trusted Only*.

-   #### Available Endpoints
	| **Arguments** | **Description** |
	|------|------|
	| `q` **(Optional)** | Search query. |
	| `s` **(Optional)** | Sorting parameter |
	| `p` **(Optional)** | Page number |
	| `f` **(Optional)** | Filter Option |
	| `o` **(Optional)** | Order of sorting. Defaults to **_Descending order_**. |

	-   **Endpoints**
		| **Category** | **Endpoint** |
		|---------|---------|
		| All | `/all` |
		| Anime | `/anime` |
		| Manga | `/manga` |
		| Audio | `/audio` |
		| Pictures | `/pictures` |
		| Live Action | `/live_action` |
		| Software | `/software` |
		| ID | `/id` |
		| User | `/user` |

	-   **Sub-Categories** (Not applicable for `/user` and `/id`)
		| **Category** | **Sub-Category** |
		|------|------|
		| Anime | `/amv`, `/eng`, `/non-eng`, `/raw` |
		| Manga | `/eng`, `/non-eng`, `/raw` |
		| Audio | `/lossy`, `/lossless` |
		| Pictures | `/photos`, `/graphics` |
		| Live Action | `/promo`, `/eng`, `/non-eng`, `/raw` |
		| Software | `/application`, `/games` |

	-   **Sorting Parameters**
		| **Arguments** | **Methods** |
		| ---- | ---- |
		| Sort | `size`, `seeders`, `leechers`, `date`, `downloads` |
		| Order | `asc`, `desc` |

-	#### Search using ID
	- 	`https://nyaa-api-go.herokuapp.com/id/{id}`

-   #### Search using category
	-   `https://nyaa-api-go.herokuapp.com/{category}?q={search_query}`
	-   `https://nyaa-api-go.herokuapp.com/{category}?q={search_query}&s={sorting_parameter}`
	-   `https://nyaa-api-go.herokuapp.com/{category}?q={search_query}&s={sorting_parameter}&p={page_number}`
	-   `https://nyaa-api-go.herokuapp.com/{category}?q={search_query}&s={sorting_parameter}&p={page_number}&o={order}`
	-   `https://nyaa-api-go.herokuapp.com/{category}?q={search_query}&s={sorting_parameter}&p={page_number}&o={order}&f={filter}`

-   #### Search using sub category
	-   `https://nyaa-api-go.herokuapp.com/{category}/{sub_category}?q={search_query}`
	-   `https://nyaa-api-go.herokuapp.com/{category}/{sub_category}?q={search_query}&s={sorting_parameter}`
	-   `https://nyaa-api-go.herokuapp.com/{category}/{sub_category}?q={search_query}&s={sorting_parameter}&p={page_number}`
	-   `https://nyaa-api-go.herokuapp.com/{category}/{sub_category}?q={search_query}&s={sorting_parameter}&p={page_number}&o={order}`
	-   `https://nyaa-api-go.herokuapp.com/{category}/{sub_category}?q={search_query}&s={sorting_parameter}&p={page_number}&o={order}&f={filter}`

-	#### Search using username
	- 	`https://nyaa-api-go.herokuapp.com/user/{username}`
	- 	`https://nyaa-api-go.herokuapp.com/user/{username}?q={search_query}`
	- 	`https://nyaa-api-go.herokuapp.com/user/{username}?q={search_query}&s={sorting_parameter}`
	- 	`https://nyaa-api-go.herokuapp.com/user/{username}?q={search_query}&s={sorting_parameter}&p={page_number}`
	- 	`https://nyaa-api-go.herokuapp.com/user/{username}?q={search_query}&s={sorting_parameter}&p={page_number}&o={order}`
	- 	`https://nyaa-api-go.herokuapp.com/user/{username}?q={search_query}&s={sorting_parameter}&p={page_number}&o={order}&f={filter}`

## License

```text
MIT License

Copyright (c) 2021 Yash Garg

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
