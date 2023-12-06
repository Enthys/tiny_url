export interface IShortUrl {
    id: number;
    url: string;
    shortUrl: string;
}

export class ShortUrl implements IShortUrl {
    public id: number;

    public url: string;

    public shortUrl: string;

    constructor(id: number, url: string, shortUrl: string) {
        this.id = id;
        this.url = url;
        this.shortUrl = shortUrl;
    }

    public getShortUrl(): string {
        return `${import.meta.env.VITE_APP_SERVER_URL}/url/${this.shortUrl}`
    }
}

export function getShortUrls(page: number, limit: number = 20): Promise<Map<number, ShortUrl>> {
    return fetch(`${import.meta.env.VITE_APP_SERVER_URL}/v1/short-urls?page=${page}&page_size=${limit}`)
        .then((res) => res.json())
        .then(({short_urls}) => {
            return short_urls.reduce((result: Map<number, ShortUrl>, shortUrl: any) => {
                result.set(shortUrl.id, new ShortUrl(shortUrl.id, shortUrl.url, shortUrl.short_url));

                return result;
            }, new Map<number, ShortUrl>());
        });
}

export function createShortUrl(url: string): Promise<ShortUrl> {
    return fetch(`${import.meta.env.VITE_APP_SERVER_URL}/v1/short-urls`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url })
    })
        .then((res) => res.json())
        .then(({short_url}) => new ShortUrl(short_url.id, short_url.url, short_url.short_url));
}