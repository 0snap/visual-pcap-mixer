import React from 'react';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import configureStore from './store/configureStore';
import App from './App';
import { loadAll } from './actions/apiActions';

const store = configureStore();
store.dispatch(loadAll());

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);