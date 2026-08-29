import { withInstall } from "@/utils/withInstall";

const modules: Record<string, any> = import.meta.glob(["./*.ts", "!./index.ts"], {
    eager: true,
});

export default withInstall(app => {
    for(const file in modules) {
        const name = file.split(/[\/.]/).at(-2);
        const module = modules[file].default || modules[file][name || ""];
        if(name && module) {
            app.directive(name, module);
        }
    }
});