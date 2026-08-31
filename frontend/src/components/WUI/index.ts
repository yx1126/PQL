import { withInstall } from "@/utils/withInstall";

// layout
import WSiderMenu, { type MenuSiderItem } from "./components/WSiderMenu.vue";
import WSiderLayout from "./components/WSiderLayout.vue";
import WPlayerLayout from "./components/WPlayerLayout.vue";

// basic
import WButton from "./components/WButton.vue";
import WVideo from "./components/WVideo.vue";
import WBacktop from "./components/WBacktop.vue";
import WPaging, { type PagingType } from "./components/WPaging.vue";
import WImage from "./components/WImage.vue";
import WCard from "./components/WCard.vue";
import WText from "./components/WText.vue";

export {
    type PagingType,
    type MenuSiderItem,
};

export default withInstall(app => {
    [
        WSiderMenu,
        WSiderLayout,
        WPlayerLayout,
        WButton,
        WVideo,
        WBacktop,
        WPaging,
        WImage,
        WCard,
        WText,
    ].forEach(v => {
        app.component(v.name!, v);
    });
});