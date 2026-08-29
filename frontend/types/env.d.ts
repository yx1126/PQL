/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_APP_TITLE: string;
    readonly VITE_APP_TITLE_SHORT: string;
    readonly VITE_APP_BASE_URL: string;
    readonly VITE_APP_MAIN_NAME: string;
    // 更多环境变量...
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}