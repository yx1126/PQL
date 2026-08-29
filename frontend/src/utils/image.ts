class Img {
    private cache = new Map<string, HTMLImageElement>();
    private urlCache = new Map<string, string>();

    async loadLocal(img: string) {
        if(this.cache.has(img)) {
            return this.cache.get(img)!;
        }
        const value = await this.request(`/local/pubg/${img}`);
        this.cache.set(img, value);
        return value;
    }

    async load(img: string) {
        if(this.cache.has(img)) {
            return this.cache.get(img)!;
        }
        const value = await this.request(img);
        this.cache.set(img, value);
        return value;
    }

    async set(key: string, value: string | HTMLImageElement) {
        if(typeof value === "string") {
            const image = await this.request(value);
            this.cache.set(key, image);
            return;
        }
        this.cache.set(key, value);
    }

    get(key: string) {
        return this.cache.get(key);
    }

    has(key: string) {
        return this.cache.has(key);
    }

    clear() {
        this.cache.clear();
        this.urlCache.clear();
    }

    async loadSvg(svg: string) {
        if(this.cache.has(svg)) {
            return this.cache.get(svg)!;
        }
        let url = this.urlCache.get(svg);
        if(!url) {
            const blob = new Blob([svg], { type: "image/svg+xml;charset=utf-8" });
            url = URL.createObjectURL(blob);
        }
        const image = await this.request(url);
        this.cache.set(svg, image);
        this.urlCache.set(svg, url);
        return image;
    }

    async revoke(svg: string) {
        const url = this.urlCache.get(svg);
        if(url) {
            URL.revokeObjectURL(url);
            this.urlCache.delete(svg);
        }
    }

    private request(img: string) {
        return new Promise<HTMLImageElement>((resolve, reject) => {
            const image = new Image();
            image.src = img;
            image.onload = () => {
                resolve(image);
            };
            image.onerror = error => {
                reject(error);
            };
        });
    }
}

export default new Img();