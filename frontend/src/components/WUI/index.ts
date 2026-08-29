import { withInstall } from "@/utils/withInstall";
import WButton from "./components/WButton.vue";
import WVideo from "./components/WVideo.vue";
import WBacktop from "./components/WBacktop.vue";
import WPaging, { type PagingType } from "./components/WPaging.vue";
import WImage from "./components/WImage.vue";
import WSiderMenu, { type MenuSiderItem } from "./components/WSiderMenu.vue";
import WSiderLayout from "./components/WSiderLayout.vue";
import WCard from "./components/WCard.vue";
import WPlayerLayout from "./components/WPlayerLayout.vue";

export {
    type PagingType,
    type MenuSiderItem,
};

export default withInstall(app => {
    [
        WButton,
        WVideo,
        WBacktop,
        WPaging,
        WImage,
        WSiderMenu,
        WSiderLayout,
        WCard,
        WPlayerLayout,
    ].forEach(v => {
        app.component(v.name!, v);
    });
});