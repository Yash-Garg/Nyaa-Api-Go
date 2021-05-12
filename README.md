# Nyaa-Api-Go

Nyaa API is an **Unofficial Nyaa API** used for searching torrents.

Made for fun, you can use [Nyaa API](https://github.com/Vivek-Kolhe/Nyaa-API) by [@Vivek-Kolhe](https://github.com/Vivek-Kolhe/) instead if you want more endpoints.

## Usage

-   #### Available Endpoints

    **If no arguments are specified, latest uploaded torrents are shown**
    | **Arguments** | **Description** |
    |------|------|
    | `q` **(Optional)** | Search query. |
    | `s` **(Optional)** | Sorting parameter |
    | `p` **(Optional)** | Page number |
    | `o` **(Optional)** | Order of sorting. Defaults to **_Descending order_**. |

    -   **Endpoints:**
        | **Category** | **Endpoint** |
        |---------|---------|
        | Anime | `/anime` |
        | Manga | `/manga` |
        | Audio | `/audio` |
        | Pictures | `/pictures` |
        | Live Action | `/live_action` |
        | Software | `/software` |

    -   **Sub-Categories:**
        | **Category** | **Sub-Category** |
        |------|------|
        | Anime | `/amv`, `/eng`, `/non-eng`, `/raw` |
        | Manga | `/eng`, `/non-eng`, `/raw` |
        | Audio | `/lossy`, `/lossless` |
        | Pictures | `/photos`, `/graphics` |
        | Live Action | `/promo`, `/eng`, `/non-eng`, `/raw` |
        | Software | `/application`, `/games` |

    -   **Sorting:**
        | **Arguments** | **Methods** |
        | ---- | ---- |
        | Sort | `size`, `seeders`, `leechers`, `date`, `downloads` |
        | Order | `asc`, `desc` |

-   #### Examples
    -   `https://nyaa-api-go.herokuapp.com/anime?q={search_query}`
    -   `https://nyaa-api-go.herokuapp.com/anime?q={search_query}`
    -   `https://nyaa-api-go.herokuapp.com/anime/{sub_category}?q={search_query}`
    -   `https://nyaa-api-go.herokuapp.com/anime/{sub_category}?q={search_query}&s={sorting_parameter}`
    -   `https://nyaa-api-go.herokuapp.com/anime/{sub_category}?q={search_query}&s={sorting_parameter}&p={page_number}`
    -   `https://nyaa-api-go.herokuapp.com/anime/{sub_category}?q={search_query}&s={sorting_parameter}&p={page_number}&o={order}`