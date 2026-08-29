class Emun<T extends string | number> {
    private _data: T[] = [];
    constructor(data: T[]) {
        this._data = data;
    }

    get data() {
        return this._data;
    }

    index(value: T) {
        return this._data.findIndex(v => v === value);
    }

    pre(v: T) {
        const i = this.index(v);
        if(i !== -1 && i > 0) {
            return this._data[i - 1];
        }
        return this._data[this._data.length - 1];
    }

    next(v: T) {
        const i = this.index(v);
        if(i !== -1 && i < this._data.length - 1) {
            return this._data[i + 1];
        }
        return this._data[0];
    }
}

export function createEmun<T extends string | number>(data: T[]) {
    return new Emun(data);
}