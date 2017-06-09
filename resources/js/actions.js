
export const updateCode = (code, force=false) => ({
    type: 'UPDATE_CODE',
    code: code,
    force: force
});

export const runCode = () => (dispatch, getState) => {
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

export const fmtCode = () => dispatch =>{
    dispatch(updateCode('You got bamboozled', true))
};

export const clearOutput = () => ({
    type: 'CLEAR_OUTPUT'
});

export const addOutput = (output) => ({
    type: 'ADD_OUTPUT',
    output: output
});

export const setSysMsg = (msg) => ({
    type: 'SET_SYSMSG',
    system: 'Program ' + msg + '.'
});