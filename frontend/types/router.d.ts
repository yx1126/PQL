import "vue-router";

declare module "vue-router" {
    interface RouteMeta {
        readonly title?: string;
        readonly subTitle?: string;
        readonly showMenu?: boolean;
        readonly activePath?: string;
        readonly keepAlive?: boolean;
        readonly keepType?: "games" | "music" | "video" | "live" | "anime";
    }
}