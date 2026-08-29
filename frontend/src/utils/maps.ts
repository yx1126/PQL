import pointSecret from "@/assets/icon/point-secret.svg";
import pointVehicle from "@/assets/icon/point-vehicle.svg";
import pointGlider from "@/assets/icon/point-glider.svg";
import pointBear from "@/assets/icon/point-bear.svg";
import pointExperiment from "@/assets/icon/point-experiment.svg";
import pointCrowbar from "@/assets/icon/point-crowbar.svg";
import pointSafedoor from "@/assets/icon/point-safedoor.svg";
import { decimalMul, decimalSqrt } from "./decimal";

export type PointType = "secretRoom" | "vehicle" | "glider" | "bear" | "experiment" | "crowbar" | "safedoor";

export interface MapItem {
    label: string;
    name: string;
    size: "4x4" | "8x8";
    thumbnail: string;
    map: string;
    points?: Partial<Record<PointType, Location[]>>;
}

export interface PointIcon {
    name: string;
    icon: string;
    png: string;
}

export const pointIcon: Record<PointType, PointIcon> = {
    secretRoom: {
        name: "密室",
        icon: "point-secret",
        png: pointSecret,
    },
    vehicle: {
        name: "载具",
        icon: "point-vehicle",
        png: pointVehicle,
    },
    glider: {
        name: "滑翔机",
        icon: "point-glider",
        png: pointGlider,
    },
    bear: {
        name: "熊洞",
        icon: "point-bear",
        png: pointBear,
    },
    experiment: {
        name: "实验营地",
        icon: "point-experiment",
        png: pointExperiment,
    },
    crowbar: {
        name: "撬棍房",
        icon: "point-crowbar",
        png: pointCrowbar,
    },
    safedoor: {
        name: "安全门",
        icon: "point-safedoor",
        png: pointSafedoor,
    },
};

export interface Location {
    x: number;
    y: number;
}

export const mapList: MapItem[] = [
    {
        label: "erangel",
        name: "艾伦格",
        size: "8x8",
        thumbnail: "Erangel.jpg",
        map: "Erangel_Main_High_Res.jpg",
        points: {
            secretRoom: [
                { x: 1389.5901036915952, y: 1828.7088167130435 },
                { x: 5166.817292558109, y: 672.9389073099528 },
                { x: 2612.4839778779096, y: 2230.3819189660594 },
                { x: 4158.892002340867, y: 1986.267837592953 },
                { x: 6580.185995764648, y: 2093.7182331436293 },
                { x: 1494.6868647860338, y: 3570.793624742747 },
                { x: 3036.573279821607, y: 3783.1067362599606 },
                { x: 5510.881120449581, y: 3453.2160151057305 },
                { x: 4699.529253905581, y: 4461.8269061806795 },
                { x: 2761.0018955128858, y: 5196.622577337241 },
                { x: 6811.583659181082, y: 4936.777236748235 },
                { x: 1273.8959279285455, y: 5576.139447553096 },
                { x: 4435.508572855535, y: 5967.939891694633 },
                { x: 3328.7643214224504, y: 6758.531467430446 },
                { x: 5720.548807521695, y: 6785.722178015443 },
            ],
            vehicle: [],
            glider: [],
        },
    },
    {
        label: "miramar",
        name: "米拉玛",
        size: "8x8",
        thumbnail: "Miramar.jpg",
        map: "Miramar_Main_High_Res.jpg",
        points: {
            secretRoom: [
                { x: 3268.857597774634, y: 1078.1315971424708 },
                { x: 1794.9346598450177, y: 1677.5604730759133 },
                { x: 4639.33837890625, y: 1447.0625 },
                { x: 6330.418703460078, y: 1957.7167998439922 },
                { x: 2810.5798391794087, y: 2512.176300384196 },
                { x: 1405.633035671858, y: 3326.743559780004 },
                { x: 5201.33750794325, y: 3210.258130874364 },
            ],
            vehicle: [],
            glider: [],
        },
    },
    {
        label: "taego",
        name: "泰戈",
        size: "8x8",
        thumbnail: "Taego.jpg",
        map: "Taego_Main_High_Res.jpg",
        points: {
            secretRoom: [
                { x: 1404.0089843750002, y: 1194.5951171875 },
                { x: 2584.1009765625004, y: 1354.3338867187501 },
                { x: 3594.6351243570234, y: 1982.3072169161621 },
                { x: 4855.35009765625, y: 1726.458740234375 },
                { x: 1265.6264525652591, y: 2711.008540153891 },
                { x: 1022.663134765625, y: 3419.252734375 },
                { x: 972.5293490794824, y: 5260.851456844348 },
                { x: 2433.1797530354443, y: 6481.311372804441 },
                { x: 4445.81201171875, y: 4993.943359375 },
                { x: 6945.079917082954, y: 2086.71767114055 },
                { x: 7132.000604180711, y: 3381.459578865848 },
                { x: 6069.64453125, y: 3886.811767578125 },
                { x: 6445.830467192773, y: 5595.123186907959 },
                { x: 4949.4345703125, y: 6454.86865234375 },
                { x: 6387.72520351065, y: 7220.779284687786 },
            ],
            vehicle: [],
            glider: [],
        },
    },
    {
        label: "sanhok",
        name: "萨诺",
        size: "4x4",
        thumbnail: "Sanhok.jpg",
        map: "Sanhok_Main_High_Res.jpg",
    },
    {
        label: "vikendi",
        name: "维寒迪",
        size: "8x8",
        thumbnail: "Vikendi.jpg",
        map: "Vikendi_Main_High_Res.jpg",
        points: {
            secretRoom: [
                { x: 2767.14501953125, y: 1588.2088623046875 },
                { x: 5444.478515625, y: 1326.7513427734375 },
                { x: 6294.15700670965, y: 2473.38951380031 },
                { x: 4132.170652859582, y: 3233.1297543394667 },
                { x: 1402.2657388962186, y: 3868.9547643280703 },
                { x: 6889.601819338149, y: 3910.763180439186 },
                { x: 4743.4746702781495, y: 4982.306769648186 },
                { x: 2394.884776296455, y: 5664.328551271763 },
                { x: 3978.5477451298375, y: 6583.8526036323865 },
                { x: 6111.013671875, y: 5911.0166015625 },
            ],
            vehicle: [],
            glider: [],
            bear: [
                { x: 2868.847412109375, y: 1367.031982421875 },
                { x: 5298.87550659883, y: 1409.5934753072593 },
                { x: 6102.52516957792, y: 1752.3605799163345 },
                { x: 5746.080458616965, y: 2945.895390584536 },
                { x: 4736.219737465377, y: 3588.2153931381945 },
                { x: 6682.548322037489, y: 3901.5517506801807 },
                { x: 3848.602544424695, y: 6105.250666677936 },
                { x: 3794.8641402294616, y: 6539.297387908438 },
                { x: 5271.425164881342, y: 5940.802986589949 },
                { x: 6280.55789207241, y: 5919.484748826683 },
            ],
            crowbar: [],
            experiment: [
                { x: 2030.403527207859, y: 2696.018031380654 },
                { x: 2281.03955078125, y: 2196.72119140625 },
                { x: 3239.0135810231154, y: 1946.5220903923346 },
                { x: 3299.425048828125, y: 1062.8783569335938 },
            ],
        },
    },
    {
        label: "rondo",
        name: "荣都",
        size: "8x8",
        thumbnail: "Rondo.jpg",
        map: "Rondo_Main_High_Res.jpg",
        points: {
            secretRoom: [
                { x: 1479.8409423828125, y: 1362.573486328125 },
                { x: 3055.537109375, y: 937.2607421875 },
                { x: 5845.3978659038485, y: 903.9684887826324 },
                { x: 5062.173788443852, y: 2067.008038557788 },
                { x: 7090.228431639154, y: 2768.307596983416 },
                { x: 3809.778076171875, y: 2673.599853515625 },
                { x: 1432.0548537994737, y: 3261.5682645544066 },
                { x: 4690.8037109375, y: 4423.8388671875 },
                { x: 1468.3936767578125, y: 4821.216796875 },
                { x: 3032.336181640625, y: 5171.462890625 },
                { x: 1246.4943787635798, y: 6554.084290091562 },
                { x: 3404.479333816936, y: 7040.57274450222 },
                { x: 5717.173842161316, y: 6115.988484218582 },
                { x: 4969.78125, y: 7425.921875 },
                { x: 6675.638938611166, y: 4295.693905715735 },
            ],
            vehicle: [],
            glider: [],
        },
    },
    {
        label: "deston",
        name: "帝斯顿",
        size: "8x8",
        thumbnail: "Deston.jpg",
        map: "Deston_Main_High_Res.jpg",
        points: {
            safedoor: [],
            vehicle: [],
            glider: [],
        },
    },
];

// 获取真实距离
export function getDistance(point1: Location, point2: Location): number {
    const dx = Math.max(point1.x, point2.x) - Math.min(point1.x, point2.x);
    const dy = Math.max(point1.y, point2.y) - Math.min(point1.y, point2.y);
    const distance = decimalSqrt(dx ** 2 + dy ** 2, 0);
    // 8192: 图片大小  8000: 地图真实渲染大小
    return Number(decimalMul(distance, (8000 / 8192)).toFixed(0));
}

// 获取中心点
export function getCenterPoint(p1: Location, p2: Location): Location {
    return {
        x: (p1.x + p2.x) / 2,
        y: (p1.y + p2.y) / 2,
    };
}