# Nyaa-Api-Go
Nyaa API is an **Unofficial Nyaa API**. It just scrapes the website to satisfy the need for an API.

## Usage
***Note:*** The API is hosted at **heroku**, so it might be slow to respond.

**The API is available at:** ```https://nyaa-api-go.herokuapp.com/```

- ### Nyaa Search
	- #### Available Endpoints
	  **If no arguments are specified, latest uploaded torrents are shown**
	  | **Arguments** | **Description** |
	  |------|------|
	  | ```q``` **(Optional)** | Search query. |
	  | ```s``` **(Optional)** | Sorting parameter |
	  | ```p``` **(Optional)** | Page number |
	  | ```o``` **(Optional)** | Order of sorting. Defaults to ***Descending order***. |

		- **Endpoints:**
		  | **Category** | **Endpoint** |
		  |---------|---------|
		  | Anime | ```/anime``` |
		  | Manga | ```/manga``` |
		  | Audio | ```/audio``` |
		  | Pictures | ```/pictures``` |
		  | Live Action | ```/live_action``` |
		  | Software | ```/software``` |

		- **Sub-Categories:**
		  | **Category** | **Sub-Category** |
		  |------|------|
		  | Anime | ```/amv```, ```/eng```, ```/non-eng```, ```/raw``` |
		  | Manga | ```/eng```, ```/non-eng```, ```/raw``` |
		  | Audio | ```/lossy```, ```/lossless``` |
		  | Pictures | ```/photos```, ```/graphics``` |
		  | Live Action | ```/promo```, ```/eng```, ```/non-eng```, ```/raw``` |
		  | Software | ```/application```, ```/games``` |

		- **Sorting:**
		  | **Arguments** | **Methods**  |
		  | ---- | ---- |
		  | Sort | ```size```, ```seeders```, ```leechers```, ```date```, ```downloads``` |
		  | Order | ```asc```, ```desc``` |
	- #### Examples
		-  ```https://nyaa-api-go.herokuapp.com/anime?q={search_query}```
		-  ```https://nyaa-api-go.herokuapp.com/anime?q={search_query}```
		-  ```https://nyaa-api-go.herokuapp.com/anime/{sub_category}?q={search_query}```
		-  ```https://nyaa-api-go.herokuapp.com/anime/{sub_category}?q={search_query}&s={sorting_parameter}```
		-  ```https://nyaa-api-go.herokuapp.com/anime/{sub_category}?q={search_query}&s={sorting_parameter}&p={page_number}```
		-  ```https://nyaa-api-go.herokuapp.com/anime/{sub_category}?q={search_query}&s={sorting_parameter}&p={page_number}&o={order}```