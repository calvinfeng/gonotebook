const path = require('path');

module.exports = {
    mode: "development",
    devtool: "inline-source-map",
    entry: "./index.tsx",
    output: {
        path: path.join(__dirname, '..', 'public'),
        filename: "index.js"
    },
    resolve: {
        extensions: [".ts", ".tsx", ".js", ".jsx"]
    },
    module: {
        rules: [
            { test: /\.tsx?$/, loader: "ts-loader" },
            { test: /\.scss$/, use: ["style-loader", "css-loader", "sass-loader"] }
        ]
    }
};