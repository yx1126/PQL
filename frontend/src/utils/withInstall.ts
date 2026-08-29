import type { App, Plugin, DefineComponent, AppContext } from "vue";

export type WithInstall<T> = T & Plugin;

type WithInstallContext<T> = WithInstall<T> & {
    _context: AppContext | null;
};

function toInstall<T extends DefineComponent>(app: App, comps: Array<T>) {
    for(const comp of comps) {
        if(import.meta.env.MODE === "development" && !(comp.name || comp.__name)) {
            throw new Error("The component name cannot be empty");
        }
        app.component(comp.name || comp.__name!, comp);
    }
}

export function withInstall<T>(main: T[]): Plugin;
export function withInstall(main: (app: App) => void): Plugin;
export function withInstall<T>(main: T): WithInstall<T>;
export function withInstall(main: unknown) {
    if(Array.isArray(main)) {
        return {
            install: (app: App) => {
                toInstall(app, main);
            },
        };
    }
    if(typeof main === "function") {
        return {
            install: main,
        };
    }
    (main as DefineComponent).install = function(app: App) {
        toInstall(app, [main as DefineComponent]);
    };
    return main;
}

export function withInstallFunction<T extends Function>(fn: T, name: string) {
    (fn as WithInstallContext<T>).install = (app: App) => {
        (fn as WithInstallContext<T>)._context = app._context;
        app.config.globalProperties[name] = fn;
    };
    return fn as WithInstallContext<T>;
}

export function withInstallProps<T extends object>(props: T): Plugin;
export function withInstallProps<T>(prop: T, name: string): Plugin;
export function withInstallProps(prop: unknown, name?: string): Plugin {
    const installs: Record<string, any> = arguments.length === 2 ? { [name!]: prop } : Object.assign({}, prop);
    return {
        install(app) {
            for(const name in installs) {
                app.config.globalProperties[name] = installs[name];
            }
        },
    };
}