// 生成随机id
export function getRandom() {
    if("crypto" in window && crypto?.randomUUID) {
        return crypto?.randomUUID?.();
    }
    return Math.random().toString(36).substring(2, 15);
}