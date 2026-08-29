import { defineCustomElement } from "vue";
import LoadIcon from "./LoadIcon.ce.vue";
import { kebabCase } from "lodash-es";

const WebComponents = [
    LoadIcon,
];

WebComponents.forEach(component => {
    if(component.name) {
        customElements.define(kebabCase(component.name), defineCustomElement(component));
    } else {
        throw new Error("[customElements]: The component name cannot be empty!");
    }
});