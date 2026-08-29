import { isArray } from "./validata";
import type { LocationQueryValue } from "vue-router";

/**
 * 格式化标题
 * @example
 * ```js
 * const query = ["test", "test1"];
 * const template = "{0}-{1}";
 * formatMenuTitle(query, template);  // "test-test1"
 *
 * const template = "{0}-{1}-{2}";
 * formatMenuTitle(query, template);  // "test-test1-"
 *
 * const template = "{0}-{1}-{2|test2}";
 * formatMenuTitle(query, template);  // "test-test1-test2"
 * ```
 */
export function formatMenuTitle(title: LocationQueryValue | LocationQueryValue[], template?: string) {
    const custom = isArray<LocationQueryValue[]>(title) ? title : [title];
    return template?.replace(/\{(\d+)(\|(.*))?\}/g, (...args) => {
        return custom[args[1]] || args[3] || "";
    });
}