import {
    UPDATE_CODE, SWITCH_LIVE, CLEAR_OUTPUT, ADD_OUTPUT, SET_SYSMSG
} from './constants';

export const updateCode = (code, force=false) => ({
    type: UPDATE_CODE,
    code: code,
    force: force
});

export const runCode = () => (dispatch, getState) => {
    const { live } = getState();
    live ? wsRun(dispatch, getState) : httpRun(dispatch, getState);
};

const wsRun = (dispatch, getState) => {
    const { code } = getState();
    dispatch(clearOutput());

    let url = 'ws://localhost:8000/ws';
    let c = new WebSocket(url);

    let send = function(data) {
        // dispatch(addOutput((new Date())+ " ==> "+data+"\n"));
        c.send(data)
    };

    c.onmessage = (msg) => {
        // dispatch(addOutput((new Date())+ " <== "+msg.data+"\n"));
        console.log(msg.data);
        console.log(msg);
        let data = JSON.parse(msg.data);
        switch(data.kind) {
            case "stdout":
            case "stderr":
                dispatch(addOutput(data.body));
                break;
            case "end":
                dispatch(setSysMsg('exited'));
                break;
            default:
                break;
        }
    };

    c.onopen = () => { send(JSON.stringify({"id": '' + new Date().getUTCMilliseconds(), "kind": "run", "body": code})) };
};

const httpRun = (dispatch, getState) => {
    const { code } = getState();
    dispatch(clearOutput());
    dispatch(addOutput('Waiting for remote server...'));
    axios
        .post('/run-code', {code: code})
        .then((resp) => {
            console.log(resp.data);
            if (resp.data.hasOwnProperty('output')) {
                dispatch(clearOutput());
                dispatch(addOutput(resp.data.output));
                dispatch(setSysMsg('exited'));
            }
        })
        .catch((err) => {
            console.log(err);
        });
};

export const fmtCode = () => (dispatch, getState) => {
    const { code } = getState();
    axios
        .post('/fmt-code', {code: code})
        .then((resp) => {
            console.log(resp.data);
            if (resp.data.hasOwnProperty('output')) {
                dispatch(updateCode(resp.data.output, true));
            }
        })
        .catch((err) => {
            console.log(err);
        });
};

export const switchLive = (live) => ({
    type: SWITCH_LIVE,
    live: live
});

export const clearOutput = () => ({
    type: CLEAR_OUTPUT
});

export const addOutput = (output) => ({
    type: ADD_OUTPUT,
    output: output
});

export const setSysMsg = (msg) => ({
    type: SET_SYSMSG,
    system: 'Program ' + msg + '.'
});