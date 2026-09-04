export {};

declare module "vue" {
    // export interface ComponentCustomProperties {

    // }

    // export interface GlobalDirectives {
    // }

    export interface GlobalComponents {
        // element-plus extra components
        ElList: typeof import("../src/plugins/element-plus/components/List/List.vue")["default"];
        ElListItem: typeof import("../src/plugins/element-plus/components/List/ListItem.vue")["default"];
        ElThing: typeof import("../src/plugins/element-plus/components/Thing/Thing.vue")["default"];
        ElCardV2: typeof import("../src/plugins/element-plus/components/CardV2/CardV2.vue")["default"];
        // local global components
        Icon: typeof import("../src/components/GlobalRegister/Icon/Icon.vue")["default"];
        Pagination: typeof import("element-plus")["ElPagination"];
        // w-ui
        WSiderMenu: typeof import("../src/components/WUI/layout/WSiderMenu.vue")["default"];
        WSiderLayout: typeof import("../src/components/WUI/layout/WSiderLayout.vue")["default"];
        WPlayerLayout: typeof import("../src/components/WUI/layout/WPlayerLayout.vue")["default"];
        WSiderActions: typeof import("../src/components/WUI/layout/WSiderActions.vue")["default"];

        WBacktop: typeof import("../src/components/WUI/components/WBacktop.vue")["default"];
        WButton: typeof import("../src/components/WUI/components/WButton.vue")["default"];
        WImage: typeof import("../src/components/WUI/components/WImage.vue")["default"];
        WPaging: typeof import("../src/components/WUI/components/WPaging.vue")["default"];
        WVideo: typeof import("../src/components/WUI/components/WVideo.vue")["default"];
        WCard: typeof import("../src/components/WUI/components/WCard.vue")["default"];
        WText: typeof import("../src/components/WUI/components/WText.vue")["default"];
    }
}