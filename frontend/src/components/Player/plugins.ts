import Artplayer from "artplayer";

export function hotKeyPlugin(art: Artplayer) {
    const keys: Record<string, Array<(event: Event) => void>> = {};
    function add(key: string, event: (event: Event) => void) {
        if(keys[key]) {
            if(!keys[key].includes(event)) {
                keys[key].push(event);
            }
        } else {
            keys[key] = [event];
        }
    }

    add("Escape", () => {
        if(art.fullscreenWeb) {
            art.fullscreenWeb = false;
        }
    });
    add("ArrowUp", () => {
        art.volume += Artplayer.VOLUME_STEP;
    });
    add("ArrowDown", () => {
        art.volume -= Artplayer.VOLUME_STEP;
    });

    art.on("document:keydown", e => {
        const event = e as KeyboardEvent;
        if(art.isFocus) {
            const tag = document.activeElement?.tagName.toUpperCase();
            const editable = document.activeElement?.getAttribute("contenteditable");
            if(
                tag !== "INPUT"
                && tag !== "TEXTAREA"
                && editable !== ""
                && editable !== "true"
                && !event.altKey
                && !event.ctrlKey
                && !event.metaKey
                && !event.shiftKey
            ) {
                const events = keys[event.code];
                if(events) {
                    event.preventDefault();
                    for(let index = 0; index < events.length; index++) {
                        events[index].call(art, event);
                    }
                    art.emit("hotkey", event);
                }
            }
        }
        art.emit("keydown", event);
    });
}