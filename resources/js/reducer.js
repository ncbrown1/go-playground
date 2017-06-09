import {
    UPDATE_CODE, SWITCH_LIVE, CLEAR_OUTPUT, ADD_OUTPUT, SET_SYSMSG
} from './constants';

const reducer = (state={}, action) => {
    switch (action.type) {
        case UPDATE_CODE:
            return Object.assign({}, state, {
                code: action.code,
                forceEditorRender: action.force
            });
        case SWITCH_LIVE:
            return Object.assign({}, state, {
                live: action.live
            });
        case CLEAR_OUTPUT:
            return Object.assign({}, state, {
                output: '',
                system: ''
            });
        case ADD_OUTPUT:
            return Object.assign({}, state, {
                output: state.output + action.output
            });
        case SET_SYSMSG:
            return Object.assign({}, state, {
                system: action.system
            });
        default:
            return state;
    }
};

export default reducer;