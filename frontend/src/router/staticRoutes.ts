import type { RouteRecordRaw } from "vue-router";

export const staticRoutes: RouteRecordRaw[] = [
    {
        path: "/",
        redirect: "/home",
    },
    {
        path: "/home",
        name: "Home",
        redirect: "/home",
        meta: {
            keepAlive: true,
            showMenu: true,
        },
        component: () => import("@/views/home/index.vue"),
        children: [
            {
                path: "",
                name: "Workbench",
                meta: { keepAlive: true },
                component: () => import("@/views/home/views/workbench/index.vue"),
            },
            {
                path: "games",
                name: "Games",
                meta: { keepAlive: true },
                component: () => import("@/views/home/views/games/index.vue"),
            },
            {
                path: "music",
                name: "Music",
                meta: { keepAlive: true },
                component: () => import("@/views/home/views/music/index.vue"),
            },
            {
                path: "live",
                name: "Live",
                meta: { keepAlive: true },
                component: () => import("@/views/home/views/live/index.vue"),
                children: [
                    {
                        path: "",
                        name: "LiveMine",
                        meta: {
                            activePath: "/home/live",
                            keepAlive: true,
                            keepType: "live",
                        },
                        component: () => import("@/views/home/views/live/mine.vue"),
                    },
                    {
                        path: "competition",
                        name: "LiveCompetition",
                        meta: {
                            activePath: "/home/live",
                            keepAlive: true,
                            keepType: "live",
                        },
                        component: () => import("@/views/home/views/live/competition.vue"),
                    },
                    // {
                    //     path: "search",
                    //     name: "LiveSearch",
                    //     meta: {
                    //         activePath: "/home/live",
                    //         keepAlive: true,
                    //         keepType: "live",
                    //     },
                    //     component: () => import("@/views/home/views/live/search.vue"),
                    // },
                ],
            },
            {
                path: "anime",
                name: "Anime",
                meta: { keepAlive: true },
                redirect: "/home/anime",
                component: () => import("@/views/home/views/anime/index.vue"),
                children: [
                    {
                        path: "",
                        name: "AnimeList",
                        meta: {
                            activePath: "/home/anime",
                            keepAlive: true,
                            keepType: "anime",
                        },
                        component: () => import("@/views/home/views/anime/list.vue"),
                    },
                    {
                        path: "search",
                        name: "AnimeSearch",
                        meta: {
                            activePath: "/home/anime",
                            keepAlive: true,
                            keepType: "anime",
                        },
                        component: () => import("@/views/home/views/anime/search.vue"),
                    },
                    {
                        path: "schedule",
                        name: "AnimeSchedule",
                        meta: {
                            activePath: "/home/anime",
                            keepAlive: true,
                            keepType: "anime",
                        },
                        component: () => import("@/views/home/views/anime/schedule.vue"),
                    },
                ],
            },
            {
                path: "video",
                name: "Video",
                meta: { keepAlive: true },
                redirect: "/home/video",
                component: () => import("@/views/home/views/video/index.vue"),
                children: [
                    {
                        path: "",
                        name: "VideoList",
                        meta: {
                            activePath: "/home/video",
                            keepAlive: true,
                            keepType: "video",
                        },
                        component: () => import("@/views/home/views/video/list.vue"),
                    },
                    {
                        path: "search",
                        name: "VideoSearch",
                        meta: {
                            activePath: "/home/video",
                            keepAlive: true,
                            keepType: "video",
                        },
                        component: () => import("@/views/home/views/video/search.vue"),
                    },
                ],
            },
        ],
    },
    {
        path: "/pubg",
        name: "Pubg",
        meta: {
            title: "PUBG",
        },
        component: () => import("@/views/pubg/index.vue"),
        children: [
            {
                path: "",
                name: "PubgHome",
                component: () => import("@/views/pubg/pubg-home/index.vue"),
            },
            {
                path: "map/:name",
                name: "PubgMap",
                component: () => import("@/views/pubg/pubg-map/index.vue"),
            },
        ],
    },
    {
        path: "/sub",
        name: "SubLayout",
        component: () => import("@/layout/SubLayout.vue"),
        children: [
            {
                path: "setting",
                name: "Setting",
                meta: {
                    subTitle: "设置",
                },
                component: () => import("@/views/setting/index.vue"),
            },
            {
                path: "download",
                name: "Download",
                meta: {
                    subTitle: "下载管理",
                },
                component: () => import("@/views/download/index.vue"),
            },
            {
                path: "message",
                name: "Message",
                meta: {
                    subTitle: "消息记录",
                },
                component: () => import("@/views/message/index.vue"),
            },
            {
                path: "source",
                name: "Source",
                meta: {
                    subTitle: "源管理",
                },
                component: () => import("@/views/source/index.vue"),
            },
        ],
    },
    {
        path: "/video-play",
        name: "VideoPlay",
        component: () => import("@/views/video-play/index.vue"),
    },
    {
        path: "/live-play",
        name: "LivePlay",
        component: () => import("@/views/live-play/index.vue"),
    },
    {
        path: "/iframe/:link",
        name: "Iframe",
        component: () => import("@/views/iframe/index.vue"),
        children: [],
    }, {
        path: "/redirect/:path*",
        name: "Redirect",
        component: () => import("@/views/redirect/index.vue"),
    }, {
        path: "/:pathMatch(.*)",
        name: "NotFound",
        component: () => import("@/views/error/404.vue"),
    },
];

export const keepNames = getKeepNames(staticRoutes);

export function getKeepNames(data: RouteRecordRaw[]) {
    return data.reduce<string[]>((pre, route) => {
        if(route.meta?.keepAlive) {
            pre.push(route.name as string);
        }
        pre.push(...getKeepNames(route.children || []));
        return pre;
    }, []);
}