import 'babel-polyfill';
import 'preact/devtools';
import { h, render } from 'preact';
import App from './app/app.js';

function renderApp() {
  render(
    <App />, document.body, document.body.lastChild
  );
}

renderApp();

if (module.hot) {
  module.hot.accept(renderApp);
}
