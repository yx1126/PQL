import type { MaybeRef, MaybeRefOrGetter, VNode } from "vue";

export {};

declare global {

    export type Unit = string | number;

    export type Noop = () => void;

    export type ObjKeys<T> = keyof T;

    export type Nullable<T> = T | null;

    export type Noable<T> = T | undefined;

    export type Empty<T> = Nullable<Noable<T>>;

    export type Arrayable<T> = T extends Array<infer A> ? A[] : T[];

    export type ArrayGet<T> = T extends Array<infer A> ? A : T;

    export type Writable<T> = {
        -readonly [P in keyof T]: T[P];
    };

    export type DeepPartial<T> = {
        [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P];
    };

    export type DeepReadonly<T> = {
        readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P];
    };

    export type DeepWritable<T> = {
        -readonly [P in keyof T]: T[P] extends object ? DeepWritable<T[P]> : T[P];
    };

    export type UnRef<T> = T extends MaybeRef<infer V> ? V : T extends MaybeRefOrGetter<infer CV> ? CV : T;

    export type UnRefable<T extends object> = {
        [K in keyof T]: UnRef<T[K]>
    };

    export type Slotsable<T extends string> = {
        [K in T]?: () => VNode[];
    };

    export type EventOn<T extends string> = `on${Capitalize<T>}`;

    export interface Models {
        readonly id?: number;
        readonly createdAt?: string;
        readonly updatedAt?: string;
    }

    export type PartialId<T extends object> = Omit<T, "id"> & {
        id?: number;
    };
}