# React Tools

React is popular enough that I don't really need to explain what it is. Let me quickly go over couple tools we need to get React working. I assume that you have `npm` installed.  The choice of modern JS syntax is ES7 and TypeScript for this article.

### Babel Loader

JavaScript development touches a lot of diverse environments: Node, Chrome, Safari, etc. These various environments have different levels of compatibility with advanced JavaScript features like JSX and ES7. To ensure that our JSX and ES7 code works in any environment, we will use a _transpiler_ called **Babel** to convert our code into ES5, the universal, vanilla JavaScript compatible with all browsers and Node.

#### Configuration

Webpack can be configured to transpile your JSX and ES7 source code into browser-compatible JavaScript when creating the bundle.

```text
npm install --save @babel/core @babel/polyfill
npm install --save @babel/preset-env @babel/preset-react
npm install --save babel-loader
```

Create a `webpack.config.js`

```javascript
const path = require('path');

module.exports = {
    mode: "development",
    devtool: "inline-source-map",
    entry: "./src/index.jsx",
    output: {
        path: path.join(__dirname, '..', 'public'),
        filename: "index.js"
    },
    resolve: {
        extensions: [".js", ".jsx"],
    },
    module: {
        rules: [
            { 
                test: /\.jsx?$/,
                loader: "babel-loader",
                query: {
                    presets: ['@babel/env', '@babel/react']
                } 
            },
            { 
                test: /\.scss$/, 
                use: ["style-loader", "css-loader", "sass-loader"]
            }
        ]
    }
};
```

Here's a little tricky thing about using the latest ES7 feature, which is `async` and `await`. You will need `@babel/polyfill` to convert those new syntax into appropriate JavaScript to run on a browser. Within your source code, make sure to do the following if you happen to use `async/await` in your code.

```javascript
import '@babel/polyfill'
```

### TS Loader

TS loader is similar to Babel loader; instead of transpiling ES7 JavaScript source code, it transpile TypeScript source code into ES5 JavaScript!

```text
npm install ts-loader typescript --save
```

#### Configuration

The configuration is more or less the same except using `ts-loader`.

```javascript
const path = require('path');

module.exports = {
    mode: "development",
    devtool: "inline-source-map",
    entry: "./src/index.tsx",
    output: {
        path: path.join(__dirname, '..', 'public'),
        filename: "index.js"
    },
    resolve: {
        extensions: [".ts", ".tsx", ".js", ".jsx"],
    },
    module: {
        rules: [
            { 
                test: /\.tsx?$/, 
                loader: "ts-loader"
            },
            { 
                test: /\.scss$/, 
                use: ["style-loader", "css-loader", "sass-loader"]
            }
        ]
    }
};
```

### Style Loaders

Notice that I also have `style-loader`, `css-loader`, and `sass-loader`, because I can embed CSS into JavaScript using these loaders. 

### Package.json

In general, your `package.json` should look like the following with everything installed, i.e. React & Redux.

#### JavaScript ES7

```javascript
{
  "name": "todojs",
  "version": "1.0.0",
  "description": "simple todo application written in React/Redux JavaScript",
  "scripts": {
    "build": "webpack --watch"
  },
  "keywords": [
    "javascript",
    "react",
    "redux",
    "ES7"
  ],
  "author": "Calvin Feng",
  "license": "ISC",
  "dependencies": {
    "@babel/core": "^7.4.5",
    "@babel/polyfill": "^7.4.4",
    "@babel/preset-env": "^7.4.5",
    "@babel/preset-react": "^7.0.0",
    "axios": "^0.19.0",
    "babel-loader": "^8.0.6",
    "css-loader": "^3.0.0",
    "node-sass": "^4.12.0",
    "react": "^16.8.6",
    "react-dom": "^16.8.6",
    "react-redux": "^7.1.0",
    "redux": "^4.0.1",
    "redux-thunk": "^2.3.0",
    "sass-loader": "^7.1.0",
    "style-loader": "^0.23.1",
    "webpack": "^4.35.0",
    "webpack-cli": "^3.3.5",
    "webpack-dev-server": "^3.7.2"
  }
}
```

#### TypeScript

```javascript
{
  "name": "todots",
  "version": "1.0.0",
  "description": "simple todo application written in React/Redux TypeScript",
  "scripts": {
    "build": "webpack --watch"
  },
  "keywords": [
    "typescript",
    "react",
    "redux"
  ],
  "author": "Calvin Feng",
  "license": "ISC",
  "dependencies": {
    "@types/react": "^16.8.22",
    "@types/react-dom": "^16.8.4",
    "@types/react-redux": "^7.1.0",
    "@types/redux": "^3.6.0",
    "axios": "^0.19.0",
    "css-loader": "^3.0.0",
    "node-sass": "^4.12.0",
    "react": "^16.8.6",
    "react-dom": "^16.8.6",
    "react-redux": "^7.1.0",
    "redux": "^4.0.1",
    "redux-thunk": "^2.3.0",
    "sass-loader": "^7.1.0",
    "style-loader": "^0.23.1",
    "ts-loader": "^6.0.4",
    "typescript": "^3.5.2",
    "webpack": "^4.35.0",
    "webpack-cli": "^3.3.5",
    "webpack-dev-server": "^3.7.2"
  }
}
```

