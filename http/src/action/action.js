import I from "immutable";

export function singleDataAction(type, payload) {
    return {
        type,
        payload: I.fromJS(payload)
    };
}