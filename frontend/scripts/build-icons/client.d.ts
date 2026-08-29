declare module "virtual:svg-icons-list" {
    const iconsNames: Array<{
        readonly name: string;
        readonly symbolId: string;
    }>;
    export default iconsNames;
}

declare module "virtual:svg-icons-load" {
    const iconsNames: string[];
    export default iconsNames;
}