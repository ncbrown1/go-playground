
require('./setup');
require('./editor');

const initial_program = `
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, playground")
}
`.trim();

function resizeInput() {
    $(this).attr('size', $(this).val().length);
}

let share_url = $('#share-url');
share_url.keyup(resizeInput)
    // resize on page load
    .each(resizeInput);
$('#share').click(() => {
    share_url.addClass('active');
    share_url.focus();
});


import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import reducer from './reducer';
import Playground from './components/playground';


const stateStore = createStore(reducer, {
    live: false,
    code: initial_program,
    forceEditorRender: false,
    output: '',
    system: '',
    sharing: false
}, applyMiddleware(thunk));

render(
    <Provider store={stateStore}>
        <Playground />
    </Provider>,
    document.getElementById('root')
);