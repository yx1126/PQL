import Decimal from "decimal.js";

export function decimalPlus(...values: Decimal.Value[]) {
    return values.reduce<number>((pre, item) => {
        // console.log(new Decimal(pre).plus(new Decimal(item)));
        if(item) {
            return new Decimal(pre).plus(new Decimal(item)).toNumber();
        }
        return pre;
    }, 0);
}

export function decimalMul(...values: Decimal.Value[]) {
    const [initialValue, ...vs] = values;
    return vs.reduce<number>((pre, item) => {
        if(item) {
            return new Decimal(pre).mul(item).toNumber();
        }
        return pre;
    }, initialValue as number);
}

export function decimalDiv(value: Decimal.Value, div: number) {
    return new Decimal(value).div(div).toNumber();
}

export function decimalFixedDiv(value: Decimal.Value, div: number, fractionDigits: number) {
    return Number(new Decimal(value).div(div).toNumber().toFixed(fractionDigits));
}

export function decimalSqrt(value: Decimal.Value, fractionDigits?: number) {
    if(fractionDigits !== undefined) {
        return Number(new Decimal(value).sqrt().toFixed(fractionDigits));
    }
    return new Decimal(value).sqrt().toNumber();
}